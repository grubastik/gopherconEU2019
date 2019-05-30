package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET", "localhost/", nil,
	)
	w := httptest.NewRecorder()
	HomeHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
