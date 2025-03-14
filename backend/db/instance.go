package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	err := godotenv.Load()
	if err != nil{
		return &gorm.DB{}, errors.New("couldn't read dotenv file");
	}

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
		errMSG := fmt.Sprintf("couldn't open db with these:\nusername:\t%s\npassword:\t%s\nhost:\t%s\nport:\t%s\ndb name:\t%s",
		user,
		pass,
		host,
		port,
		name);
		return &gorm.DB{}, errors.New(errMSG);
	}
	return db, nil;
}