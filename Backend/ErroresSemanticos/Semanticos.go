package ErroresSemanticos

type ErrorSemantico struct {
	Line   int
	Column int
	Msg    string
}

type ErroresSemanticos struct {
	Errores []ErrorSemantico
}

var Errores ErroresSemanticos

func RegresarErrores() ErroresSemanticos {
	return Errores
}

func LimpiarLista() {
	var nueva ErroresSemanticos
	Errores = nueva
}

func AgregarErrores(nuevo ErrorSemantico) {
	Errores.Errores = append(Errores.Errores, nuevo)
}
