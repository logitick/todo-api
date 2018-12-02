package jsonapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/logitick/todo-api/pkg/adding"
	"github.com/logitick/todo-api/pkg/listing"
)

// Handler wires the http handlers to the services
func Handler(l listing.Service, a adding.Service) http.Handler {

	router := httprouter.New()

	router.GET("/api/todo", getTodos(l))
	router.POST("/api/todo", addTodo(a))
	// router.GET("/todo/:id", getBeer(l))

	return router
}

func getTodos(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetTodos()
		json.NewEncoder(w).Encode(list)
	}
}

func addTodo(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		var todo adding.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = s.AddTodo(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
