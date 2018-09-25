package main

import "github.com/gin-gonic/gin"

func main() {								                  // gin is like flask , a microframework
  r := gin.Default()                          // Initialize server with default settings , r is your router
  r.GET("/paryul", func(c *gin.Context) {           // Creates route for root(/), second argument is function to execute
    c.JSON(200, gin.H{
      "message": "Hello World",
    })
  })
  r.Run()                                     // By default, listen and serve on http://localhost:8080
}
