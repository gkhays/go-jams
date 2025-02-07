package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gkh/fips"
)

// User represents a user structure for JSON payload
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FIPSConfig represents FIPS configuration information
type FIPSConfig struct {
	FIPSMode bool `json:"fips_mode"`
}

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// Handler for GET endpoint
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	logger.Info("User info",
		slog.Int("user-id", user.ID),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler for FIPS mode endpoint
func getFIPSModeHandler(w http.ResponseWriter, r *http.Request) {
	fipsMode := fips.IsFIPSModeEnabled()

	logger.Info("FIPS mode",
		slog.Bool("fips-enabled", fipsMode),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(FIPSConfig{FIPSMode: fipsMode})
}

// Handler for POST endpoint
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode incoming JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log received user (in real app, you'd save to database)
	log.Printf("Received user: %+v", user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func main() {
	logger.Info("Starting server...", slog.String("port", ":8080"))

	// User endpoints
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUserHandler(w, r)
		case http.MethodPost:
			createUserHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// FIPS mode endpoint
	http.HandleFunc("/fips-mode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		getFIPSModeHandler(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error("REST service failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
