package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", hanleEvents)
	r.POST("/with", hanleWith)
	r.POST("/without", hanleWithout)
	r.NoRoute(func(c *gin.Context) {
		fmt.Printf("gc: %+v\n", gc)
		c.Data(404, "text/plain", []byte(toString(gc)))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
