package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()
	ginEngine := testRoutes.(*gin.Engine)

	routes := []struct {
		URL    string
		Method string
	}{
		{URL: "/v1/records", Method: "GET"},
		{URL: "/v1/record", Method: "POST"},
	}

	for _, route := range routes {
		routeExists(t, ginEngine, route)
	}

	testApp.InfoLog.Println("\033[32m✅ All routes exist as expected\033[0m")
}

func routeExists(t *testing.T, engine *gin.Engine, route struct {
	URL    string
	Method string
}) {
	found := false

	for _, r := range engine.Routes() {
		// Check if the route matches the expected URL and method
		if r.Method == route.Method && r.Path == route.URL {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Did not find %s in registered routes", route)
	}
}
