package jsonapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/logitick/todo-api/pkg/adding"
	"github.com/logitick/todo-api/pkg/listing"
	"github.com/logitick/todo-api/pkg/updating"
)

// Handler wires the http handlers to the services
func Handler(l listing.Service, a adding.Service, u updating.Service) http.Handler {

	router := httprouter.New()

	router.GET("/api/todo", getTodos(l))
	router.POST("/api/todo", addTodo(a))
	router.PATCH("/api/todo/:id", updateTodo(u))

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

func updateTodo(s updating.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		var todo updating.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := p.ByName("id")
		todo.ID, err = strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := s.UpdateTodo(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(t)
	}
}
