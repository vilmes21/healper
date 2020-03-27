package main

import (
	// "../db1"
	// "log"
	// "fmt"
	"net/http"
	// "dbbase/sql"

	// "github.com/gorilla/mux"
	// _ "github.com/lib/pq"
)

func midw(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Content-Type", "application/json")
        f(w, r)
    }
}