package main

import (
	"log"
	"net/http"
	"time"

	"gopractice/mini_proj_A/db"
	"gopractice/mini_proj_A/handlers"
	"gopractice/mini_proj_A/repository"
)

// loggingMiddleware logs HTTP method, path, and duration for each request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s took %v\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func main() {
	// 1. Initialize SQLite Database
	dbPath := "mini_proj_A/users.db"
	database, err := db.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer database.Close()

	// 2. Wire Dependencies (Repo -> Handler)
	userRepo := repository.NewUserRepository(database)
	userHandler := handlers.NewUserHandler(userRepo)

	// 3. Setup Go 1.22+ Standard ServeMux
	mux := http.NewServeMux()

	// Route definitions with HTTP method prefixes
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Wrap mux with logging middleware
	handlerWithMiddleware := loggingMiddleware(mux)

	// 4. Start Server
	port := ":8080"
	log.Printf("User CRUD Server running on http://localhost%s\n", port)
	log.Println("Available endpoints:")
	log.Println("  POST   /users       -> Create a user")
	log.Println("  GET    /users       -> List all users")
	log.Println("  GET    /users/{id}  -> Get user by ID")
	log.Println("  PUT    /users/{id}  -> Update user by ID")
	log.Println("  DELETE /users/{id}  -> Delete user by ID")

	if err := http.ListenAndServe(port, handlerWithMiddleware); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}
