package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestDB struct {
	instance *gorm.DB
}

func (d *TestDB) Close() {
	db, _ := d.instance.DB()

	db.Close()
}

func NewTestDB(dst ...interface{}) *TestDB {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("I failed to load the .env file")
	}

	dsn := os.Getenv("DSN_TEST")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})

	if err != nil {
		panic("I failed to connect the database")
	}

	db.Migrator().DropTable(dst...)

	db.AutoMigrate(dst...)

	return &TestDB{db}
}

type TestClient struct {
	server *httptest.Server
}

func CreateTestClient(router *gin.Engine) *TestClient {
	return &TestClient{server: httptest.NewServer(router)}
}

func (s *TestClient) Close() {
	s.server.Close()
}

func (s *TestClient) Request(uri string, method string, reqBody interface{}) (*http.Response, error) {
	buf, err := json.Marshal(reqBody)

	if err != nil {
		return &http.Response{}, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", s.server.URL, uri), bytes.NewBuffer(buf))

	if err != nil {
		return &http.Response{}, err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return res, nil
}
