package users

import (
	"net/http"
	"strconv"

	"github.com/arun6783/bookstore_users-api/domain/users"
	"github.com/arun6783/bookstore_users-api/services"
	"github.com/arun6783/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(paramUserId string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(paramUserId, 10, 64)
	if userErr != nil {

		return 0, errors.NewBadResuestError("user id should be a number")
	}
	return userId, nil
}
func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(int(userErr.Status), userErr)
		return
	}

	result, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(int(getErr.Status), getErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {

	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(int(userErr.Status), userErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadResuestError("Invalid json body")
		c.JSON(int(restErr.Status), restErr)
		return
	}

	user.Id = userId
	var isPartialUpdate = c.Request.Method == http.MethodPut

	result, getErr := services.UpdateUser(isPartialUpdate, user)

	if getErr != nil {
		c.JSON(int(getErr.Status), getErr)
		return
	}

	c.JSON(http.StatusOK, result)

}
func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(int(userErr.Status), userErr)
		return
	}

	delresult := services.Delete(userId)
	if delresult != nil {
		c.JSON(int(delresult.Status), delresult)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := services.Search(status)

	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}

	c.JSON(http.StatusOK, users)
}
