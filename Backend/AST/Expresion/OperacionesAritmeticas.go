package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	interfaces2 "OLC2_TAREA1/AST/Interfaces"
	"fmt"
)

var SumaDominante = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.REAL, Entornos.STRING, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.STRING, Entornos.NULL, Entornos.NULL},
	{Entornos.STRING, Entornos.STRING, Entornos.STRING, Entornos.STRING, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.STRING, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

var RestaDominante = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

var MultiDominante = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

var DiviDominante = [5][5]Entornos.TipoDato{
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

var ModDominante = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

type Operacion struct {
	Izquierda interfaces2.Expresion
	Operador  string
	Derecha   interfaces2.Expresion
	Unario    bool
}

func NewAritmetica(Izq interfaces2.Expresion, Operador string, Der interfaces2.Expresion, unario bool) Operacion {
	return Operacion{Izq, Operador, Der, unario}
}

func (Op Operacion) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	var izq Entornos.RetornoType
	var der Entornos.RetornoType

	if Op.Unario == true {
		der = Op.Derecha.ObtenerValor(entorno)
		if der.Tipo == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: der.Tipo, Valor: -der.Valor.(int)}
		} else if der.Tipo == Entornos.REAL {
			return Entornos.RetornoType{Tipo: der.Tipo, Valor: -der.Valor.(float64)}
		}
	} else {
		izq = Op.Izquierda.ObtenerValor(entorno)
		der = Op.Derecha.ObtenerValor(entorno)
	}

	var dominante Entornos.TipoDato

	switch Op.Operador {
	case "+":
		dominante = SumaDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(int) + der.Valor.(int)}
		} else if dominante == Entornos.REAL {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(float64) + der.Valor.(float64)}
		} else if dominante == Entornos.STRING {
			r1 := fmt.Sprintf("%v", izq.Valor)
			r2 := fmt.Sprintf("%v", der.Valor)
			return Entornos.RetornoType{Tipo: dominante, Valor: r1 + r2}
		} else if dominante == Entornos.NULL {
			fmt.Println("NO SE PUEDEN SUMAR ESTOS VALORES")
		}
		break
	case "-":
		dominante = RestaDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(int) - der.Valor.(int)}
		} else if dominante == Entornos.REAL {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(float64) - der.Valor.(float64)}
		} else if dominante == Entornos.NULL {
			fmt.Println("NO SE PUEDEN RESTAR ESTOS VALORES")
		}
		break
	case "*":
		dominante = MultiDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(int) * der.Valor.(int)}
		} else if dominante == Entornos.REAL {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(float64) * der.Valor.(float64)}
		} else if dominante == Entornos.NULL {
			fmt.Println("NO SE PUEDEN MULTIPLICAR ESTOS VALORES")
		}
		break
	case "/":
		dominante = DiviDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.REAL {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(float64) / der.Valor.(float64)}
		} else if dominante == Entornos.NULL {
			fmt.Println("NO SE PUEDEN DIVIDIR ESTOS VALORES")
		}
		break
	case "%":
		dominante = ModDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: dominante, Valor: izq.Valor.(int) % der.Valor.(int)}
		} else if dominante == Entornos.NULL {
			fmt.Println("NO SE PUEDE OBTENER EL MÓDULO DE ESTOS VALORES")
		}
		break
	}

	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
