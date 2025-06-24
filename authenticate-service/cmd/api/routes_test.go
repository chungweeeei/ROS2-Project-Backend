package main

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_routes_exists(t *testing.T) {

	testApp := Config{}
	testRoutes := testApp.routes()
	ginEngine := testRoutes.(*gin.Engine)

	routes := []map[string]string{
		{"URL": "/v1/authenticate/login", "Method": "POST"},
		{"URL": "/v1/authenticate/signup", "Method": "POST"},
	}

	for _, route := range routes {
		routeExists(t, ginEngine, route)
	}

	fmt.Println("All routes exist as expected")
}

func routeExists(t *testing.T, engine *gin.Engine, route map[string]string) {

	found := true

	for _, r := range engine.Routes() {
		if r.Method == route["Method"] && r.Path == route["URL"] {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Did not find %s in registered routes", route)
	}

}
