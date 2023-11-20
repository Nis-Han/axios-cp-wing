package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/handlers/handlers"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func makeRequest(method, url string, body interface{}, headers map[string]string, db database.Querier) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	writer := httptest.NewRecorder()
	apiHandler := handlers.Api{DB: db}
	router := handlers.SetupRoutes(&apiHandler)
	router.ServeHTTP(writer, request)
	return writer
}
