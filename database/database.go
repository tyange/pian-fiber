package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
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

	DBConn = db
}
