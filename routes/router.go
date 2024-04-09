package routes

import (
	"encoding/json"
	"net/http"

	"github.com/barbatoast/bank/database"
)

func HandleAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	account := database.DB.GetAccount(r.PathValue("user"))
	json.NewEncoder(w).Encode(account)
}
