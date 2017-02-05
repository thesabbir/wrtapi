package main

import (
    "net/http"
    "log"
)

func main() {
    router := RouterInit()
    log.Fatal(http.ListenAndServe(":8080", router))
}
