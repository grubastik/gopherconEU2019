package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET", "http://localhost", nil,
	)
	w := httptest.NewRecorder()
	HomeHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHomeHandler1(t *testing.T) {
	logger, _ := test.NewNullLogger()
	req := httptest.NewRequest(
		"GET", "http://localhost", nil,
	)
	w := httptest.NewRecorder()
	handler := HomeHandler1(logger)
	handler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
