package Personal

import (
	"OLC2_TAREA1/AST/Entornos"
	"OLC2_TAREA1/AST/Interfaces"
	parser2 "OLC2_TAREA1/Gramatica"
	"fmt"
)

type TreeShapeListener struct {
	*parser2.BaseGramaticaListener
	Cadena []string
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitIni(ctx *parser2.IniContext) {
	result := ctx.GetLista()
	entornoGlobal := Entornos.NuevoEntorno("Global", nil)
	for i := 0; i < result.Len(); i++ {
		r := result.GetValue(i)
		if r != nil {
			val := result.GetValue(i).(Interfaces.Instruccion).Ejecutar(entornoGlobal)
			if val != nil {
				this.Cadena = append(this.Cadena, fmt.Sprintf("%v", val.(Entornos.RetornoType).Valor))
			}
		}

	}
}
