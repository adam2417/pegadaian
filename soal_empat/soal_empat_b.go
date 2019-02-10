package main

import (
    "fmt"
    "strconv"
    "github.com/gorilla/mux"
    "net/http"
    "html/template"
    "log"
)

func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func Index(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("soal_empat_a.tmpl")
    t.Execute(w,nil)
}

func ShowNumeric(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        _v := r.FormValue("txtInput")
        if isNumeric(_v) == true {
            fmt.Fprintf(w, "%v adalah angka",_v)
        } else {
            fmt.Fprintf(w, "%v bukan angka",_v)
        }
    }
}

func handleRequests() {
    log.Println("Server started on: http://localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true) 
    myRouter.HandleFunc("/", Index)
    myRouter.HandleFunc("/showisnumber", ShowNumeric)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main(){
    handleRequests()
} 