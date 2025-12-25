package handlers

import (
	"encoding/json"
	"geobill_golang_versions/service"
	"net/http"
)

type Handler struct {
	Service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{Service: s}
}

type RegistrationPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegistrationPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate Role (Optional but good practice)
	// For now validation happens in service or DB constraints or trusting user input for this demo

	err := h.Service.Register(payload.Username, payload.Password, payload.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// For login endpoint, we just verify credentials and return success message
	// Since the requirement is Basic Auth for protection, this endpoint is just a check.
	// Or we could return the user info.

	user, err := h.Service.Login(payload.Username, payload.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user":    user,
	})
}
