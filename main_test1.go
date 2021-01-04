package main

import (
	"log"
	"net/http"
	"time"
)

func IndexHandler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "test1"))
    log.Println(time.Now(),":test1")
}

func main() {
	http.HandleFunc("/test1", IndexHandler1)
	http.ListenAndServe(":9000", nil)
}