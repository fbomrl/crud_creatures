package handlers

import (
	service "crud_creatures/internal/services"
	"html/template"
	"net/http"
)

var tempCreatures = template.Must(template.ParseGlob("templates/*.html"))

func CreaturesHandlers(service *service.CreatureService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creatures, err := service.FindAllCreaturesService()
		if err != nil {
			http.Error(w, "Erro ao buscar criaturas", http.StatusInternalServerError)
			return
		}

		if err := tempCreatures.ExecuteTemplate(w, "creatures", creatures); err != nil {
			http.Error(w, "erro ao renderizar template de criaturas", http.StatusInternalServerError)
			return
		}
	}
}
