package Funcion

import (
	"OLC2_TAREA1/AST/Entornos"
	interfaces2 "OLC2_TAREA1/AST/Interfaces"
)

type Imprimir struct {
	Expresiones interfaces2.Expresion
}

func NewImprimir(val interfaces2.Expresion) Imprimir {
	return Imprimir{val}
}

func (Imp Imprimir) Ejecutar(entorno Entornos.Entorno) interface{} {
	impr := Imp.Expresiones.ObtenerValor(entorno)
	return impr
}
