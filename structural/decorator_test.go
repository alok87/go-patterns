package structural

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestDecorateV1(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	d := DecorateV1(hello)
	out := capture.Output(func() {
		d("raj")
	})
	assert.Equal(t, out, "pre post:raj")
}

func TestDecorateV2(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	d := DecorateV2(hello)
	out := capture.Output(func() {
		d("vic")
	})
	assert.Equal(t, out, "pre post:vic")
}

func TestDecorateV3(t *testing.T) {
	hello := func(s string) string {
		return s
	}
	out := capture.Output(func() {
		DecorateV3(hello, "vic_raj8")
	})
	assert.Equal(t, out, "pre post:vic_raj8")
}

// TODO: convert the test to table driven for redability
func TestDecoratorMiddlewares(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// basic test
	rr := httptest.NewRecorder()
	basicHandler := http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			io.WriteString(rw, "basic handler test")
		},
	)
	basicHandler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Error("basic request failed")
	}
	assert.Equal(t, rr.Body.String(), "basic handler test")

	// newrelic handler test
	rr = httptest.NewRecorder()
	newRelicHandler := NewRelicMiddleWare(basicHandler)
	newRelicHandler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Error("newrelic middleware request failed")
	}
	assert.Equal(
		t,
		rr.Body.String(),
		"newrelic_startbasic handler testnewrelic_end",
	)

	// authorization failed test
	rr = httptest.NewRecorder()
	authMiddleware := AuthorizeMiddleware(newRelicHandler)
	authMiddleware.ServeHTTP(rr, req)
	if rr.Code != http.StatusForbidden {
		t.Error("auth middleware request failed")
	}

	// authorization success test
	rr = httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer usertoken")
	authMiddleware.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Error("auth middleware request failed")
	}
	assert.Equal(
		t,
		rr.Body.String(),
		"access_allowednewrelic_startbasic handler testnewrelic_end",
	)

}
