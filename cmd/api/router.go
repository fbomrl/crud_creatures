package api

import (
	"crud_creatures/internal/services"
	"net/http"
)

func Routers(moveService *services.MoveService, creatureService *services.CreatureService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/creatures", http.StatusFound)
	})

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
