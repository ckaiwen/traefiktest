package main

import (
	"log"
	"net/http"
	"time"
)

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "test2"))
    log.Println(time.Now(),":test2")
}

func main() {
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":9000", nil)
}