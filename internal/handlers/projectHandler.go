package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"projectBit/internal/services"
	"projectBit/models"
)

type ProjectHandler interface {
	AddProject(w http.ResponseWriter, r *http.Request)
}

type projectHandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) ProjectHandler {
	return &projectHandler{projectService: projectService}
}

func (h *projectHandler) AddProject(w http.ResponseWriter, r *http.Request) {
	var newProject models.Projects
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		log.Println(err)
		return
	}

	createdProject, err := h.projectService.AddProject(newProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdProject)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
