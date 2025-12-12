package main

import (
	"crud_creatures/internal/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	instance := os.Getenv("DB_INSTANCE")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || name == "" {
		log.Fatalf("missing DB env vars (USER=%s HOST=%s NAME=%s)", user, host, name)
	}

	log.Printf("➡️ Connecting with:\nHOST=%s\nINSTANCE=%s\nDB=%s\n", host, instance, name)

	if err := database.EnsureDBAndMigrate(
		user,
		pass,
		host,
		instance,
		name,
		"migration",
	); err != nil {
		log.Fatal(err)
	}

	log.Println("migratrion aplicada com sucesso!")
}
