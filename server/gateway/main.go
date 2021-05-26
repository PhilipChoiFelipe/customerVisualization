package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()

	log.Printf("server is listening at port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
