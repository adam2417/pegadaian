package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "strconv"
)

type Envelope struct {
	Cube []struct {
		Date  string `xml:"time,attr"`
		Rates []struct {
			Currency string `xml:"currency,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"Cube"`
	} `xml:"Cube>Cube"`
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

type Rate struct {
    _tgl string
    _curr string
    _rate float64
}

func InsertRate(_rate Rate) {
    db := dbConn()    
    insForm, err := db.Prepare("INSERT INTO tb_rates(tgl,curr,rates) VALUES(?,?,?)")
    if err != nil {
        panic(err.Error())
    }        
    insForm.Exec(_rate._tgl,_rate._curr,_rate._rate)
    
    defer db.Close()
}

func DeleteAllRate() {
    db := dbConn()
    delForm, err := db.Prepare("DELETE FROM tb_rates")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec()    
    defer db.Close()    
}

func main() {	
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	xmlCurrenciesData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var env Envelope
	err = xml.Unmarshal(xmlCurrenciesData, &env)

	if err != nil {
		log.Fatal(err)
	}
    
    fmt.Println("Delete all rate first")
    DeleteAllRate()
    fmt.Println("Insert rates to table")
    
    for _, _dt := range env.Cube {        
        for _, v := range _dt.Rates {
            _rt := Rate{}
            _rt._curr = v.Currency
            
            f, err := strconv.ParseFloat(v.Rate, 64)
            if err != nil{
                log.Fatal(err)
            }
            
            _rt._rate = f
            _rt._tgl = _dt.Date
            InsertRate(_rt)            
        }
    }
    fmt.Println("All data inputted!")
}
