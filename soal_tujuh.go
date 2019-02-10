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

type Minmaxcurr struct {
    Curr string
    Rate float64
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

func ShowMinMaxCurr(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    parm := mux.Vars(r)
    _pRateDt := parm["dt"]
    selDB, err := db.Query("select curr,rates from tb_rates where tgl = ? order by rates",_pRateDt)
    if err != nil {
        panic(err.Error())
    }
    minMax := Minmaxcurr{}
    res := []Minmaxcurr{}
    for selDB.Next() {
        var currcd string
        var minmaxrates float64
        err = selDB.Scan(&currcd, &minmaxrates)
        if err != nil {
            panic(err.Error())
        }
        minMax.Curr = currcd
        minMax.Rate = minmaxrates
        res = append(res, minMax)
    }
    fmt.Println(res)
    json.NewEncoder(w).Encode(res)
    
    defer db.Close()
}

func handleRequests() {
    log.Println("Server started on: http://localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true) 
    myRouter.HandleFunc("/rates/{dt}", ShowMinMaxCurr)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    handleRequests()
}