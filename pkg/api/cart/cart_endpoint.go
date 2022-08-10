package cart

import (
	"bohemian/practice/pkg/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "cart"

func CreateCart(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "CreateCart"})
}

func UpdateCart(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "UpdateCart"})
}

func DeleteCart(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "DeleteCart"})
}

func Cart(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Cart"})
}

func Carts(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Carts"})
}
