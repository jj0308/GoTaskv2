package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jj0308/GoTaskv2/config"
	"github.com/jj0308/GoTaskv2/storage"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	cfg := config.LoadConfig()

	db := storage.SetupDatabase(cfg)
	defer db.Close()

	// Create a new router using the SetupRouter function
	router := SetupRouter(db)

	// Define test cases as a list of anonymous structs
	tests := []struct {
    method string
    path   string
    code   int
	}{
		{http.MethodPost, "/users/:id/events", http.StatusBadRequest},
		{http.MethodPost, "/meetings", http.StatusBadRequest},
		{http.MethodPut, "/invitations/:id", http.StatusBadRequest},
		{http.MethodGet, "/users/:id/invitations", http.StatusInternalServerError},
		{http.MethodGet, "/users/:id/meetings", http.StatusInternalServerError},
	}

	// Iterate over test cases and execute each
	for _, test := range tests {
		t.Run(test.method+" "+test.path, func(t *testing.T) {
			// Create a new request with the test method and path
			req, _ := http.NewRequest(test.method, test.path, nil)
			resp := httptest.NewRecorder()

			// Process the request using the router
			router.ServeHTTP(resp, req)

			// Assert that the response code matches the expected code
			assert.Equal(t, test.code, resp.Code)
		})
	}
}
