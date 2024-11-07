package main

import (
    "fmt"
    "net/http"
)

func handlerStatus(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "OK! I'm alive!")    
}