package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/controllers/usersController"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/ws"
)

func UserRoutes(incomingRoutes *gin.Engine, wsHandler *ws.Handler) {

	incomingRoutes.GET("/user/getuserbyusername/:username", usersController.GetUserByUsername)
	incomingRoutes.GET("/user/addtorecentlyviewed/:uid", usersController.AddUserToRecentlyViewed)
	incomingRoutes.GET("/user/fetchrecentusers/:uid", usersController.FetchRecentUsers)
	incomingRoutes.GET("/user/removerecentuser/:uid", usersController.RemoveRecentUser)
	incomingRoutes.GET("/user/addtorecentchats/:uid", usersController.AddUserToRecentChatsAndCreateRoom)
	incomingRoutes.GET("/user/fetchhomedata/:uid", usersController.FetchHomeUsers)
	incomingRoutes.GET("/user/removehomeuser/:uid", usersController.RemoveHomeUsers)
	incomingRoutes.GET("/user/getroomidbyusernames", usersController.GetRoomIDByUsernames)
	incomingRoutes.GET("/ws/joinroom/:roomId", wsHandler.JoinRoom)
	incomingRoutes.GET("/user/fetchchatsfromdatabase/:roomId", usersController.FetchChatsFromDB)
}
