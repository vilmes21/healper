package main

import (
	// "encoding/json"
	// "log"
	// "net/http"
	// "../db1"
	// "../config"
	// "../diseOrgsympField"
	// "../fn"
	// "../model"
	// "../viewmodel"
	"fmt"
    // "strings"
	// "dbbase/sql"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func test () {
	// var t1 = []string {"no", "oi"}
	// fmt.Println(`fn.Sql.SelectFrom---` + (fn.Sql).SelectFrom(t1, `tst`)) 

	fmt.Println(`fn.Sql.SelectFrom---`) 

	fmt.Println(sqlFn.Select([]string {"first", "id"}).From("booo").End())
	
	fmt.Println(`Test DONE.`)
}
