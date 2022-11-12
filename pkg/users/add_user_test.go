package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"harold29/yourkeyswallet/pkg/common/config"
	"harold29/yourkeyswallet/pkg/common/db"
	"harold29/yourkeyswallet/pkg/common/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Test_AddUser_OK(t *testing.T) {
	a := assert.New(t)

	database, err := setupDb()
	if err != nil {
		a.Error(err)
	}

	dateParsed, err := time.Parse("2006-01-02", "1980-01-01")
	if err != nil {
		a.Error(err)
	}

	user := models.User{
		FirstName:    "Test",
		LastName:     "Test",
		Email:        "test@test.com",
		PhoneNumber1: "+541111111111",
		PhoneNumber2: "+542222222222",
		Gender:       "potato",
		Birthday:     dateParsed,
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setAddUserRouter(database, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := user

	a.Equal(expected.FirstName, actual.FirstName)
	a.Equal(expected.LastName, actual.LastName)
	a.Equal(expected.Email, actual.Email)
	a.Equal(expected.PhoneNumber1, actual.PhoneNumber1)
	a.Equal(expected.PhoneNumber2, actual.PhoneNumber2)
	a.Equal(expected.Gender, actual.Gender)
	a.Equal(expected.Birthday, actual.Birthday)

	clearAndClose()
}

func Test_AddUser_MissingInformation(t *testing.T) {
	a := assert.New(t)

	database, err := setupDb()
	if err != nil {
		a.Error(err)
	}

	dateParsed, err := time.Parse("2006-01-02", "1980-01-01")
	if err != nil {
		a.Error(err)
	}

	user := models.User{
		LastName:     "Test",
		Email:        "test@test.com",
		PhoneNumber1: "+541111111111",
		PhoneNumber2: "+542222222222",
		Gender:       "potato",
		Birthday:     dateParsed,
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setAddUserRouter(database, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")
	clearAndClose()
}

func Test_AddUser_DuplicatedRecord(t *testing.T) {
	a := assert.New(t)

	database, err := setupDb()
	if err != nil {
		a.Error(err)
	}

	dateParsed, err := time.Parse("2006-01-02", "1980-01-01")
	if err != nil {
		a.Error(err)
	}

	user1 := models.User{
		FirstName:    "Potato1",
		LastName:     "Test1",
		Email:        "test@test.com",
		PhoneNumber1: "+541111111111",
		PhoneNumber2: "+542222222222",
		Gender:       "potato",
		Birthday:     dateParsed,
	}

	user2 := models.User{
		FirstName:    "Potato1",
		LastName:     "Test1",
		Email:        "test@test.com",
		PhoneNumber1: "+541111111111",
		PhoneNumber2: "+542222222222",
		Gender:       "potato",
		Birthday:     dateParsed,
	}

	reqBody1, err := json.Marshal(user1)
	if err != nil {
		a.Error(err)
	}

	reqBody2, err := json.Marshal(user2)
	if err != nil {
		a.Error(err)
	}

	req1, w1, err := setAddUserRouter(database, bytes.NewBuffer(reqBody1))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req1.Method, "HTTP request method error")
	a.Equal(http.StatusCreated, w1.Code, "HTTP request should return created")

	req2, w2, err := setAddUserRouter(database, bytes.NewBuffer(reqBody2))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req2.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, w2.Code, "HTTP request should return error")

	clearAndClose()
}

func setAddUserRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	h := &handler{DB: db}

	r.POST("/", h.AddUser)

	req, err := http.NewRequest(http.MethodPost, "/", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func setupDb() (*gorm.DB, error) {
	viper.SetConfigFile("../common/envs/dev.yaml")

	conf, errConf := config.LoadConfig("../common/envs/")

	if errConf != nil {
		fmt.Printf("Error loading configuration, %s", errConf)
		return nil, errConf
	}

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "127.0.0.1", conf.Db.PostgresUser, conf.Db.PostgresPass, conf.Db.PostgresDB, conf.Db.PostgresPort)

	return db.Init(dbInfo), nil
}

func clearAndClose() {
	db.ClearTable()
}
