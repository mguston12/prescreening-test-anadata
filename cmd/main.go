package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"Screening-Test-Anagata/go-project/internal/config"
	"Screening-Test-Anagata/go-project/internal/entity"
	"Screening-Test-Anagata/go-project/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	db := config.InitDB()

	// ➕ Tambahkan baris ini untuk membuat tabel `users`
	db.AutoMigrate(&entity.User{})

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := routes.NewRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
