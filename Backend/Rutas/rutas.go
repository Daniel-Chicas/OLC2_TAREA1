package Rutas

import (
	"OLC2_TAREA1/Controladores"
	"github.com/gorilla/mux"
)

func RutasCompiladores(router *mux.Router) {
	router.HandleFunc("/", Controladores.Inicio()).Methods("GET")
	router.HandleFunc("/Parser", Controladores.Data()).Methods("POST")
	router.HandleFunc("/Errores", Controladores.ErroresEntrada()).Methods("POST")
}
