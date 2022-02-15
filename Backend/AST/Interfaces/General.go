package Interfaces

import (
	"OLC2_TAREA1/AST/Entornos"
)

type Expresion interface {
	ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType
}
type Instruccion interface {
	Ejecutar(entorno Entornos.Entorno) interface{}
}
