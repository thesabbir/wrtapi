package main

import (
    "encoding/json"
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "wrtapi/utils"
    "log"
)

func ubusResponse(s func() ([]byte, error), res *gin.Context) {
    stat, err := s()
    log.Printf("%s %s", stat, err)
    if err != nil {
        res.JSON(http.StatusNotImplemented, gin.H{"error": "ubus failed"})
    } else {
        var parsed interface{}
        json.Unmarshal(stat, &parsed)
        res.JSON(http.StatusOK, parsed)
    }
}

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.BoardInfo, c)
    })

    router.GET("/wan", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.WanStatus, c)
    })

    router.GET("/lan", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.LanStatus, c)
    })

    router.GET("/wlan", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.WirelessStatus, c)
    })

    router.GET("/system", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.Info, c)
    })

    router.GET("/services", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.ServiceList, c)
    })

    router.GET("/interfaces", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.InterfacesList, c)
    })

    router.GET("/config/network",  func(c *gin.Context) {
        ubusResponse(utils.UbusCall.NetworkConfig, c)
    })

    router.Run() // 0.0.0.0:8080
}
