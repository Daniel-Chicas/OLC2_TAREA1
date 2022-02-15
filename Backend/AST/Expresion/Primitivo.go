package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	"strings"
)

type Primitivo struct {
	Valor interface{}
	Tipo  Entornos.TipoDato
}

func (Pri Primitivo) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	return Entornos.RetornoType{Pri.Tipo, Pri.Valor}
}

func NewPrimitivo(val interface{}, tipo Entornos.TipoDato) Primitivo {
	if tipo == Entornos.BOOLEAN {
		if strings.ToLower(val.(string)) == "true" {
			return Primitivo{true, tipo}
		} else if strings.ToLower(val.(string)) == "false" {
			return Primitivo{false, tipo}
		}
	}
	return Primitivo{val, tipo}
}
