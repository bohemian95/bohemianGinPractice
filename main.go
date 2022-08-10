package main

import (
	"bohemian/practice/configs"
	"bohemian/practice/pkg/api"
	"bohemian/practice/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	configs.InitializeConfig(engine)

	engine.GET("/healthCheck", api.HealthCheck)

	router := engine.Group("api")
	{
		routes.CartEndpointRegistration(router)
		routes.ProductEndpointRegistration(router)
		routes.UserEndpointRegistration(router)
	}

	err := engine.Run()

	if err != nil {
		return
	}
}
