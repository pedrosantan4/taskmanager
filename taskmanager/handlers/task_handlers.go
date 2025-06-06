package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pedrosantan4/taskmanager/database"
	"github.com/pedrosantan4/taskmanager/models"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	respondWithJSON(w, http.StatusOK, tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("⚠️ ID inválido: %v", chi.URLParam(r, "id"))
		respondWithError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	log.Printf("🔍 Buscando task com ID: %d", id)

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		log.Printf("❌ Erro ao buscar task: %v", result.Error)
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	log.Printf("✅ Task encontrada: %+v", task)
	respondWithJSON(w, http.StatusOK, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Garante que o ID do body e o da URL sejam os mesmos
	task.ID = uint(id)

	result := database.DB.Save(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

var validate = validator.New()

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := validate.Struct(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
