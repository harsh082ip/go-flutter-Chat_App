package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/controllers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/users/signup", controllers.SignUp)
	incomingRoutes.POST("/users/login", controllers.Login)
	incomingRoutes.POST("/users/getrefreshtoken", authhelper.GenerateRefreshToken)
}
