package Entornos

type Simbolo struct {
	Linea         int
	Columna       int
	Identificador string
	Valor         interface{}
	Tipo          TipoDato
	Constante     bool
	Funcion       bool
}

func NuevoSimboloIdentificador(linea int, columna int, identificador string, tipo TipoDato) *Simbolo {
	return &Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Tipo:          tipo,
		Constante:     false,
		Funcion:       false}
}

func NuevoSimboloIdentificadorVal(linea int, columna int, identificador string, valor interface{}, dato TipoDato) *Simbolo {
	return &Simbolo{linea, columna, identificador, valor, dato, false, false}
}
