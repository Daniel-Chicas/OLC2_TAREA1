package Controladores

import (
	parser2 "OLC2_TAREA1/Gramatica"
	"OLC2_TAREA1/Personal"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"net/http"
)

type Entrada struct {
	Texto string `json:"text"`
}

func Inicio() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]interface{}{"ok": "Exitoso"})
	}
}

func Data() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var solicitu Entrada

		if err := json.NewDecoder(r.Body).Decode(&solicitu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Status": http.StatusBadRequest, "Message": "Error"})
			return
		}

		errores := &Personal.CustomErrorListener{}

		is := antlr.NewInputStream(solicitu.Texto)

		// Creación de lexer
		lexer := parser2.NewGramaticaLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errores)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Creando el parser
		p := parser2.NewGramaticaParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(errores)

		p.BuildParseTrees = true

		// Obteniendo la raiz
		tree := p.Ini()

		//fmt.Printf("\nErrores este punto %v", errores.Errors)
		//fmt.Println()
		// Listener para recorrer el arbol
		var listener *Personal.TreeShapeListener = Personal.NewTreeShapeListener()

		if len(errores.Errors) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}
		fmt.Println("------------------------------------------")
		fmt.Println(listener.Cadena)
		fmt.Println("------------------------------------------")
		p.RemoveErrorListeners()
		json.NewEncoder(w).Encode(map[string]interface{}{"val": listener.Cadena})
	}
}
