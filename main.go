package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Web API todo start")
	r := gin.Default()

	r.Run()
}
