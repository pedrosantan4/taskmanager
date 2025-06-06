package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pedrosantan4/taskmanager/handlers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rota de health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Task Manager API"))
	})

	// Rotas para Tasks
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetAllTasks) // GET /tasks - Lista todas
		r.Post("/", handlers.CreateTask) // POST /tasks - Cria nova

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetTaskByID)   // GET /tasks/{id} - Busca por ID
			r.Put("/", handlers.UpdateTask)    // PUT /tasks/{id} - Atualiza
			r.Delete("/", handlers.DeleteTask) // DELETE /tasks/{id} - Remove
		})
	})

	return r
}
