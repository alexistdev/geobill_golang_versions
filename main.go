package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"geobill_golang_versions/handlers"
	"geobill_golang_versions/middleware"
	"geobill_golang_versions/models"
	"geobill_golang_versions/repository"
	"geobill_golang_versions/service"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database Connection
	// user:password@tcp(host:port)/dbname
	dsn := "root:325339@tcp(localhost:3306)/mybill?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to MySQL database")

	// Init Layers
	repo := repository.NewMySQLRepository(db)
	authService := service.NewAuthService(repo)
	authMiddleware := middleware.NewMiddleware(authService)
	authHandler := handlers.NewHandler(authService)

	// Router Setup
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Public Routes
	r.Post("/v1/auth/login", authHandler.Login)
	r.Post("/v1/auth/registration", authHandler.Register)

	// Protected Routes
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.BasicAuthMiddleware)

		// Administrator
		r.With(middleware.CheckRole(models.RoleAdmin)).Get("/v1/admin/dashboard", dashboardHandler("Administrator"))

		// Staff
		r.With(middleware.CheckRole(models.RoleStaff)).Get("/v1/staff/dashboard", dashboardHandler("Staff"))

		// User
		r.With(middleware.CheckRole(models.RoleUser)).Get("/v1/user/dashboard", dashboardHandler("User"))
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func dashboardHandler(roleName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		response := map[string]string{
			"message": fmt.Sprintf("Welcome to the %s Dashboard", roleName),
			"user":    user.Username,
			"role":    user.Role,
			"time":    time.Now().Format(time.RFC3339),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
