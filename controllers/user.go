package controllers

import (
	"fmt"
	models "models"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		fmt.Println("Error in getting all the users")
		return
	}
	c.IndentedJSON(http.StatusOK, users)

}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error in reading the response", err)
		return
	}
	id, err := models.AddUser(&user)
	if err != nil {
		fmt.Println("Error in adding the user", err)
		return
	}
	fmt.Println(id)
	c.IndentedJSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error in reading the response", err)
		return
	}
	err := models.UpdateUser(&user)
	if err != nil {
		fmt.Println("Error in updating the user", err)
		return
	}
	c.IndentedJSON(http.StatusOK, "successful")

}
