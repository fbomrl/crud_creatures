package api

import (
	"crud_creatures/internal/handlers"
	"crud_creatures/internal/services"
	"html/template"
	"net/http"
)

func Routers(moveService *services.MoveService, creatureService *services.CreatureService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// se houver criaturas, encaminha para /creatures; caso contrário mostra página inicial
		if _, err := creatureService.FindAllCreaturesService(); err == nil {
			http.Redirect(w, r, "/creatures", http.StatusFound)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		if err := tmpl.ExecuteTemplate(w, "index", nil); err != nil {
			http.Error(w, "erro ao renderizar página inicial", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/moves", handlers.MovesHandlers(moveService))
	mux.HandleFunc("/creatures", handlers.CreaturesHandlers(creatureService))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
