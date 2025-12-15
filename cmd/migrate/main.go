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
	instance := os.Getenv("DB_INSTANCE")
	name := os.Getenv("DB_NAME")

	    dsn := /* montar DSN conforme internal/database.SqlServer espera, ex: sqlserver://... */


	if user == "" || pass == "" || host == "" || name == "" {
		log.Fatalf("missing DB env vars (USER=%s HOST=%s NAME=%s)", user, host, name)
	}

	log.Printf("aplicando migrations usando:\nHOST=%s\nINSTANCE=%s\nDB=%s\n",
		host, instance, name,
	)

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

	log.Println("migration aplicada com sucesso!")
}
