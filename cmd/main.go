package main // package declaration

import (
	"fmt"      // import fmt package
	"log"      // import log package
	"net/http" // import http package
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, Go internet monitor!") })

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
