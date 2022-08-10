package product

import (
	"bohemian/practice/pkg/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "product"

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "CreateProduct"})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "UpdateProduct"})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "DeleteProduct"})
}

func Product(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Product"})
}

func Products(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Products"})
}
