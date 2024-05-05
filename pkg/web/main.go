package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johnoppenheimer/boustifaille/web/pages"
	"github.com/johnoppenheimer/boustifaille/web/point"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := pages.Index(point.Points)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//

	getHandler(w, r)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}

		getHandler(w, r)
	})

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error listening %v", err)
	}
}
