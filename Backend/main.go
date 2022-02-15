package main

import (
	"OLC2_TAREA1/Rutas"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	Rutas.RutasCompiladores(r)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	srv := &http.Server{
		Handler:      corsWrapper.Handler(r),
		Addr:         ":3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
