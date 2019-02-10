package main

import (
    "fmt"
    "database/sql"
    "net/http"
    "log"
    "html/template"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

type Minmax struct {
    Minmaxval int
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

func CalcMinMax(p_minmax []Minmax,w http.ResponseWriter){
    var arr []Minmax
    var n, biggest, smallest int
    arr = p_minmax
        
    for _,rgbig := range arr {
        if rgbig.Minmaxval > n {
          n = rgbig.Minmaxval
          biggest = n
        }
    }
    fmt.Println("Angka terbesar: ", biggest)
    
    for _,rgsmall := range arr {
        if rgsmall.Minmaxval < n {          
          n = rgsmall.Minmaxval
          smallest = n
        } 
    }
    fmt.Println("Angka terkecil: ", smallest)
    
    _allres := make(map[string]int)
    _allres["angka_terbesar"] = biggest
    _allres["angka_terkecil"] = smallest
    
    json.NewEncoder(w).Encode(_allres)
}

func Index(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("soal_dua_a.tmpl")
    t.Execute(w,nil)
}

func ShowMinMax(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM tb_max_min")
    if err != nil {
        panic(err.Error())
    }
    minMax := Minmax{}
    res := []Minmax{}
    for selDB.Next() {
        var valminmax int
        err = selDB.Scan(&valminmax)
        if err != nil {
            panic(err.Error())
        }
        minMax.Minmaxval = valminmax
        res = append(res, minMax)
    }
    CalcMinMax(res,w)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        _v := r.FormValue("txtVal")
        insForm, err := db.Prepare("INSERT INTO tb_max_min(val) VALUES(?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(_v)
        log.Println("INSERT: tb_max_min Value: " + _v)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func handleRequests() {
    log.Println("Server started on: http://localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true) 
    myRouter.HandleFunc("/", Index)
    myRouter.HandleFunc("/insert", Insert)
    myRouter.HandleFunc("/get_hitung_angka", ShowMinMax)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main(){
    handleRequests()
}