package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
)

type Logico struct {
	izq        Interfaces.Expresion
	TipoLogico string
	der        Interfaces.Expresion
	Unario     bool
}

func NewLogico(Izq Interfaces.Expresion, Operador string, Der Interfaces.Expresion, unario bool) Logico {
	return Logico{Izq, Operador, Der, unario}
}

func (Log Logico) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	var izq = Log.izq.ObtenerValor(entorno)

	switch Log.TipoLogico {
	case "&&":
		var der = Log.der.ObtenerValor(entorno)
		var result = izq.Valor.(bool) && der.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	case "||":
		var der = Log.der.ObtenerValor(entorno)
		var result = izq.Valor.(bool) || der.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	case "!":
		var result = !izq.Valor.(bool)
		return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
