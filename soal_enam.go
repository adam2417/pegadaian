package main

import (
    "fmt"
    "database/sql"
    //"encoding/json"
    _ "github.com/go-sql-driver/mysql"
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

func ShowMinMaxCurr() {
    db := dbConn()
    selDB, err := db.Query("select curr,rates from tb_rates where tgl = (select max(tgl) from tb_rates) order by rates")
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
    fmt.Println("By rate: ", res)
    
    defer db.Close()
}

func main() {
    ShowMinMaxCurr()
}