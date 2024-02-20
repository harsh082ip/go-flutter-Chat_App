package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/controllers/usersController"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/user/getuserbyusername/:username", usersController.GetUserByUsername)
	incomingRoutes.GET("/user/addtorecentlyviewed/:uid", usersController.AddUserToRecentlyViewed)
	incomingRoutes.GET("/user/fetchhomedata/:uid", usersController.FetchRecentUsers)
}
