package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
	"OLC2_TAREA1/ErroresSemanticos"
)

type Logico struct {
	Linea      int
	Columna    int
	izq        Interfaces.Expresion
	TipoLogico string
	der        Interfaces.Expresion
	Unario     bool
}

func NewLogico(Linea int, Columna int, Izq Interfaces.Expresion, Operador string, Der Interfaces.Expresion, unario bool) Logico {
	return Logico{Linea, Columna, Izq, Operador, Der, unario}
}

func (Log Logico) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	var izq = Log.izq.ObtenerValor(entorno)

	switch Log.TipoLogico {
	case "&&":
		var der = Log.der.ObtenerValor(entorno)
		if izq.Tipo != Entornos.BOOLEAN && der.Tipo != Entornos.BOOLEAN {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Log.Linea, Column: Log.Columna, Msg: "No se puede tener un resultado lógico con estos valores"})
			return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
		}
		var result = izq.Valor.(bool) && der.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	case "||":
		var der = Log.der.ObtenerValor(entorno)
		if izq.Tipo != Entornos.BOOLEAN && der.Tipo != Entornos.BOOLEAN {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Log.Linea, Column: Log.Columna, Msg: "No se puede tener un resultado lógico con estos valores"})
			return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
		}
		var result = izq.Valor.(bool) || der.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	case "!":
		if izq.Tipo != Entornos.BOOLEAN {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Log.Linea, Column: Log.Columna, Msg: "No se puede tener un resultado lógico con estos valores"})
			return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
		}
		var result = !izq.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
