package main

import "goProject/router"

func main() {
	//dao.ConnAndInsertAndSelect()
	r := router.SetUpRouter()
	r.Run(":8080")
}
