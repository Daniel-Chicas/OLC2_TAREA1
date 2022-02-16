package Funcion

import (
	"OLC2_TAREA1/AST/Entornos"
	interfaces2 "OLC2_TAREA1/AST/Interfaces"
	"fmt"
)

type Imprimir struct {
	Expresiones interfaces2.Expresion
}

func NewImprimir(val interfaces2.Expresion) Imprimir {
	return Imprimir{val}
}

var lista []string

func (Imp Imprimir) Ejecutar(entorno Entornos.Entorno) interface{} {
	impr := Imp.Expresiones.ObtenerValor(entorno)
	if impr.Valor != nil {
		lista = append(lista, fmt.Sprintf("%v", impr.Valor))
	}
	return impr
}

func RegresarLista() []string {
	return lista
}

func LimpiarLista() {
	var nueva []string
	lista = nueva
}
