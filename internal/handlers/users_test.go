package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	fmt.Println(recorder.Code)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func signUp(user models.SignUpInput) (models.User, error) {
	r := SetupRouters()
	recorder := httptest.NewRecorder()
	url := "/user/signUp"
	testSignUp := models.SignUpInput{
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
	}
	data, err := json.Marshal(testSignUp)
	if err != nil {
		return models.User{}, err
	}
	bodyReader := bytes.NewReader(data)
	request, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return models.User{}, err
	}
	var userDB models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &userDB)
	if err != nil {
		return models.User{}, err
	}
	r.ServeHTTP(recorder, request)
	return userDB, nil
}
