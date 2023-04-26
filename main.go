package main

import (
	"fmt"
	"gemm123/grovego-api/db"
	"gemm123/grovego-api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		host, user, pass, dbName)
	_, err = db.InitDB(dsn)
	if err != nil {
		log.Fatal("can't connect to database: " + err.Error())
	}
	fmt.Println("connected to database")

	router := routes.Routes()
	router.Run()
}
