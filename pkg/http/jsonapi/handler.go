package jsonapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/logitick/todo-api/pkg/listing"
)

// Handler wires the http handlers to the services
func Handler(l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/todo", getTodos(l))
	// router.GET("/todo/:id", getBeer(l))

	// router.POST("/todo", addBeer(a))

	return router
}

func getTodos(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetTodos()
		json.NewEncoder(w).Encode(list)
	}
}
