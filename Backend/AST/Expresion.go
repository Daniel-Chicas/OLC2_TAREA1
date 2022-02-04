package ast

type Expresion struct {
	Der  *Expresion
	Op   string
	Izq  *Expresion
	Id   *Identificadores
	Prim *Primitivos
}
