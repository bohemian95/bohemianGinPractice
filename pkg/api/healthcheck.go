package api

import (
	"bohemian/practice/pkg/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ResponseEntity{Result: "Live"})
}
