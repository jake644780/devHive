package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("couldn't load dotenv")
	}
	log.Println("dotenv loaded")

	user := os.Getenv("DB_USER");
	pass := os.Getenv("DB_PASS");
	host := os.Getenv("DB_HOST");
	port := os.Getenv("DB_PORT");
	name := os.Getenv("DB_NAME");

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	user,
	pass,
	host,
	port,
	name);

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});
	if err != nil{
		log.Fatalf("couldn't open db with these:\nuser:\t%s\npass:\t%s\nhost:\t%s\nport:\t%s\nname:\t%s",
		user,
		pass,
		host,
		port,
		name);
	}
	log.Println("connected to db")
	Instance = db
}