package Expresion

import "OLC2_TAREA1/AST/Entornos"

type Identificador struct {
	Linea   int
	Columna int
	Id      string
}

func NuevoIdentificador(Linea int, Columna int, identificador string) Identificador {
	return Identificador{Linea, Columna, identificador}
}

func (ide Identificador) ObtenerValor(E Entornos.Entorno) Entornos.RetornoType {
	encontrado := E.ExisteSimbolo(ide.Id)
	if !encontrado {
		return Entornos.RetornoType{Valor: nil, Tipo: Entornos.NULL}
	}
	simb := E.ObtenerSimbolo(ide.Id)
	return Entornos.RetornoType{Valor: simb.Valor, Tipo: simb.Tipo}
}
