package Entornos

type Simbolo struct {
	Identificador string
	Valor         interface{}
	Tipo          TipoDato
	Constante     bool
	Funcion       bool
}

func NuevoSimboloIdentificador(identificador string, tipo TipoDato) *Simbolo {
	return &Simbolo{
		Identificador: identificador,
		Tipo:          tipo,
		Constante:     false,
		Funcion:       false}
}

func NuevoSimboloIdentificadorVal(identificador string, valor interface{}, dato TipoDato) *Simbolo {
	return &Simbolo{identificador, valor, dato, false, false}
}
