// Date:  2021/9/28 22:04
// Desc:
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home : Displaying the home page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("Hello MoonFoxBox"))
}

// showSnippet : Display a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID: %d", id)
}

// createSnippet:Create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// Suppressing System-Generated Headers
		// w.Header()["Date"] = nil

		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method Not Allowed"))
		// 上面可以简化于下面
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

// main function
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Staring server on 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
