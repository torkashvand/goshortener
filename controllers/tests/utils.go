package tests

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torkashvand/goshortener/models"
	"github.com/torkashvand/goshortener/routers"
)

func performGetRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func performPostRequest(r http.Handler, path string, data interface{}) *httptest.ResponseRecorder {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(jsonData))

	w := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	r.ServeHTTP(w, req)

	return w
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str
}

func initApp() (*gin.Engine, *models.SQLite) {
	sqlite := &models.SQLite{}
	sqlite.Open()
	sqlite.AutoMigrate()

	router := routers.SetupRouter(sqlite)

	return router, sqlite
}
