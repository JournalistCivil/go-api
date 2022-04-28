package internal

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Environment struct {
	AppPort string
	AppEnv  string
	DbUser  string
	DbPass  string
	DbName  string
	DbHost  string
	DbPort  string
}

func checkEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file")
	}

	envs := []string{
		"APP_PORT",
		"APP_ENV",
		"DB_USER",
		"DB_PASS",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
	}

	for _, env := range envs {
		value := os.Getenv(env)
		if value == "" {
			msg := "env variable missing: " + env
			log.Fatal(msg)
		}
	}
}

func GetEnvironment() Environment {
	checkEnvironmentVariables()

	return Environment{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		DbUser:  os.Getenv("DB_USER"),
		DbPass:  os.Getenv("DB_PASS"),
		DbName:  os.Getenv("DB_NAME"),
		DbHost:  os.Getenv("DB_HOST"),
		DbPort:  os.Getenv("DB_PORT"),
	}
}

func GetDbClient(e *Environment) *gorm.DB {
	username := e.DbUser
	password := e.DbPass
	dbName := e.DbName
	dbHost := e.DbHost
	dbPort := e.DbPort

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("unexpected db error, when connecting to db")
	}
	// See "Important settings" section.
	sqlDB, err := client.DB()
	if err != nil {
		log.Fatal("unexpected db error, when trying to set connection pool")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return client
}

func GetDB() *gorm.DB {
	return db
}
