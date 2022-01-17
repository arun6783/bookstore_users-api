package users

import (
	"net/http"

	"github.com/arun6783/bookstore_users-api/domain/users"
	"github.com/arun6783/bookstore_users-api/services"
	"github.com/arun6783/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadResuestError("Invalid json body")

		c.JSON(int(restErr.Status), restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {

		c.JSON(int(saveErr.Status), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {

}

func SearchUser(c *gin.Context) {

}
