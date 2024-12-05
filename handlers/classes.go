package handlers

import (
	"encoding/json"
	"net/http"
	"glofox-api/models"
	"glofox-api/storage"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var newClass models.Class
	if err := json.NewDecoder(r.Body).Decode(&newClass); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !newClass.IsValid() {
		http.Error(w, "Invalid class details", http.StatusBadRequest)
		return
	}

	storage.AddClass(newClass)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Class created successfully"})
}
