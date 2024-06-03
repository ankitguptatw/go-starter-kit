package securityheaders_test

import (
	"myapp/crossCutting/api/middleware/securityheaders"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	securityheaders.Add()(ctx)

	testcases := []struct {
		header string
		value  string
	}{
		{"Cache-Control", "no-store"},
		{"Content-Security-Policy", "frame-ancestors 'none'"},
		{"Content-Type", "application/json"},
		{"Strict-Transport-Security", "max-age=31536000 ; includeSubDomains"},
		{"X-Frame-Options", "DENY"},
		{"X-XSS-Protection", "1; mode=block"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.header, func(t *testing.T) {
			assert.Equal(t, testcase.value, w.Header().Get(testcase.header))
		})
	}

}
