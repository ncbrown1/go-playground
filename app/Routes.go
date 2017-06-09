package app

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) (*gin.Engine){

    r.LoadHTMLFiles("public/index.html")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    r.POST("/run-code", RunCode)
    r.POST("/fmt-code", FmtCode)

    authenticated := r.Group("/", gin.BasicAuth(gin.Accounts{
        "cgaucho": "foobar0",  // 1. user:cgaucho  password:foobar0
        "starbuck": "c0vf3f3", // 2. user:starbuck password:c0vf3f3
        "admin": "admin",      // 3. user:admin    password:admin
    }))
    authenticated.GET("/ws", func(c *gin.Context) {
        wshandler(c.Writer, c.Request)
    })

    return r
}