package main

import (
	_ "questions/pkg/database"
	"questions/pkg/config"
	"questions/pkg/logger"
	"questions/routes"
)

var err error

func main() {

	r := routes.RegisterRouters()

	logger.Info("http server started, listened on port ", config.Config.Port)
	r.Run(":8080")
}
