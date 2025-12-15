package main

import (
	"log"
	"net/http"
	"os"

	"crud_creatures/cmd/api"
	"crud_creatures/internal/database"
	"crud_creatures/internal/repository"
	"crud_creatures/internal/services"
)

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	instance := os.Getenv("DB_INSTANCE")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || name == "" {
		log.Fatal("missing DB env vars")
	}

	dsn := database.BuildDSN(user, pass, host, instance, name)

	db, err := database.SqlServer(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	moveRepo := &repository.MovesRepository{DB: db}
	creatureRepo := &repository.CreatureRepository{DB: db}

	moveService := &services.MoveService{RepoMove: moveRepo}
	creatureService := &services.CreatureService{RepoCreature: creatureRepo}

	handler := api.Routers(moveService, creatureService)

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "localhost.pem", "localhost-key.pem", handler))
}
