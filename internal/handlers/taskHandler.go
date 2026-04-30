package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"projectBit/internal/services"
	"projectBit/models"
)

type TaskHandler interface {
	AddTask(w http.ResponseWriter, r *http.Request)
	ChangeTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(s services.TaskService) TaskHandler {
	return &taskHandler{taskService: s}
}

func (h *taskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Tasks
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println(err)
		return
	}

	createTask, err := h.taskService.AddTask(newTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createTask)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *taskHandler) ChangeTask(w http.ResponseWriter, r *http.Request) {
	var updTask models.Tasks
	err := json.NewDecoder(r.Body).Decode(&updTask)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.taskService.ChangeTask(updTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
