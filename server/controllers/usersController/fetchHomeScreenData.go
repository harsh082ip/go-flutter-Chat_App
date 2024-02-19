package usersController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers"
	authhelper "github.com/harsh082ip/go-flutter-Chat_App/tree/main/server/helpers/authHelper"
)

func FetchHomeScreenData(c *gin.Context) {

	uid := c.Param("uid")
	jwt_token := c.Query("jwtkey")

	if uid != "" && jwt_token != "" {

		_, err := authhelper.CheckAuthorization(jwt_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized Access",
				"error":  err.Error(),
			})
			return
		}

		collName := "Recently_Viewed"
		docExists, _ := helpers.CheckIfDocumentExists(uid, collName, true)
		if !docExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "User id is invalid",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Missing Details",
			"error":  "Please provide all the details",
		})
		return
	}
}
