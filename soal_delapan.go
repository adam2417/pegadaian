package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "encoding/json"
)

type CurrAnalyze struct {
    Curr string
    Minrate float64
    Maxrate float64
    Avgrate float64
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "pegadaian"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func ShowCurrAnalyze(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("select curr,min(rates) _min, max(rates)_max, avg(rates) _avg from tb_rates group by curr")
    if err != nil {
        panic(err.Error())
    }
    minMax := CurrAnalyze{}
    res := []CurrAnalyze{}
    for selDB.Next() {
        var currcd string
        var _minrate float64
        var _maxrate float64
        var _avgrate float64
        err = selDB.Scan(&currcd, &_minrate, &_maxrate, &_avgrate)
        if err != nil {
            panic(err.Error())
        }
        minMax.Curr = currcd
        minMax.Minrate = _minrate
        minMax.Maxrate = _maxrate
        minMax.Avgrate = _avgrate
        res = append(res, minMax)
    }
    fmt.Println(res)
    json.NewEncoder(w).Encode(res)
    
    defer db.Close()
}

func handleRequests() {
    log.Println("Server started on: http://localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true) 
    myRouter.HandleFunc("/rates/analyze", ShowCurrAnalyze)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    handleRequests()
}