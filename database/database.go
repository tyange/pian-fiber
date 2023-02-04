package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	mysqlStorage "github.com/gofiber/storage/mysql"
	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn       *gorm.DB
	SessionStore *session.Store
)

func ConnectDb() {
	godotenv.Load()
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlUserPassword := os.Getenv("MYSQL_PW")
	dbName := "test_burger4"

	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlUsername, mysqlUserPassword, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	if db.AutoMigrate(&models.Burger{}, &models.User{}) != nil {
		log.Fatal("Failed DB auto migration.")
	}

	DBforSql, _ := db.DB()

	store := mysqlStorage.New(mysqlStorage.Config{
		Db:         DBforSql,
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	SessionStore = session.New(session.Config{
		Storage: store,
	})

	DBConn = db
}
