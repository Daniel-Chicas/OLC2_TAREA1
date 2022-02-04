// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica // Gramatica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterIni is called when entering the ini production.
	EnterIni(c *IniContext)

	// EnterFuncionMetodo is called when entering the funcionMetodo production.
	EnterFuncionMetodo(c *FuncionMetodoContext)

	// EnterCuerpoFunciones is called when entering the cuerpoFunciones production.
	EnterCuerpoFunciones(c *CuerpoFuncionesContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

	// EnterListaExpresiones is called when entering the listaExpresiones production.
	EnterListaExpresiones(c *ListaExpresionesContext)

	// EnterSentenceIf is called when entering the sentenceIf production.
	EnterSentenceIf(c *SentenceIfContext)

	// EnterSentenceElse is called when entering the sentenceElse production.
	EnterSentenceElse(c *SentenceElseContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterIdentificadores is called when entering the identificadores production.
	EnterIdentificadores(c *IdentificadoresContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterPrimitivos is called when entering the primitivos production.
	EnterPrimitivos(c *PrimitivosContext)

	// ExitIni is called when exiting the ini production.
	ExitIni(c *IniContext)

	// ExitFuncionMetodo is called when exiting the funcionMetodo production.
	ExitFuncionMetodo(c *FuncionMetodoContext)

	// ExitCuerpoFunciones is called when exiting the cuerpoFunciones production.
	ExitCuerpoFunciones(c *CuerpoFuncionesContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitListaExpresiones is called when exiting the listaExpresiones production.
	ExitListaExpresiones(c *ListaExpresionesContext)

	// ExitSentenceIf is called when exiting the sentenceIf production.
	ExitSentenceIf(c *SentenceIfContext)

	// ExitSentenceElse is called when exiting the sentenceElse production.
	ExitSentenceElse(c *SentenceElseContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitIdentificadores is called when exiting the identificadores production.
	ExitIdentificadores(c *IdentificadoresContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitPrimitivos is called when exiting the primitivos production.
	ExitPrimitivos(c *PrimitivosContext)
}
