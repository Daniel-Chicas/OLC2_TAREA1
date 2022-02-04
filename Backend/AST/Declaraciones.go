package ast

type Declaraciones struct {
	If       *SentenciaIf
	Consola  *Imprimir
	Variable []*Variables
}
