package api

import (
	"crud_creatures/internal/handlers"
	"crud_creatures/internal/services"
	"net/http"
)

func Routers(moveService *services.MoveService, creatureService *services.CreatureService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/creatures", http.StatusFound)
	})

	mux.HandleFunc("/moves", handlers.MovesHandlers(moveService))
	mux.HandleFunc("/creatures", handlers.CreaturesHandlers(creatureService))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
