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
		{"URL": "/v1/logs", "Method": "GET"},
		{"URL": "/v1/log", "Method": "POST"},
	}

	for _, route := range routes {
		routeExists(t, ginEngine, route)
	}

	fmt.Printf("\033[32m✅ All routes exist as expected\033[0m\n")
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
