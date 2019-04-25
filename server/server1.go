//server1 is a echo server
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)//echo request transfer process program
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//process program echo request URL r's path
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
