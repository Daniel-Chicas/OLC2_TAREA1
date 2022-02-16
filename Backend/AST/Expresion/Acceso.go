package Expresion

import "OLC2_TAREA1/AST/Entornos"

type Acceso struct {
	Linea   int
	Columna int
	Id      string
}

func NewAcceso(Linea int, Columna int, id string) Acceso {
	return Acceso{Linea, Columna, id}
}

func (A Acceso) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	existe := entorno.ExisteSimbolo(A.Id)
	if existe {
		simb := entorno.ObtenerSimbolo(A.Id)
		return Entornos.RetornoType{Tipo: simb.Tipo, Valor: simb.Valor}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
