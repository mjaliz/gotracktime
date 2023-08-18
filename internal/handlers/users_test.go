package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/mjaliz/gotracktime/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBRepo_SignIn(t *testing.T) {
	testCases := []struct {
		name               string
		input              models.SignInInput
		expectedStatusCode int
		errorExpected      bool
	}{
		{
			"OK",
			models.SignInInput{
				Email:    "mrph1@gmail.com",
				Password: "12345",
			},
			http.StatusOK,
			false,
		},
		{
			"BadRequest-Without Email",
			models.SignInInput{
				Password: "12345",
			},
			http.StatusBadRequest,
			true,
		},
		{
			"BadRequest-Without Password",
			models.SignInInput{
				Email: "test@gmail.com",
			},
			http.StatusBadRequest,
			true,
		},
		{
			"Unauthorized",
			models.SignInInput{
				Email:    "unauthorized@gmail.com",
				Password: "12345",
			},
			http.StatusUnauthorized,
			true,
		},
		{
			"DBError",
			models.SignInInput{
				Email:    "internalServerError@gmail.com",
				Password: "12345",
			},
			http.StatusInternalServerError,
			true,
		},
		{
			"Hash password error",
			models.SignInInput{
				Email:    "hashPasswordError@gmail.com",
				Password: "12345",
			},
			http.StatusUnauthorized,
			true,
		},
	}
	url := "/user/signIn"
	r := setUpTestRoutes()
	for _, tc := range testCases {
		recorder := httptest.NewRecorder()
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Error(err)
		}
		bodyReader := bytes.NewReader(data)
		request, err := http.NewRequest(http.MethodPost, url, bodyReader)
		if err != nil && !tc.errorExpected {
			t.Error(err)
		}
		r.ServeHTTP(recorder, request)
		assert.Equal(t, tc.expectedStatusCode, recorder.Code)
	}
}

func signUp(user models.SignUpInput) (models.User, error) {
	r := setUpTestRoutes()
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

func setUpTestRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := NewRouters()
	return r
}
