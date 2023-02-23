package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	godotenv.Load()
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlUserPassword := os.Getenv("MYSQL_PW")
	dbName := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%v:%v@tcp(containers-us-west-195.railway.app:5501)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlUsername, mysqlUserPassword, dbName)

	sqlDB, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Failed to connect to sql DB. \n", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	if db.AutoMigrate(&models.Burger{}, &models.User{}) != nil {
		log.Fatal("Failed DB auto migration.")
	}

	DBConn = db
}
