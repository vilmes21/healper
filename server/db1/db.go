package db1

import (
	"../key"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	// "time"
)

func Init() *sql.DB {
	connectInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		key.Host, key.Port, key.User, key.Dbname)

	db, err := sql.Open("postgres", connectInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

