package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedrosantan4/taskmanager/database"
	"github.com/pedrosantan4/taskmanager/models"
	"github.com/pedrosantan4/taskmanager/router"
)

func main() {
	// Conectar ao banco
	database.Connect()

	// Migrar modelo Task
	err := database.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Configurar rotas
	r := router.SetupRouter()

	// Iniciar servidor
	log.Println("🚀 Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Could not start server: %v", err)
	}

	newTask := models.Task{
		Title:       "Primeira Task",
		Description: "Conectar ao PostgreSQL",
		Completed:   false,
	}
	database.DB.Create(&newTask)
	fmt.Println("Task criada com ID:", newTask.ID)

	// Busca uma task (READ)
	var task models.Task
	database.DB.First(&task, 1) // Busca a task com ID = 1
	fmt.Println("Task encontrada:", task.Title)

}
