package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MCanhisares/chessbic/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB * database.Queries
}

func main() {
  fmt.Println("hello world")
  
	godotenv.Load()

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT not set in environment")
  }
  fmt.Println("Port: ", portString)
	
	dbUrl := os.Getenv("DB_URL")
  if portString == "" {
		log.Fatal("DB_URL not set in environment")
  }
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to DB", err)
	}

	apiCfg := apiConfig {
		DB: database.New(conn)	,
	}


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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Printf("Server starting on port: %v", portString)

	srvErr := srv.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
