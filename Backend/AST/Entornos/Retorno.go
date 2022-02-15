package Entornos

type TipoDato int

const (
	INTEGER TipoDato = iota
	REAL
	STRING
	BOOLEAN
	NULL
)

type RetornoType struct {
	Tipo  TipoDato
	Valor interface{}
}
