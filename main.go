package main

import (
	"fmt"
	"web/todo/modules/common/config"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Web API todo start")
	config.LoadConfig()
	r := gin.Default()

	r.Run(config.Value.Port)
}
