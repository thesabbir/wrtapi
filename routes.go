package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func RouterInit() http.Handler {
    router := httprouter.New()
    router.GET("/", Handler(func(w http.ResponseWriter) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"wrtapi": "0.1"}`))
    }));
    router.GET("/system", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.Info, w)
    }))
    router.GET("/system/board", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.BoardInfo, w)
    }))
    router.GET("/system/services", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.ServiceList, w)
    }))
    router.GET("/network", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.NetworkConfig, w)
    }))
    router.GET("/network/interfaces", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.InterfacesList, w)
    }))
    router.GET("/network/interfaces/wan", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.WanStatus, w)
    }))
    router.GET("/network/interfaces/lan", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.LanStatus, w)
    }))
    router.GET("/wlan", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.WirelessStatus, w)
    }))
    router.GET("/wlan/clients", Handler(func(w http.ResponseWriter) {
        UbusResponse(UbusCall.WlanClients, w)
    }))
    return router
}