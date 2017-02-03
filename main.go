package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    "lede-mobile/utils"
    "encoding/json"
    "net/http"
)

func responder(stat []byte, err error, res *gin.Context) {
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
        stat, err := utils.SystemInfo()
        responder(stat, err, c)
    })

    router.GET("/wan", func(c *gin.Context) {
        stat, err := utils.StatusWan()
        responder(stat, err, c)
    })

    router.Run() // listen and serve on 0.0.0.0:8080
}