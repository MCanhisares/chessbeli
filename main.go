package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)
func main() {
  fmt.Println("hello world")
  
	godotenv.Load()

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT not set in environment")
  }
  fmt.Println("Port: ", portString)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{    
    AllowedOrigins:   []string{"https://*", "http://*"},    
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"*"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, 
  }))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Printf("Server starting on port: %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
