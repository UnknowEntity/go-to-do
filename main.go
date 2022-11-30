package main

import (
	"fmt"
	"web/todo/pkg/common/config"
	"web/todo/pkg/common/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: Create function to handle initial loading
	fmt.Println("Web API todo start")
	config.LoadConfig()
	db.InitDB(config.Value.DB_URL)
	r := gin.Default()

	r.Run(config.Value.Port)
}
