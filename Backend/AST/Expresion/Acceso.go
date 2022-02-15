package Expresion

import "OLC2_TAREA1/AST/Entornos"

type Acceso struct {
	Id    string
	Linea int
}

func NewAcceso(id string, line int) Acceso {
	return Acceso{id, line}
}

func (A Acceso) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	existe := entorno.ExisteSimbolo(A.Id)
	if existe {
		simb := entorno.ObtenerSimbolo(A.Id)
		return Entornos.RetornoType{Tipo: simb.Tipo, Valor: simb.Valor}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
