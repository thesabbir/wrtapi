package main

import (
    "github.com/julienschmidt/httprouter"
    "wrtapi/utils"
    "net/http"
)

func RouterInit() http.Handler {
    router := httprouter.New()
    router.GET("/", utils.Handler(func(w http.ResponseWriter) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"wrtapi": "0.1"}`))
    }));
    router.GET("/system", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.Info, w)
    }))
    router.GET("/system/board", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.BoardInfo, w)
    }))
    router.GET("/system/services", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.ServiceList, w)
    }))
    router.GET("/network", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.NetworkConfig, w)
    }))
    router.GET("/network/interfaces", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.InterfacesList, w)
    }))
    router.GET("/network/interfaces/wan", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.WanStatus, w)
    }))
    router.GET("/network/interfaces/lan", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.LanStatus, w)
    }))
    router.GET("/wlan", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.WirelessStatus, w)
    }))
    router.GET("/wlan/clients", utils.Handler(func(w http.ResponseWriter) {
        utils.UbusResponse(utils.UbusCall.WlanClients, w)
    }))
    return router
}