package Funcion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type SentenciaIf struct {
	Condicion Interfaces.Expresion
	Cuerpo    *arrayList.List
	SenElse   *arrayList.List
}

func NewIf(condicion Interfaces.Expresion, cuerpo *arrayList.List, senelse *arrayList.List) SentenciaIf {
	return SentenciaIf{condicion, cuerpo, senelse}
}

func (Sen SentenciaIf) Ejecutar(entorno Entornos.Entorno) interface{} {
	var nuevoEntorno = Entornos.NuevoEntorno("if", &entorno)
	var condicion = Sen.Condicion.ObtenerValor(nuevoEntorno)
	if condicion.Tipo != Entornos.BOOLEAN {
		fmt.Println("\n\tError semántico, no es posible evaluar la sentencia \"if\" sin un valor boolean.")
	} else {
		if condicion.Valor.(bool) {
			result := Sen.Cuerpo
			for i := 0; i < result.Len(); i++ {
				r := result.GetValue(i)
				if r != nil {
					return result.GetValue(i).(Interfaces.Instruccion).Ejecutar(nuevoEntorno)
				}
			}
		} else if Sen.SenElse != nil {
			var nuevo = Entornos.NuevoEntorno("else", &entorno)
			result := Sen.SenElse
			return result.GetValue(0).(Interfaces.Instruccion).Ejecutar(nuevo)
		}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
