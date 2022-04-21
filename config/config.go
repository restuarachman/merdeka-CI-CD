package config

import (
	"fmt"
	"os"
	"restapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	username := os.Getenv("APP_DB_USERNAME")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("APP_DB_PASSWORD")
	if password == "" {
		password = ""
	}
	port := os.Getenv("APP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	host := os.Getenv("APP_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	name := os.Getenv("APP_DB_NAME")
	if name == "" {
		name = "test"
	}
	config := Config{
		DB_Username: username,
		DB_Password: password,
		DB_Port:     port,
		DB_Host:     host,
		DB_Name:     name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
func InitialMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
}
