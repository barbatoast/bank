package main

import (
	"net/http"

	"github.com/barbatoast/bank/database"
	"github.com/barbatoast/bank/routes"
	"github.com/rs/cors"
)

func main() {
	database.DB.Connect("./migration/db.sqlite3")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/accounts/{user}", routes.HandleAccounts)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
