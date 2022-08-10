package user

import (
	"bohemian/practice/pkg/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "user"

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "CreateUser"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "UpdateUser"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "DeleteUser"})
}

func User(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "User"})
}

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Users"})
}
