package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/barbatoast/bank/account"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB Database

type Database struct {
	driver *sqlx.DB
	mutex  *sync.Mutex
}

func (database *Database) Connect(path string) {
	driver, err := sqlx.Open("sqlite3", path)
	if err != nil {
		log.Fatalln(err)
	}
	database.driver = driver
	database.mutex = &sync.Mutex{}
}

func (database *Database) GetAccount(user string) account.Account {
	database.mutex.Lock()
	defer database.mutex.Unlock()
	query := fmt.Sprintf(`SELECT currency, description, balance
		FROM accounts where user = '%s'`, user)
	row := database.driver.QueryRow(query)
	var currency string
	var description string
	var balance float32
	_ = row.Scan(&currency, &description, &balance)

	result := account.Account{
		User: user, Currency: currency,
		Description: description, Balance: balance}

	query = fmt.Sprintf(`SELECT id, date, object, amount
		FROM transactions where user = '%s'`, user)
	rows, _ := database.driver.Query(query)
	for rows.Next() {
		var id int
		var date string
		var object string
		var amount float32
		_ = rows.Scan(&id, &date, &object, &amount)
		result.Transactions = append(result.Transactions,
			account.Transaction{Id: id, Date: date, Object: object, Amount: amount})
	}

	return result
}
