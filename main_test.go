package main

import (
	"bytes"
	"encoding/json"
	"gin-fleamarket/dto"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"gin-fleamarket/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	db.Exec("DELETE FROM items")
	db.Exec("DELETE FROM users")
	items := []models.Item{
		{Name: "テストアイテム１", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム２", Price: 2000, Description: "テスト２", SoldOut: true, UserID: 1},
		{Name: "テストアイテム３", Price: 3000, Description: "テスト３", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	//テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	//APIリクエストの実行
	router.ServeHTTP(w, req)

	//APIの実行結果を取得
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	//アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	//テストのセットアップ
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム４",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	//APIリクエストの実行
	router.ServeHTTP(w, req)

	//APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	//アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	// assert.Equal(t, uint(4), res["data"].ID)
	assert.Equal(t, "テストアイテム４", res["data"].Name)
}

func TestCreateUnauthorized(t *testing.T) {
	//テストのセットアップ
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム４",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	//APIリクエストの実行
	router.ServeHTTP(w, req)

	//APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	//アサーション
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
