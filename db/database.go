package db

import (
	"fmt"
	"log"
	"os"

	"test/domain/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabaseWithGorm() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みできませんでした: %v", err)
	}

	DBUser := os.Getenv("DBUser")
	DBPass := os.Getenv("DBPass")
	DBProtocol := os.Getenv("DBProtocol")
	DBName := os.Getenv("DBName")
	
	connectTemplate := "%s:%s@%s/%s?%s"
  connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName,"parseTime=true")
		
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Post{})
	return db, nil
}