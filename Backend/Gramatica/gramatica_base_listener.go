// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica // Gramatica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIni is called when production ini is entered.
func (s *BaseGramaticaListener) EnterIni(ctx *IniContext) {}

// ExitIni is called when production ini is exited.
func (s *BaseGramaticaListener) ExitIni(ctx *IniContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseGramaticaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseGramaticaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterDeclaraciones is called when production declaraciones is entered.
func (s *BaseGramaticaListener) EnterDeclaraciones(ctx *DeclaracionesContext) {}

// ExitDeclaraciones is called when production declaraciones is exited.
func (s *BaseGramaticaListener) ExitDeclaraciones(ctx *DeclaracionesContext) {}

// EnterSentenceIf is called when production sentenceIf is entered.
func (s *BaseGramaticaListener) EnterSentenceIf(ctx *SentenceIfContext) {}

// ExitSentenceIf is called when production sentenceIf is exited.
func (s *BaseGramaticaListener) ExitSentenceIf(ctx *SentenceIfContext) {}

// EnterSentenceElse is called when production sentenceElse is entered.
func (s *BaseGramaticaListener) EnterSentenceElse(ctx *SentenceElseContext) {}

// ExitSentenceElse is called when production sentenceElse is exited.
func (s *BaseGramaticaListener) ExitSentenceElse(ctx *SentenceElseContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGramaticaListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGramaticaListener) ExitVariable(ctx *VariableContext) {}

// EnterIdentificadores is called when production identificadores is entered.
func (s *BaseGramaticaListener) EnterIdentificadores(ctx *IdentificadoresContext) {}

// ExitIdentificadores is called when production identificadores is exited.
func (s *BaseGramaticaListener) ExitIdentificadores(ctx *IdentificadoresContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseGramaticaListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseGramaticaListener) ExitTipo(ctx *TipoContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseGramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseGramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterLogicas is called when production logicas is entered.
func (s *BaseGramaticaListener) EnterLogicas(ctx *LogicasContext) {}

// ExitLogicas is called when production logicas is exited.
func (s *BaseGramaticaListener) ExitLogicas(ctx *LogicasContext) {}

// EnterRelacionales is called when production relacionales is entered.
func (s *BaseGramaticaListener) EnterRelacionales(ctx *RelacionalesContext) {}

// ExitRelacionales is called when production relacionales is exited.
func (s *BaseGramaticaListener) ExitRelacionales(ctx *RelacionalesContext) {}

// EnterAritmeticas is called when production aritmeticas is entered.
func (s *BaseGramaticaListener) EnterAritmeticas(ctx *AritmeticasContext) {}

// ExitAritmeticas is called when production aritmeticas is exited.
func (s *BaseGramaticaListener) ExitAritmeticas(ctx *AritmeticasContext) {}

// EnterPrimitivos is called when production primitivos is entered.
func (s *BaseGramaticaListener) EnterPrimitivos(ctx *PrimitivosContext) {}

// ExitPrimitivos is called when production primitivos is exited.
func (s *BaseGramaticaListener) ExitPrimitivos(ctx *PrimitivosContext) {}
