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

func NewSignInTestClient(testDB *TestDB) *TestClient {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	auth := &auth.Module{
		Router: router.Group("auth"),
		Db:     testDB.instance,
	}

	auth.AddSignUpUseCase()

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

	resBody := &auth.SignInUseCaseResponseBody{}

	json.NewDecoder(res2.Body).Decode(resBody)

	assert.Equal(t, http.StatusCreated, res2.StatusCode)

	assert.NotEmpty(t, resBody.AccessToken)

	assert.NotEmpty(t, resBody.RefreshToken)
}

func TestSignInWithWrongUsername(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewSignInTestClient(db)

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
		"username": "user",
		"password": "password",
	}

	res2, err := client.Request("auth/issue", http.MethodPost, reqBody2)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusNotFound, res2.StatusCode)
}

func TestSignInWithWrongPassword(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewSignInTestClient(db)

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
		"password": "wrongpass",
	}

	res2, err := client.Request("auth/issue", http.MethodPost, reqBody2)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusUnauthorized, res2.StatusCode)
}
