package Controladores

import (
	"OLC2_TAREA1/AST/Funcion"
	"OLC2_TAREA1/ErroresSemanticos"
	parser2 "OLC2_TAREA1/Gramatica"
	"OLC2_TAREA1/Personal"
	"encoding/json"
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

		erroresLexicos := &Personal.CustomErrorListener{}
		erroresSintacticos := &Personal.CustomErrorListener{}

		is := antlr.NewInputStream(solicitu.Texto)

		// Creación de lexer
		lexer := parser2.NewGramaticaLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(erroresLexicos)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Creando el parser
		p := parser2.NewGramaticaParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(erroresSintacticos)
		p.BuildParseTrees = true
		tree := p.Ini() //raíz

		// Listener para recorrer el arbol
		var listener *Personal.TreeShapeListener = Personal.NewTreeShapeListener()

		if len(erroresSintacticos.Errors) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}

		ErroresSemanticos.LimpiarLista()
		w.WriteHeader(http.StatusCreated)

		lista := Funcion.RegresarLista()
		Funcion.LimpiarLista()
		json.NewEncoder(w).Encode(map[string]interface{}{"val": lista})
	}
}

func ErroresEntrada() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var solicitu Entrada

		if err := json.NewDecoder(r.Body).Decode(&solicitu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Status": http.StatusBadRequest, "Message": "Error"})
			return
		}

		erroresLexicos := &Personal.CustomErrorListener{}
		erroresSintacticos := &Personal.CustomErrorListener{}

		is := antlr.NewInputStream(solicitu.Texto)

		// Creación de lexer
		lexer := parser2.NewGramaticaLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(erroresLexicos)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Creando el parser
		p := parser2.NewGramaticaParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(erroresSintacticos)
		p.BuildParseTrees = true
		tree := p.Ini() //raíz

		// Listener para recorrer el arbol
		var listener *Personal.TreeShapeListener = Personal.NewTreeShapeListener()

		if len(erroresSintacticos.Errors) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}

		erroresSemanticos := ErroresSemanticos.RegresarErrores()

		var errores []e
		for i := 0; i < len(erroresLexicos.Errors); i++ {
			er := erroresLexicos.Errors[i]
			errores = append(errores, e{er.Line, er.Column, "Lexico", er.Msg})
		}

		for i := 0; i < len(erroresSintacticos.Errors); i++ {
			er := erroresSintacticos.Errors[i]
			errores = append(errores, e{er.Line, er.Column, "Sintáctico", er.Msg})
		}

		for i := 0; i < len(erroresSemanticos.Errores); i++ {
			er := erroresSemanticos.Errores[i]
			errores = append(errores, e{er.Line, er.Column, "Semántico", er.Msg})
		}

		ErroresSemanticos.LimpiarLista()

		w.WriteHeader(http.StatusCreated)

		Funcion.LimpiarLista()
		json.NewEncoder(w).Encode(map[string]interface{}{"message": errores})
	}
}

type e struct {
	Linea   int
	Columna int
	Tipo    string
	Mensaje string
}
