// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica // Gramatica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterIni is called when entering the ini production.
	EnterIni(c *IniContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

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

	// EnterLogicas is called when entering the logicas production.
	EnterLogicas(c *LogicasContext)

	// EnterRelacionales is called when entering the relacionales production.
	EnterRelacionales(c *RelacionalesContext)

	// EnterAritmeticas is called when entering the aritmeticas production.
	EnterAritmeticas(c *AritmeticasContext)

	// EnterPrimitivos is called when entering the primitivos production.
	EnterPrimitivos(c *PrimitivosContext)

	// ExitIni is called when exiting the ini production.
	ExitIni(c *IniContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

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

	// ExitLogicas is called when exiting the logicas production.
	ExitLogicas(c *LogicasContext)

	// ExitRelacionales is called when exiting the relacionales production.
	ExitRelacionales(c *RelacionalesContext)

	// ExitAritmeticas is called when exiting the aritmeticas production.
	ExitAritmeticas(c *AritmeticasContext)

	// ExitPrimitivos is called when exiting the primitivos production.
	ExitPrimitivos(c *PrimitivosContext)
}
