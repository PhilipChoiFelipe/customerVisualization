package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/info441-sp21/final-project/server/userDataService/handlers"
	"github.com/info441-sp21/final-project/server/userDataService/models/customers"
	"github.com/info441-sp21/final-project/server/userDataService/models/items"
	// "github.com/info441-sp21/final-project/server/gateway/handlers"
)

func main() {
	dsn := os.Getenv("DSN")

	//session id key
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("sql error pinging database: %v\n", err)
	} else {
		fmt.Printf("sql successfully connected!\n")
	}

	itemStorage := items.NewSqlStorage(db)
	customerStorage := customers.NewSqlStorage(db)

	httpHandler := handlers.HttpHandler{ItemStorage: itemStorage, CustomerStorage: customerStorage}

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	router := mux.NewRouter()
	//customer handlers
	router.HandleFunc("/v1/user/{user_id}/customers", httpHandler.CustomersHandler)
	router.HandleFunc("/v1/user/{user_id}/customers/{customer_id}", httpHandler.SpecificCustomerHandler)

	//item handlers
	router.HandleFunc("/v1/user/{user_id}/items", httpHandler.ItemsHandler)
	router.HandleFunc("/v1/user/{user_id}/items/{item_id}", httpHandler.SpecificItemHandler)

	log.Printf("server is listening at port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
