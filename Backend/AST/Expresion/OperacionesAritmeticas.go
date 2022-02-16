package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	interfaces2 "OLC2_TAREA1/AST/Interfaces"
	"OLC2_TAREA1/ErroresSemanticos"
	"fmt"
	"strconv"
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
	{Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

type Operacion struct {
	Linea     int
	Columna   int
	Izquierda interfaces2.Expresion
	Operador  string
	Derecha   interfaces2.Expresion
	Unario    bool
}

func NewAritmetica(linea int, columna int, Izq interfaces2.Expresion, Operador string, Der interfaces2.Expresion, unario bool) Operacion {
	return Operacion{linea, columna, Izq, Operador, Der, unario}
}

func (Op Operacion) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	var izq Entornos.RetornoType
	var der Entornos.RetornoType

	if Op.Unario == true {
		der = Op.Derecha.ObtenerValor(entorno)
		if der.Tipo == Entornos.INTEGER {
			return Entornos.RetornoType{Tipo: der.Tipo, Valor: -der.Valor.(int)}
		} else if der.Tipo == Entornos.REAL {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			return Entornos.RetornoType{Tipo: der.Tipo, Valor: -val1}
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
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			resultado, _ := strconv.ParseInt(fmt.Sprintf("%v", val1+val2), 0, 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: resultado}
		} else if dominante == Entornos.REAL {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: val1 + val2}
		} else if dominante == Entornos.STRING {
			r1 := fmt.Sprintf("%v", izq.Valor)
			r2 := fmt.Sprintf("%v", der.Valor)
			return Entornos.RetornoType{Tipo: dominante, Valor: r1 + r2}
		} else if dominante == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se pueden sumar los valores"})
		}
		break
	case "-":
		dominante = RestaDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			resultado, _ := strconv.ParseInt(fmt.Sprintf("%v", val1-val2), 0, 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: resultado}
		} else if dominante == Entornos.REAL {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: val1 - val2}
		} else if dominante == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se pueden restar los valores"})
		}
		break
	case "*":
		dominante = MultiDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.INTEGER {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			resultado, _ := strconv.ParseInt(fmt.Sprintf("%v", val1*val2), 0, 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: resultado}
		} else if dominante == Entornos.REAL {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: val1 * val2}
		} else if dominante == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se pueden multiplicar los valores"})
		}
		break
	case "/":
		dominante = DiviDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.REAL {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", izq.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", der.Valor), 64)
			if val2 == 0 {
				ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se permite dividir entre 0"})
			}
			return Entornos.RetornoType{Tipo: dominante, Valor: val1 / val2}
		} else if dominante == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se pueden dividir los valores"})
		}
		break
	case "%":
		dominante = ModDominante[izq.Tipo][der.Tipo]
		if dominante == Entornos.REAL {
			val1, _ := strconv.ParseInt(fmt.Sprintf("%v", izq.Valor), 0, 64)
			val2, _ := strconv.ParseInt(fmt.Sprintf("%v", der.Valor), 0, 64)
			return Entornos.RetornoType{Tipo: dominante, Valor: val1 % val2}
		} else if dominante == Entornos.NULL {
			ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Op.Linea, Column: Op.Columna, Msg: "No se puede obtener el módulo de los valores"})
		}
		break
	}

	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
