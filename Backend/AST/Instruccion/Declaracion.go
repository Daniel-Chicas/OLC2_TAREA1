package Instruccion

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Expresion"
	"OLC2_TAREA1/AST/Interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

var tipoDef = [5][5]Entornos.TipoDato{
	{Entornos.INTEGER, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.REAL, Entornos.REAL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.STRING, Entornos.NULL, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.BOOLEAN, Entornos.NULL},
	{Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL, Entornos.NULL},
}

type Declaracion struct {
	Inicio    Interfaces.Expresion
	Tipo      Entornos.TipoDato
	Variables *arrayList.List
}

func NuevaDeclaracion(lista *arrayList.List, tipo Entornos.TipoDato) *Declaracion {
	return &Declaracion{
		Tipo:      tipo,
		Variables: lista,
	}
}

func NuevaDeclaracionInicio(lista *arrayList.List, tipo Entornos.TipoDato, inicial Interfaces.Expresion) *Declaracion {
	return &Declaracion{
		Inicio:    inicial,
		Tipo:      tipo,
		Variables: lista,
	}
}

func (D *Declaracion) Ejecutar(E Entornos.Entorno) interface{} {
	if D.Inicializado() {
		if D.Variables.Len() > 1 {
			return nil
		}
		retorno := D.Inicio.ObtenerValor(E)

		tipoExp := retorno.Tipo
		tipoDecl := D.Tipo
		tipoResult := tipoDef[tipoExp][tipoDecl]
		if tipoResult == Entornos.NULL {
			return nil
		}

		for i := 0; i < D.Variables.Len(); i++ {
			vari := D.Variables.GetValue(i).(Expresion.Identificador)
			if E.ExisteSimbolo(vari.Id) {
				fmt.Printf("\nError, la variable %s ya está declarada \n", vari.Id)
			} else {
				simboloT := Entornos.NuevoSimboloIdentificadorVal(0, 0, vari.Id, retorno.Valor, tipoResult)
				E.AgregarSimbolo(vari.Id, simboloT)
			}

		}
	} else if !D.Inicializado() {
		if D.Variables.Len() > 1 {
			return nil
		}

		for i := 0; i < D.Variables.Len(); i++ {
			vari := D.Variables.GetValue(i).(Expresion.Identificador)

			if E.ExisteSimbolo(vari.Id) {
				fmt.Printf("\nError, la variable %s ya está declarada \n", vari.Id)
			} else {
				if D.Tipo == Entornos.INTEGER {
					simboloT := Entornos.NuevoSimboloIdentificadorVal(0, 0, vari.Id, 0, D.Tipo)
					E.AgregarSimbolo(vari.Id, simboloT)
				} else if D.Tipo == Entornos.REAL {
					var x float64 = 0.0
					simboloT := Entornos.NuevoSimboloIdentificadorVal(0, 0, vari.Id, x, D.Tipo)
					E.AgregarSimbolo(vari.Id, simboloT)
				} else if D.Tipo == Entornos.STRING {
					simboloT := Entornos.NuevoSimboloIdentificadorVal(0, 0, vari.Id, "", D.Tipo)
					E.AgregarSimbolo(vari.Id, simboloT)
				} else if D.Tipo == Entornos.BOOLEAN {
					simboloT := Entornos.NuevoSimboloIdentificadorVal(0, 0, vari.Id, true, D.Tipo)
					E.AgregarSimbolo(vari.Id, simboloT)
				} else {
					fmt.Printf("\nERROR SEMÁNTICO, TIPO DE DATO NO VÁLIDO: %d\n", D.Tipo)
				}
			}

		}
	}
	//data, err := json.MarshalIndent(E, "", "  ")
	//if err != nil {
	//	panic(err)
	//}

	//stringEsQuery := string(data)
	//fmt.Println(stringEsQuery)
	//fmt.Printf("%v", D.Variables)
	return nil
}

func (D *Declaracion) Inicializado() bool {
	return D.Inicio != nil
}
