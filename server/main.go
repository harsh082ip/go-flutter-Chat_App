package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/routes"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/ws"
)

const (
	WEBPORT = "0.0.0.0:8006"
)

func main() {

	router := gin.Default()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()
	ws.SaveChatsToDatabase()

	routes.AuthRoutes(router)
	routes.MiscRoutes(router)
	routes.UserRoutes(router, wsHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Status": "Welcome to the chat-hub Api"})
	})

	router.Run(WEBPORT)

}
