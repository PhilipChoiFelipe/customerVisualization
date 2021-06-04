package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/info441-sp21/final-project/server/gateway/handlers"
	"github.com/info441-sp21/final-project/server/gateway/models/users"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
)

func main() {
	//redis address
	redisAdd := os.Getenv("REDISADDR")
	//sql address
	dsn := os.Getenv("DSN")

	//session id key
	signingKey := os.Getenv("SESSIONKEY")

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

	userStorage := users.NewSqlStorage(db)

	client := redis.NewClient(&redis.Options{
		Addr: redisAdd,
		DB:   0,
	})
	sessionStore := sessions.NewRedisStore(client, time.Hour)

	httpHandler := handlers.HttpHandler{SigningKey: signingKey, SessionStore: sessionStore, UserStorage: userStorage}

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	if len(tlsKeyPath) == 0 || len(tlsCertPath) == 0 {
		log.Print("appropriate tls path is not provided")
		os.Exit(1)
		return
	}

	userDSAddrs := os.Getenv("USERDSADDR")
	userDSURLs := GetURLs(userDSAddrs)
	userDSProxy := &httputil.ReverseProxy{Director: CustomDirector(userDSURLs, &httpHandler)}

	router := mux.NewRouter()
	//user handlers
	router.HandleFunc("/v1/user", httpHandler.UsersHandler)
	router.HandleFunc("/v1/user/{user_id}", httpHandler.SpecificUserHandler)
	router.HandleFunc("/v1/sessions", httpHandler.SessionsHandler)
	router.HandleFunc("/v1/sessions/{session_id}", httpHandler.SpecificSessionHandler)

	//MICROSERVICE
	//customer handlers
	router.Handle("/v1/user/{user_id}/customers", userDSProxy)
	router.Handle("/v1/user/{user_id}/customers/{customer_id}", userDSProxy)

	//item handlers
	router.Handle("/v1/user/{user_id}/items", userDSProxy)
	router.Handle("/v1/user/{user_id}/items/{item_id}", userDSProxy)

	wrappedMux := handlers.NewCORSHandler(router)

	log.Printf("server is listening at port: %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
	// log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
