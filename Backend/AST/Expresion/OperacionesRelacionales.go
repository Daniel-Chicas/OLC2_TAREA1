package Expresion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
	"fmt"
)

type Relacional struct {
	izq            Interfaces.Expresion
	TipoRelacional string
	der            Interfaces.Expresion
	Unario         bool
}

func NewRelacional(Izq Interfaces.Expresion, Operador string, Der Interfaces.Expresion, unario bool) Relacional {
	return Relacional{Izq, Operador, Der, unario}
}

func (Rel Relacional) ObtenerValor(entorno Entornos.Entorno) Entornos.RetornoType {
	var izq = Rel.izq.ObtenerValor(entorno)
	var der = Rel.der.ObtenerValor(entorno)

	switch Rel.TipoRelacional {
	case "<=":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) <= der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) <= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) <= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) <= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) <= der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	case ">=":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) >= der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) >= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) >= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) >= der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) >= der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	case "<":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) < der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) < der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) < der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) < der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) < der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	case ">":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) > der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) > der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) > der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) > der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) > der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	case "==":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) == der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) == der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) == der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) == der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) == der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	case "!=":
		if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(int) != der.Valor.(int)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.INTEGER && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) != der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.INTEGER {
			var result = izq.Valor.(float32) != der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.REAL && der.Tipo == Entornos.REAL {
			var result = izq.Valor.(float32) != der.Valor.(float32)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else if izq.Tipo == Entornos.STRING && der.Tipo == Entornos.STRING {
			var result = izq.Valor.(string) != der.Valor.(string)
			return Entornos.RetornoType{Tipo: Entornos.BOOLEAN, Valor: result}
		} else {
			fmt.Println("NO SE PUEDEN RELACIONAR ESTOS VALORES")
		}
		break
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
