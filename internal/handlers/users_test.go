package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/mjaliz/gotracktime/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBRepo_SignIn(t *testing.T) {
	r := SetupRouters()
	recorder := httptest.NewRecorder()
	url := "/user/signIn"
	testSignIn := models.SignUpInput{
		Email:    "mrph14@gmial.com",
		Password: "12345",
	}
	data, err := json.Marshal(testSignIn)
	if err != nil {
		t.Error(err)
	}
	bodyReader := bytes.NewReader(data)
	request, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
