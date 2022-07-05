package database

import (
	"book-server/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect(db *gorm.DB) *gorm.DB {
	//config, err := godotenv.Read()
	//if err != nil {
	//	log.Fatal("Error reading .env file")
	//}

	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	fmt.Println(dataSource)

	connect, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error", err)
	}

	return connect
}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&domain.Book{})
	return db
}
