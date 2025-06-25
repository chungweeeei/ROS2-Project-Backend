package main

import (
	"authenticate-service/cmd/api/docs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func (app *Config) routes() http.Handler {

	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	apiV1 := e.Group("/v1")
	{
		apiV1.POST("/authenticate/login", app.Authenticate)
		apiV1.POST("/authenticate/signup", app.Signup)
	}

	docs.SwaggerInfo.Title = "Authenticate Service API"
	docs.SwaggerInfo.Description = "API for user authentication and management"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:80"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	url := ginSwagger.URL("http://localhost:80/swagger/doc.json") // The url pointing to API definition
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return e
}
