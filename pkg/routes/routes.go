package routes

import (
	"bohemian/practice/pkg/api/cart"
	"bohemian/practice/pkg/api/product"
	"bohemian/practice/pkg/api/user"
	"github.com/gin-gonic/gin"
)

func CartEndpointRegistration(router *gin.RouterGroup) {
	endpoint := router.Group(cart.Path)
	{
		endpoint.GET("/:id", cart.Cart)
		endpoint.GET("", cart.Carts)
		endpoint.POST("/create", cart.CreateCart)
		endpoint.POST("/update", cart.UpdateCart)
		endpoint.POST("/delete", cart.DeleteCart)
	}
}

func ProductEndpointRegistration(router *gin.RouterGroup) {
	endpoint := router.Group(product.Path)
	{
		endpoint.GET("/:id", product.Product)
		endpoint.GET("", product.Products)
		endpoint.POST("/create", product.CreateProduct)
		endpoint.POST("/update", product.UpdateProduct)
		endpoint.POST("/delete", product.DeleteProduct)
	}
}

func UserEndpointRegistration(router *gin.RouterGroup) {
	endpoint := router.Group(user.Path)
	{
		endpoint.GET("/:id", user.User)
		endpoint.GET("", user.Users)
		endpoint.POST("/create", user.CreateUser)
		endpoint.POST("/update", user.UpdateUser)
		endpoint.POST("/delete", user.DeleteUser)
	}
}
