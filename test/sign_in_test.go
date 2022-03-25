package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/auth"
	"github.com/oh-jinsu/helloworld/modules/users"
	"github.com/stretchr/testify/assert"
)

func NewSignInTestClient(testDB *TestDB) *TestClient {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	users := &users.Module{
		Router: router.Group("users"),
		DB:     testDB.instance,
	}

	users.AddCreateUserUseCase()

	auth := &auth.Module{
		Router: router.Group("auth"),
		Db:     testDB.instance,
	}

	auth.AddSignInUseCase()

	client := CreateTestClient(router)

	return client
}

func TestSignIn(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewSignInTestClient(db)

	defer client.Close()

	reqBody1 := gin.H{
		"username": "username",
		"password": "password",
	}

	res1, err := client.Request("users", http.MethodPost, reqBody1)

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

	resBody := &auth.SignInUseCaseResponseBody{}

	json.NewDecoder(res2.Body).Decode(resBody)

	assert.Equal(t, http.StatusCreated, res2.StatusCode)

	assert.NotNil(t, resBody.AccessToken)

	assert.NotNil(t, resBody.RefreshToken)
}
