package account

type Transaction struct {
	User   string  `json:"user,omitempty" db:"user"`
	Id     int     `json:"id" db:"id"`
	Date   string  `json:"date" db:"date"`
	Object string  `json:"object" db:"object"`
	Amount float32 `json:"amount" db:"amount"`
}

type Account struct {
	User         string        `json:"user" db:"user"`
	Currency     string        `json:"currency" db:"currency"`
	Description  string        `json:"description" db:"description"`
	Balance      float32       `json:"balance" db:"balance"`
	Transactions []Transaction `json:"transactions"`
}
