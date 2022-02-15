package Expresion

import "OLC2_TAREA1/AST/Entornos"

type Identificador struct {
	Id string
}

func NuevoIdentificador(identificador string) Identificador {
	return Identificador{identificador}
}

func (ide Identificador) ObtenerValor(E Entornos.Entorno) Entornos.RetornoType {
	encontrado := E.ExisteSimbolo(ide.Id)
	if !encontrado {
		return Entornos.RetornoType{Valor: nil, Tipo: Entornos.NULL}
	}
	simb := E.ObtenerSimbolo(ide.Id)
	return Entornos.RetornoType{Valor: simb.Valor, Tipo: simb.Tipo}
}
