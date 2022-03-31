package main

import (
	"goProject/logger"
	"goProject/router"
)

func main() {
	//dao.ConnAndInsertAndSelect()

	r := router.SetUpRouter()
	logger.Mlogger.Warn("server begin running")
	//	logger.Mlogger.Sync()
	r.Run(":8080")
}
