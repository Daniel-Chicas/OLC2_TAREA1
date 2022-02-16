package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	"strings"
)

type Primitivo struct {
	Linea   int
	Columna int
	Valor   interface{}
	Tipo    Entornos.TipoDato
}

func (Pri Primitivo) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	return Entornos.RetornoType{Pri.Tipo, Pri.Valor}
}

func NewPrimitivo(Linea int, columna int, val interface{}, tipo Entornos.TipoDato) Primitivo {
	if tipo == Entornos.BOOLEAN {
		if strings.ToLower(val.(string)) == "true" {
			return Primitivo{Linea, columna, true, tipo}
		} else if strings.ToLower(val.(string)) == "false" {
			return Primitivo{Linea, columna, false, tipo}
		}
	}
	return Primitivo{Linea, columna, val, tipo}
}
