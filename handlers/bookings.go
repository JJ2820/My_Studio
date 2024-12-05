package handlers

import (
	"encoding/json"
	"net/http"
	"glofox-api/models"
	"glofox-api/storage"
)

func BookClassHandler(w http.ResponseWriter, r *http.Request) {
	var newBooking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !newBooking.IsValid() {
		http.Error(w, "Invalid booking details", http.StatusBadRequest)
		return
	}

	storage.AddBooking(newBooking)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking created successfully"})
}
