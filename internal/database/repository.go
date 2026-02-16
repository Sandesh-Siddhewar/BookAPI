package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Sandesh-Siddhewar/eBook/BOOKAPI/internal/auth"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		log.Fatal("Unable to load DB User")
	}

	dbpass := os.Getenv("DB_PASS")
	if dbpass == "" {
		log.Fatal("Unable to load DB pass.")
	}

	dbname := os.Getenv("DB_SNAME")
	if dbname == "" {
		log.Fatal("Incorrect schema")
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		log.Fatal("Incorrect host name.")
	}

	dbport := os.Getenv("DB_PORT")
	if dbport == "" {
		log.Fatal("Incorrect Port.")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	connect, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database connected successfully")

	//defer connect.Close()

	err = connect.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Ping completed")

	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&auth.Users{})
	if err != nil {
		panic("Failed to connect:" + err.Error())
	}
	DB = db
}
