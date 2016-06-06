package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "log"
	"runtime"
  "bytes"
)

var (
	buf    bytes.Buffer
	logger *log.Logger = log.New(&buf, "log: ", log.Lshortfile)
)

func main() {
    logger.Println("Application starts.")
    // User all cpu cores
    runtime.GOMAXPROCS(runtime.NumCPU())

    router := gin.Default()

    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    // However, this one will match /user/john/ and also /user/john/send
    // If no other routers match /user/john, it will redirect to /user/john/
    router.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        message := name + " is " + action
        c.String(http.StatusOK, message)
    })

    // PORT environment variable was defined.
    router.Run(":80") // for a hard coded port
}
