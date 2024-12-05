package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateClassHandler(t *testing.T) {
	// Mock request body
	body := []byte(`{
		"name": "Pilates",
		"start_date": "2024-12-01",
		"end_date": "2024-12-20",
		"capacity": 10
	}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createClassHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, status)
	}

	// Check the response body
	expected := `{"message":"Class created successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected response %s, got %s", expected, rr.Body.String())
	}
}

func TestBookClassHandler(t *testing.T) {
	// Mock request body
	body := []byte(`{
		"member_name": "John Doe",
		"date": "2024-12-01"
	}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bookClassHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, status)
	}

	// Check the response body
	expected := `{"message":"Booking created successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected response %s, got %s", expected, rr.Body.String())
	}
}
