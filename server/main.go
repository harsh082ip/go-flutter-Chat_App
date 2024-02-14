package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/routes"
)

const (
	WEBPORT = ":8006"
)

func main() {

	router := gin.Default()

	routes.AuthRoutes(router)
	routes.MiscRoutes(router)

	router.Run(WEBPORT)

}
