package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/auth"
	"github.com/stretchr/testify/assert"
)

func NewRefreshTestClient(testDB *TestDB) *TestClient {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	auth := &auth.Module{
		Router: router.Group("auth"),
		Db:     testDB.instance,
	}

	auth.AddSignUpUseCase()

	auth.AddSignInUseCase()

	auth.AddRefreshUseCase()

	client := CreateTestClient(router)

	return client
}

func TestRefresh(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewRefreshTestClient(db)

	defer client.Close()

	reqBody1 := gin.H{
		"username": "username",
		"password": "password",
	}

	res1, err := client.Request("auth", http.MethodPost, reqBody1)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusCreated, res1.StatusCode)

	reqBody2 := gin.H{
		"username": "username",
		"password": "password",
	}

	res2, err := client.Request("auth/issue", http.MethodPost, reqBody2)

	if err != nil {
		t.Fatal(err.Error())
	}

	resBody1 := &auth.SignInUseCaseResponseBody{}

	json.NewDecoder(res2.Body).Decode(resBody1)

	assert.Equal(t, http.StatusCreated, res2.StatusCode)

	refreshToken := resBody1.RefreshToken

	reqBody3 := gin.H{
		"refresh_token": refreshToken,
	}

	res3, err := client.Request("auth/refresh", http.MethodPost, reqBody3)

	if err != nil {
		t.Fatal(err.Error())
	}

	resBody2 := &auth.RefreshUseCaseResponseBody{}

	json.NewDecoder(res3.Body).Decode(resBody2)

	assert.Equal(t, http.StatusCreated, res3.StatusCode)

	assert.NotEmpty(t, resBody2.AccessToken)
}
