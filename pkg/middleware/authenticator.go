package middleware

import (
	"bohemian/practice/pkg/structs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authenticator(c *gin.Context) {
	var token = "token"
	var err error

	if err != nil {
		log.Printf("token is unauthorized. (err: %+v)", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, structs.AuthorizationException{Error: "unauthorized."})
	}

	c.Set("token", token)
	c.Next()
}
