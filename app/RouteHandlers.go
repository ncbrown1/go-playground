package app

import (
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
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
    log.Println(run_code.Code)
    c.Writer.WriteHeader(http.StatusNoContent)
}