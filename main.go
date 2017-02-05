package main

import (
    "encoding/json"
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "wrtapi/utils"
)

func ubusResponse(s func() ([]byte, error), res *gin.Context) {
    stat, err := s()
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

    router.GET("/wlan", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.WirelessStatus, c)
    })

    router.GET("/system", func(c *gin.Context) {
        ubusResponse(utils.UbusCall.Info, c)
    })

    router.Run() // 0.0.0.0:8080
}
