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

	//starting the server on port 8080

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
