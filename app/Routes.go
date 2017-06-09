package app

import (
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) (*gin.Engine){

    r.LoadHTMLFiles("public/index.html")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    r.POST("/run-code", RunCode)

    return r
}