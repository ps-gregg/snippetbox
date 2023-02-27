package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// sippetview handler
func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

// snippetCreate handler
func snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// the http.NewServeMux func initializes a new servermux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// http.ListenAndServe to start a new web server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
