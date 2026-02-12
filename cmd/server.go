package cmd

import (
	"footnotez/internal"
	"log"
	"net/http"
)

func main() {
	//initializing the database
	db, err := internal.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	err = internal.CreateTables(db)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	// Create repository, service, and handler
	articleRepo := internal.NewArticleRepository(db)
	articleService := internal.NewArticleService(articleRepo)
	articleHandler := internal.NewArticleHandler(articleService)

	// setup the routes

	mux := http.NewServeMux()
	mux.HandleFunc("/api/article/create", articleHandler.CreateArticle)
	mux.HandleFunc("/api/article/", articleHandler.GetArticle)

	//starting the server on port 8080

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
