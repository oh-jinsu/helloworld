package test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/users"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("I failed to load the .env file")
	}

	code := t.Run()

	os.Exit(code)
}

func NewCreateUserTestClient(testDB *TestDB) *TestClient {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	module := &users.Module{
		Router: router.Group(""),
		DB:     testDB.instance,
	}

	module.AddCreateUserUseCase()

	client := CreateTestClient(router)

	return client
}

func TestCreateUser(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "username",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	resBody := &users.CreateUserResponseBody{}

	json.NewDecoder(res.Body).Decode(resBody)

	assert.Equal(t, http.StatusCreated, res.StatusCode)

	assert.NotNil(t, resBody.Id)

	assert.NotNil(t, resBody.CreatedAt)

	assert.Equal(t, "username", resBody.Username)
}

func TestCreateUserWithTheSameName(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "username",
		"password": "password",
	}

	res1, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusCreated, res1.StatusCode)

	res2, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusConflict, res2.StatusCode)
}

func TestCreateUserWithUsernameIncludingKoreanCharacter(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "ㄱㄴㄷㄹ",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithUsernameIncludingSpecialCharacter(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "유저이름-",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithUsernameIncludingSpace(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "user name",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooShortEnglishUsername(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "us",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooShortKoreanUsername(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "유",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooLongEnglishUsername(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "usernameusernameu",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooLongKoreanUsername(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "유저이름유저이름유",
		"password": "password",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooShortPassword(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "username",
		"password": "passwor",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithTooLongPassword(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "username",
		"password": "passwordpasswordpasswordp",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithPasswordIncludingSpaceCharacter(t *testing.T) {
	db := NewTestDB(&models.User{})

	defer db.Close()

	client := NewCreateUserTestClient(db)

	defer client.Close()

	reqBody := gin.H{
		"username": "username",
		"password": "pass word",
	}

	res, err := client.Request("/", http.MethodPost, reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
