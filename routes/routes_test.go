package routes_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
    "url-shortener/routes"
)

// setup function initializes the Gin engine with the routes
func setup() *gin.Engine {
    return routes.SetupRouter()
}

// TestRegister tests the /register endpoint
func TestRegister(t *testing.T) {
    router := setup()

    w := httptest.NewRecorder()
    jsonStr := []byte(`{"email": "test@example.com", "test@example.com": "test@example.com"}`)
    req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "User registered successfully", response["message"])
}

// TestLogin tests the /login endpoint
func TestLogin(t *testing.T) {
    router := setup()

    w := httptest.NewRecorder()
    jsonStr := []byte(`{"email": "test@example.com", "test@example.com": "test@example.com"}`)
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.NotNil(t, response["token"])
}

// TestRedirectURL tests the /:short_url endpoint
func TestRedirectURL(t *testing.T) {
    router := setup()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/short_url", nil)

    router.ServeHTTP(w, req)

    assert.Equal(t, 302, w.Code)
}

// TestCreateShortURL tests the /shorten endpoint
func TestCreateShortURL(t *testing.T) {
    router := setup()

    // Simulate login to get token
    loginResponse := loginUser(t, router)

    w := httptest.NewRecorder()
    jsonStr := []byte(`{"long_url": "http://example.com"}`)
    req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", "Bearer "+loginResponse["token"].(string))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.NotNil(t, response["short_url"])
}

// TestGetUserURLs tests the /urls endpoint
func TestGetUserURLs(t *testing.T) {
    router := setup()

    // Simulate login to get token
    loginResponse := loginUser(t, router)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/urls", nil)
    req.Header.Set("Authorization", "Bearer "+loginResponse["token"].(string))

    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    var response []map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.GreaterOrEqual(t, len(response), 0)
}

// loginUser is a helper function to simulate user login and retrieve token
func loginUser(t *testing.T, router *gin.Engine) map[string]interface{} {
    w := httptest.NewRecorder()
    jsonStr := []byte(`{"email": "test@example.com", "test@example.com": "test@example.com"}`)
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    return response
}
