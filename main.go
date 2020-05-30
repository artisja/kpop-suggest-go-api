package main

import(
"fmt"
"log"
"net/http"
"github.com/gorilla/mux"
)

func home(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintf(writer, "Hello Wano!")
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home)
    log.Println("Escaping Whole Cake Island!")
    log.Fatal(http.ListenAndServe(":8080", router))
}