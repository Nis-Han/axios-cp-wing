package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func makeRequest(method, url string, body interface{}, headers map[string]string) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	writer := httptest.NewRecorder()
	router := SetupRoutes()
	router.ServeHTTP(writer, request)
	return writer
}
