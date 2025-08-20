package main

import (
	"net/http"
	"record-service/proto/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func (app *Config) Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header not provided",
			})
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header format"})
			return
		}

		token := headerParts[1]

		resp, err := app.AuthClient.CheckAuthenticate(c.Request.Context(), &auth.AuthenticateRequest{
			Token: token,
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !resp.IsAuthenticated {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Set("Email", resp.Email)

		c.Next()
	}
}
