package main

import (
	"OLC2_TAREA1/Gramatica"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {
	slide, err := fmt.Fprintf(w, "The server is active on port: 3000")
	if err != nil {
		return
	}
	fmt.Println(slide)
}

func main() {
	//router := mux.NewRouter()
	//router.HandleFunc("/", Init).Methods("GET")
	//log.Fatal(http.ListenAndServe(":3000", router))

	is := antlr.NewInputStream("\"hola\"+23")
	lexer := Gramatica.NewGramaticaLexer(is)
	fmt.Println(lexer)

}
