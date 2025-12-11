package main

import (
	"log"
	"os"

	"crud_creatures/internal/database"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || name == "" {
		log.Fatal("missing DB env vars")
	}

	if err := database.EnsureDBAndMigrate(user, pass, host, name, "migration"); err != nil {
		log.Fatal(err)
	}

	log.Println("migrations applied")
}
