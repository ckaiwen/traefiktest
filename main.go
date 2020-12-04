package main

import (
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "hello1"))
    fmt.Println(time.Now(),":hello1")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":9000", nil)
}