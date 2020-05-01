package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	Mux := mux.NewRouter()
	Mux.HandleFunc("/", mainPageHandler)
	Mux.HandleFunc("/mainPage.css", mainPageHandler1)
	Mux.HandleFunc("/about", aboutHandler)
	Mux.HandleFunc("/mainPage.js", mainPagejsHandler)
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      Mux,
	}
	fmt.Println(port)
	srv.ListenAndServe()
}
