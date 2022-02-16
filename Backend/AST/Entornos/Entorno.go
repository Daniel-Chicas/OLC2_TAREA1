package Entornos

import "strings"

type Entorno struct {
	Nombre   string
	Anterior *Entorno
	Tabla    map[string]interface{}
}

func NuevoEntorno(nombre string, Anterior *Entorno) Entorno {
	nuevo := Entorno{Nombre: nombre, Anterior: Anterior}
	nuevo.Tabla = make(map[string]interface{})
	return nuevo
}

func (E *Entorno) ExisteSimbolo(identificador string) bool {
	id := strings.ToLower(identificador)
	for actual := E; actual != nil; actual = actual.Anterior {
		for key, _ := range actual.Tabla {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (E *Entorno) AgregarSimbolo(identificador string, simbolo *Simbolo) {
	id := strings.ToLower(identificador)
	E.Tabla[id] = simbolo
}

func (E *Entorno) ObtenerSimbolo(identificador string) *Simbolo {
	id := strings.ToLower(identificador)
	for actual := E; actual != nil; actual = actual.Anterior {
		for key, simboloElement := range actual.Tabla {
			if key == id {
				return simboloElement.(*Simbolo)
			}
		}
	}
	var simboloNil *Simbolo
	return simboloNil
}

func (E *Entorno) CambiarValor(identificador string, nuevo *Simbolo) bool {
	id := strings.ToLower(identificador)
	for actual := E; actual != nil; actual = actual.Anterior {
		for key, simboloElement := range actual.Tabla {
			if key == id {
				if simboloElement.(*Simbolo).Tipo == nuevo.Tipo {
					actual.Tabla[id] = nuevo
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}

//Faltarían las de funciones
