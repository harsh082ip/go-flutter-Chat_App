package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/controllers/usersController"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/user/getuserbyusername/:username", usersController.GetUserByUsername)
	incomingRoutes.GET("/user/addtorecentlyviewed/:uid", usersController.AddUserToRecentlyViewed)
	incomingRoutes.GET("/user/fetchrecentusers/:uid", usersController.FetchRecentUsers)
	incomingRoutes.GET("/user/removerecentuser/:uid", usersController.RemoveRecentUser)
	incomingRoutes.GET("/user/addtorecentchats/:uid", usersController.AddUserToRecentChats)
	incomingRoutes.GET("/user/fetchhomedata/:uid", usersController.FetchHomeUsers)
	incomingRoutes.GET("/user/removehomeuser/:uid", usersController.RemoveHomeUsers)
}
