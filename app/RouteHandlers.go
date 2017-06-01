package app

import (
    "net/http"
    //"log"
    "github.com/gin-gonic/gin"
    "github.com/ncbrown1/go-playground/app/runtime"
)

type RunCodeJSON struct {
    Code string `form:"code" json:"code" binding:"required"`
}

func RunCode(c *gin.Context) {
    var run_code RunCodeJSON
    if err := c.BindJSON(&run_code); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
        return
    }
    //log.Println(run_code.Code)
    result := runtime.RunCode(run_code.Code)
    c.JSON(http.StatusOK, result)
}
