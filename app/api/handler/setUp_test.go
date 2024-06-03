package handler

import (
	"encoding/json"
	"myapp/crossCutting/api/validator"
	"os"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type httpCallType int

const (
	Get httpCallType = iota
	Put
	Post
	Delete
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func toJSON(any interface{}) string {
	bytes, _ := json.Marshal(any)
	return string(bytes)
}

func getEngine(path string, method httpCallType, handler gin.HandlerFunc) (*gin.Engine, *gofight.RequestConfig) {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	binding.Validator = validator.NewStructValidator()

	switch method {
	case Get:
		engine.GET(path, handler)
	case Post:
		engine.POST(path, handler)
	case Put:
		engine.PUT(path, handler)
	case Delete:
		engine.DELETE(path, handler)
	}
	request := gofight.New()
	return engine, request
}
