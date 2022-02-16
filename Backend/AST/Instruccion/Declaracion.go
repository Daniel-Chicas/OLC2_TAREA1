package Instruccion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Expresion"
	"OLC2_TAREA1/AST/Interfaces"
	"OLC2_TAREA1/ErroresSemanticos"
	arrayList "github.com/colegno/arraylist"
)

var tipoDef = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.STRING, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.BOOLEAN, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

type Declaracion struct {
	Linea     int
	Columna   int
	Inicio    Interfaces.Expresion
	Tipo      Entornos.TipoDato
	Variables *arrayList.List
}

func NuevaDeclaracion(Linea int, Columna int, lista *arrayList.List, tipo Entornos.TipoDato) *Declaracion {
	return &Declaracion{
		Linea:     Linea,
		Columna:   Columna,
		Tipo:      tipo,
		Variables: lista,
	}
}

func NuevaDeclaracionInicio(Linea int, Columna int, lista *arrayList.List, tipo Entornos.TipoDato, inicial Interfaces.Expresion) *Declaracion {
	return &Declaracion{
		Linea:     Linea,
		Columna:   Columna,
		Inicio:    inicial,
		Tipo:      tipo,
		Variables: lista,
	}
}

func (D *Declaracion) Ejecutar(E Entornos.Entorno) interface{} {

	if D.Tipo == Entornos.NULL {
		if D.Variables.Len() > 1 {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "Variable inicializada de manera incorrecta"})
			return nil
		}

		retorno := D.Inicio.ObtenerValor(E)

		tipoExp := retorno.Tipo

		for i := 0; i < D.Variables.Len(); i++ {
			vari := D.Variables.GetValue(i).(Expresion.Identificador)

			if E.ExisteSimbolo(vari.Id) {
				simboloT := Entornos.NuevoSimboloIdentificadorVal(vari.Id, retorno.Valor, tipoExp)
				agregado := E.CambiarValor(vari.Id, simboloT)
				if !agregado {
					ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "La variable debe tener un valor igual al que fue declarada"})
				}
			} else {
				ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "La variable ya fue declarada"})
			}

		}
	} else if D.Inicializado() {
		if D.Variables.Len() > 1 {
			return nil
		}
		retorno := D.Inicio.ObtenerValor(E)

		tipoExp := retorno.Tipo
		tipoDecl := D.Tipo
		tipoResult := tipoDef[tipoExp][tipoDecl]
		if tipoResult == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "Variable inicializada de manera incorrecta"})
			return nil
		}

		for i := 0; i < D.Variables.Len(); i++ {
			vari := D.Variables.GetValue(i).(Expresion.Identificador)
			if E.ExisteSimbolo(vari.Id) {
				ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "La variable ya fue declarada"})
			} else {
				simboloT := Entornos.NuevoSimboloIdentificadorVal(vari.Id, retorno.Valor, tipoResult)
				E.AgregarSimbolo(vari.Id, simboloT)
			}

		}
	} else if !D.Inicializado() {
		if D.Variables.Len() > 1 {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "Variable inicializada de manera incorrecta"})
			return nil
		}
		for i := 0; i < D.Variables.Len(); i++ {
			vari := D.Variables.GetValue(i).(Expresion.Identificador)

			if E.ExisteSimbolo(vari.Id) {
				ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: D.Linea, Column: D.Columna, Msg: "La variable ya fue declarada"})
			} else {
				simboloT := Entornos.NuevoSimboloIdentificador(vari.Id, D.Tipo)
				E.AgregarSimbolo(vari.Id, simboloT)

			}

		}
	}
	return nil
}

func (D *Declaracion) Inicializado() bool {
	return D.Inicio != nil
}
