package main

import (
	"log"
	"net/http"
	"time"
)

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "hello1"))
    log.Println(time.Now(),":hello1")
}

func main() {
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":9000", nil)
}