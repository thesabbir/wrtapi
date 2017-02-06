package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func UbusResponse(s func() ([]byte, error), w http.ResponseWriter) {
    stat, err := s()
    w.Header().Set("Content-Type", "application/json")
    if err != nil {
        w.Write([]byte(`{"error":"ubus failed"}`))
    } else {
        w.Write(stat)
    }
}

func Handler(x func(http.ResponseWriter)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        x(w)
    }
}