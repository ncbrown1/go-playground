package app

import (
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) (*gin.Engine){

    r.POST("/run-code", RunCode)

    return r
}