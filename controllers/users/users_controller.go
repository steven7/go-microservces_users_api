package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-microservces_users_api/domain/users"
	"go-microservces_users_api/services"
	"go-microservces_users_api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	// convert user id from string to int
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		return 0, err
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(user)

	result, dbError  := services.UsersService.CreateUser(user)
	fmt.Println(result)
	fmt.Println(dbError)

	if dbError != nil {
		c.JSON(dbError.Status, dbError)
		return
	}
	c.JSON(http.StatusCreated, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	// convert user id from string to int
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getError  := services.UsersService.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

//func Search(c *gin.Context) {
//	status := c.Query("status")
//
//	users, err := services.UsersService.Search(status)
//	if err != nil {
//		c.JSON(err.Status, err)
//		return
//	}
//	result := make([]interface{}, len(users))
//	for index, user := range users {
//		result[index] = user.Marshall(c.GetHeader("X-Public") == "true")
//	}
//	c.JSON(http.StatusOK, result)
//}