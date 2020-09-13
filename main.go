package main

import(
"fmt"
"time"
"log"
"net/http"
"github.com/gorilla/mux"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
)

const homeMessage = "Hello Wano!"

func home(writer http.ResponseWriter, request *http.Request){
    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte(homeMessage))
}


func main() {
//     srv := &http.Server{
//         Addr: ":8080",
//         ReadTimeout: 5 * time.Second,
//         WriteTimeout: 10 * time.Second,
//         IdleTimeout: 120 * time.Second,
//         Handler: mux,
//     }
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", home)
    log.Println("Escaping Whole Cake Island!")
    err := http.ListenAndServe(":8080", router)
    fmt.Println(err)
    log.Fatalf("Kaido failed the server",err)
}