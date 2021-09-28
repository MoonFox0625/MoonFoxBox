// Date:  2021/9/28 22:04
// Desc:
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello MoonFoxBox"))
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Staring server on 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
