package main

import (
	"fmt"
	"web/todo/pkg/common/config"
	"web/todo/pkg/common/db"
	"web/todo/pkg/common/helpers"

	"github.com/gin-gonic/gin"
)

var Modules []helpers.Module = []helpers.Module{
	&config.ConfigModule{},
	&db.DBModule{},
}

func main() {
	helpers.InitializeModule(Modules)
	fmt.Println("Web API todo start")
	r := gin.Default()

	r.Run(config.Env.Port)
}
