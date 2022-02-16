package Funcion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
	"OLC2_TAREA1/ErroresSemanticos"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type SentenciaIf struct {
	Linea     int
	Columna   int
	Condicion Interfaces.Expresion
	Cuerpo    *arrayList.List
	SenElse   *arrayList.List
}

func NewIf(Linea int, Columna int, condicion Interfaces.Expresion, cuerpo *arrayList.List, senelse *arrayList.List) SentenciaIf {
	return SentenciaIf{Linea, Columna, condicion, cuerpo, senelse}
}

func (Sen SentenciaIf) Ejecutar(entorno Entornos.Entorno) interface{} {
	var nuevoEntorno = Entornos.NuevoEntorno("if", &entorno)
	var condicion = Sen.Condicion.ObtenerValor(nuevoEntorno)
	if condicion.Tipo != Entornos.BOOLEAN {
		ErroresSemanticos.AgregarErrores(ErroresSemanticos.ErrorSemantico{Line: Sen.Linea, Column: Sen.Columna, Msg: "No es posible evaluar la condición de la sentencia"})
	} else {
		if condicion.Valor.(bool) {
			result := Sen.Cuerpo
			for i := 0; i < result.Len(); i++ {
				r := result.GetValue(i)
				if r != nil {
					result.GetValue(i).(Interfaces.Instruccion).Ejecutar(nuevoEntorno)
				}
			}
			return nil
		} else if Sen.SenElse != nil {
			var nuevo = Entornos.NuevoEntorno("else", &entorno)
			result := Sen.SenElse
			esIf := false
			for i := 0; i < result.Len(); i++ {
				r := result.GetValue(i)
				if r != nil && reflect.TypeOf(r).String() != "Funcion.SentenciaIf" {
					result.GetValue(i).(Interfaces.Instruccion).Ejecutar(nuevoEntorno)
					esIf = true
				}
			}
			if !esIf {
				return result.GetValue(0).(Interfaces.Instruccion).Ejecutar(nuevo)
			}
		}
	}
	return Entornos.RetornoType{Tipo: Entornos.NULL, Valor: nil}
}
