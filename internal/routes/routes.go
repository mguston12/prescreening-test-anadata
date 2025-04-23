package routes

import (
	data "Screening-Test-Anagata/go-project/internal/data"
	"Screening-Test-Anagata/go-project/internal/handler"
	"Screening-Test-Anagata/go-project/internal/service"

	"Screening-Test-Anagata/go-project/internal/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	userData := data.NewUserData(db)
	userService := service.NewUserService(userData)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)
	api.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateByID).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.Register).Methods("POST")

	return r
}
