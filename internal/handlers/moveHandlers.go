package handlers

import (
	service "crud_creatures/internal/services"
	"html/template"
	"net/http"
)

var tempMoves = template.Must(template.ParseGlob("templates/*.html"))

func MovesHandlers(service *service.MoveService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		moves, err := service.FindAllMovesService()
		if err != nil {
			http.Error(w, "Erro ao buscar movimentos", http.StatusInternalServerError)
			return
		}

		if err := tempMoves.ExecuteTemplate(w, "moves", moves); err != nil {
			http.Error(w, "erro ao renderizar template de movimentos", http.StatusInternalServerError)
			return
		}
	}
}
