// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica // Gramatica
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 188,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 53, 10, 4, 12, 4, 14, 4, 56, 11, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 66, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 97, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 116, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 124,
	10, 10, 12, 10, 14, 10, 127, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 140, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 181, 10, 12, 12, 12, 14,
	12, 184, 11, 12, 3, 13, 3, 13, 3, 13, 2, 6, 6, 10, 18, 22, 14, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 4, 2, 10, 14, 50, 50, 3, 2, 43,
	47, 2, 200, 2, 26, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 47, 3, 2, 2, 2, 8,
	65, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 78, 3, 2, 2, 2, 14, 96, 3, 2, 2,
	2, 16, 115, 3, 2, 2, 2, 18, 117, 3, 2, 2, 2, 20, 128, 3, 2, 2, 2, 22, 139,
	3, 2, 2, 2, 24, 185, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2, 27, 28, 7, 2, 2, 3,
	28, 3, 3, 2, 2, 2, 29, 30, 7, 7, 2, 2, 30, 31, 7, 6, 2, 2, 31, 32, 7, 40,
	2, 2, 32, 33, 7, 48, 2, 2, 33, 34, 7, 38, 2, 2, 34, 35, 7, 15, 2, 2, 35,
	36, 7, 34, 2, 2, 36, 37, 7, 11, 2, 2, 37, 38, 7, 36, 2, 2, 38, 39, 7, 37,
	2, 2, 39, 40, 7, 35, 2, 2, 40, 41, 7, 40, 2, 2, 41, 42, 7, 16, 2, 2, 42,
	43, 7, 38, 2, 2, 43, 44, 5, 6, 4, 2, 44, 45, 7, 39, 2, 2, 45, 46, 7, 39,
	2, 2, 46, 5, 3, 2, 2, 2, 47, 48, 8, 4, 1, 2, 48, 49, 5, 8, 5, 2, 49, 54,
	3, 2, 2, 2, 50, 51, 12, 4, 2, 2, 51, 53, 5, 8, 5, 2, 52, 50, 3, 2, 2, 2,
	53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 7, 3, 2,
	2, 2, 56, 54, 3, 2, 2, 2, 57, 66, 5, 12, 7, 2, 58, 59, 7, 5, 2, 2, 59,
	60, 7, 34, 2, 2, 60, 61, 5, 10, 6, 2, 61, 62, 7, 35, 2, 2, 62, 63, 7, 41,
	2, 2, 63, 66, 3, 2, 2, 2, 64, 66, 5, 16, 9, 2, 65, 57, 3, 2, 2, 2, 65,
	58, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 9, 3, 2, 2, 2, 67, 68, 8, 6, 1,
	2, 68, 69, 5, 22, 12, 2, 69, 75, 3, 2, 2, 2, 70, 71, 12, 4, 2, 2, 71, 72,
	7, 42, 2, 2, 72, 74, 5, 22, 12, 2, 73, 70, 3, 2, 2, 2, 74, 77, 3, 2, 2,
	2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77, 75,
	3, 2, 2, 2, 78, 79, 7, 3, 2, 2, 79, 80, 7, 40, 2, 2, 80, 81, 7, 34, 2,
	2, 81, 82, 5, 22, 12, 2, 82, 83, 7, 35, 2, 2, 83, 84, 7, 38, 2, 2, 84,
	85, 5, 6, 4, 2, 85, 86, 7, 39, 2, 2, 86, 87, 5, 14, 8, 2, 87, 13, 3, 2,
	2, 2, 88, 89, 7, 4, 2, 2, 89, 90, 7, 38, 2, 2, 90, 91, 5, 6, 4, 2, 91,
	92, 7, 39, 2, 2, 92, 97, 3, 2, 2, 2, 93, 94, 7, 4, 2, 2, 94, 97, 5, 12,
	7, 2, 95, 97, 3, 2, 2, 2, 96, 88, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 96, 95,
	3, 2, 2, 2, 97, 15, 3, 2, 2, 2, 98, 99, 7, 9, 2, 2, 99, 100, 5, 18, 10,
	2, 100, 101, 5, 20, 11, 2, 101, 102, 7, 41, 2, 2, 102, 116, 3, 2, 2, 2,
	103, 104, 7, 9, 2, 2, 104, 105, 5, 18, 10, 2, 105, 106, 5, 20, 11, 2, 106,
	107, 7, 22, 2, 2, 107, 108, 5, 22, 12, 2, 108, 109, 7, 41, 2, 2, 109, 116,
	3, 2, 2, 2, 110, 111, 5, 18, 10, 2, 111, 112, 7, 22, 2, 2, 112, 113, 5,
	22, 12, 2, 113, 114, 7, 41, 2, 2, 114, 116, 3, 2, 2, 2, 115, 98, 3, 2,
	2, 2, 115, 103, 3, 2, 2, 2, 115, 110, 3, 2, 2, 2, 116, 17, 3, 2, 2, 2,
	117, 118, 8, 10, 1, 2, 118, 119, 7, 48, 2, 2, 119, 125, 3, 2, 2, 2, 120,
	121, 12, 4, 2, 2, 121, 122, 7, 42, 2, 2, 122, 124, 5, 18, 10, 5, 123, 120,
	3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2,
	2, 2, 126, 19, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 129, 9, 2, 2, 2,
	129, 21, 3, 2, 2, 2, 130, 131, 8, 12, 1, 2, 131, 132, 7, 21, 2, 2, 132,
	140, 5, 22, 12, 19, 133, 134, 7, 34, 2, 2, 134, 135, 5, 22, 12, 2, 135,
	136, 7, 35, 2, 2, 136, 140, 3, 2, 2, 2, 137, 140, 7, 48, 2, 2, 138, 140,
	5, 24, 13, 2, 139, 130, 3, 2, 2, 2, 139, 133, 3, 2, 2, 2, 139, 137, 3,
	2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 182, 3, 2, 2, 2, 141, 142, 12, 18,
	2, 2, 142, 143, 7, 23, 2, 2, 143, 181, 5, 22, 12, 19, 144, 145, 12, 17,
	2, 2, 145, 146, 7, 24, 2, 2, 146, 181, 5, 22, 12, 18, 147, 148, 12, 16,
	2, 2, 148, 149, 7, 20, 2, 2, 149, 181, 5, 22, 12, 17, 150, 151, 12, 15,
	2, 2, 151, 152, 7, 21, 2, 2, 152, 181, 5, 22, 12, 16, 153, 154, 12, 14,
	2, 2, 154, 155, 7, 18, 2, 2, 155, 181, 5, 22, 12, 15, 156, 157, 12, 13,
	2, 2, 157, 158, 7, 19, 2, 2, 158, 181, 5, 22, 12, 14, 159, 160, 12, 12,
	2, 2, 160, 161, 7, 17, 2, 2, 161, 181, 5, 22, 12, 13, 162, 163, 12, 10,
	2, 2, 163, 164, 7, 30, 2, 2, 164, 181, 5, 22, 12, 11, 165, 166, 12, 9,
	2, 2, 166, 167, 7, 31, 2, 2, 167, 181, 5, 22, 12, 10, 168, 169, 12, 8,
	2, 2, 169, 170, 7, 29, 2, 2, 170, 181, 5, 22, 12, 9, 171, 172, 12, 7, 2,
	2, 172, 173, 7, 26, 2, 2, 173, 181, 5, 22, 12, 8, 174, 175, 12, 6, 2, 2,
	175, 176, 7, 28, 2, 2, 176, 181, 5, 22, 12, 7, 177, 178, 12, 5, 2, 2, 178,
	179, 7, 27, 2, 2, 179, 181, 5, 22, 12, 6, 180, 141, 3, 2, 2, 2, 180, 144,
	3, 2, 2, 2, 180, 147, 3, 2, 2, 2, 180, 150, 3, 2, 2, 2, 180, 153, 3, 2,
	2, 2, 180, 156, 3, 2, 2, 2, 180, 159, 3, 2, 2, 2, 180, 162, 3, 2, 2, 2,
	180, 165, 3, 2, 2, 2, 180, 168, 3, 2, 2, 2, 180, 171, 3, 2, 2, 2, 180,
	174, 3, 2, 2, 2, 180, 177, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 23, 3, 2, 2, 2, 184, 182, 3, 2,
	2, 2, 185, 186, 9, 3, 2, 2, 186, 25, 3, 2, 2, 2, 11, 54, 65, 75, 96, 115,
	125, 139, 180, 182,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'%'", "'*'",
	"'/'", "'+'", "'-'", "'='", "'&&'", "'||'", "'!'", "'<='", "'<'", "'>'",
	"'>='", "'=='", "'!='", "'_'", "'.'", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "':'", "';'", "','",
}
var symbolicNames = []string{
	"", "IF", "ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", "DOUBLE",
	"STRING", "INT", "REAL", "BOOLEAN", "PRINCIPAL", "METODO", "MOD", "POR",
	"DIVIDIR", "MAS", "MENOS", "IGUAL", "AND", "OR", "NOT", "MENI", "MEN",
	"MAY", "MAYI", "IGUALI", "DIFERENCIA", "GUIONB", "PUNTO", "PARA", "PARC",
	"LLABRE", "LLACIE", "CABRE", "CCIER", "DPUNTOS", "PCOMA", "COMA", "CADENA",
	"ENTERO", "DECIMAL", "TRUE", "FALSE", "ID", "ESPACIOS", "CHAR",
}

var ruleNames = []string{
	"ini", "funcionMetodo", "cuerpoFunciones", "declaraciones", "listaExpresiones",
	"sentenceIf", "sentenceElse", "variable", "identificadores", "tipo", "expresion",
	"primitivos",
}

type GramaticaParser struct {
	*antlr.BaseParser
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserIF         = 1
	GramaticaParserELSE       = 2
	GramaticaParserIMPRIMIR   = 3
	GramaticaParserMAIN       = 4
	GramaticaParserPUBLIC     = 5
	GramaticaParserCLASS      = 6
	GramaticaParserDECLARAR   = 7
	GramaticaParserDOUBLE     = 8
	GramaticaParserSTRING     = 9
	GramaticaParserINT        = 10
	GramaticaParserREAL       = 11
	GramaticaParserBOOLEAN    = 12
	GramaticaParserPRINCIPAL  = 13
	GramaticaParserMETODO     = 14
	GramaticaParserMOD        = 15
	GramaticaParserPOR        = 16
	GramaticaParserDIVIDIR    = 17
	GramaticaParserMAS        = 18
	GramaticaParserMENOS      = 19
	GramaticaParserIGUAL      = 20
	GramaticaParserAND        = 21
	GramaticaParserOR         = 22
	GramaticaParserNOT        = 23
	GramaticaParserMENI       = 24
	GramaticaParserMEN        = 25
	GramaticaParserMAY        = 26
	GramaticaParserMAYI       = 27
	GramaticaParserIGUALI     = 28
	GramaticaParserDIFERENCIA = 29
	GramaticaParserGUIONB     = 30
	GramaticaParserPUNTO      = 31
	GramaticaParserPARA       = 32
	GramaticaParserPARC       = 33
	GramaticaParserLLABRE     = 34
	GramaticaParserLLACIE     = 35
	GramaticaParserCABRE      = 36
	GramaticaParserCCIER      = 37
	GramaticaParserDPUNTOS    = 38
	GramaticaParserPCOMA      = 39
	GramaticaParserCOMA       = 40
	GramaticaParserCADENA     = 41
	GramaticaParserENTERO     = 42
	GramaticaParserDECIMAL    = 43
	GramaticaParserTRUE       = 44
	GramaticaParserFALSE      = 45
	GramaticaParserID         = 46
	GramaticaParserESPACIOS   = 47
	GramaticaParserCHAR       = 48
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_ini              = 0
	GramaticaParserRULE_funcionMetodo    = 1
	GramaticaParserRULE_cuerpoFunciones  = 2
	GramaticaParserRULE_declaraciones    = 3
	GramaticaParserRULE_listaExpresiones = 4
	GramaticaParserRULE_sentenceIf       = 5
	GramaticaParserRULE_sentenceElse     = 6
	GramaticaParserRULE_variable         = 7
	GramaticaParserRULE_identificadores  = 8
	GramaticaParserRULE_tipo             = 9
	GramaticaParserRULE_expresion        = 10
	GramaticaParserRULE_primitivos       = 11
)

// IIniContext is an interface to support dynamic dispatch.
type IIniContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIniContext differentiates from other interfaces.
	IsIniContext()
}

type IniContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIniContext() *IniContext {
	var p = new(IniContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_ini
	return p
}

func (*IniContext) IsIniContext() {}

func NewIniContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IniContext {
	var p = new(IniContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_ini

	return p
}

func (s *IniContext) GetParser() antlr.Parser { return s.parser }

func (s *IniContext) FuncionMetodo() IFuncionMetodoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionMetodoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionMetodoContext)
}

func (s *IniContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *IniContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IniContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IniContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIni(s)
	}
}

func (s *IniContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIni(s)
	}
}

func (p *GramaticaParser) Ini() (localctx IIniContext) {
	localctx = NewIniContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_ini)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.FuncionMetodo()
	}
	{
		p.SetState(25)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IFuncionMetodoContext is an interface to support dynamic dispatch.
type IFuncionMetodoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncionMetodoContext differentiates from other interfaces.
	IsFuncionMetodoContext()
}

type FuncionMetodoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionMetodoContext() *FuncionMetodoContext {
	var p = new(FuncionMetodoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_funcionMetodo
	return p
}

func (*FuncionMetodoContext) IsFuncionMetodoContext() {}

func NewFuncionMetodoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionMetodoContext {
	var p = new(FuncionMetodoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_funcionMetodo

	return p
}

func (s *FuncionMetodoContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionMetodoContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPUBLIC, 0)
}

func (s *FuncionMetodoContext) MAIN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAIN, 0)
}

func (s *FuncionMetodoContext) AllDPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserDPUNTOS)
}

func (s *FuncionMetodoContext) DPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserDPUNTOS, i)
}

func (s *FuncionMetodoContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *FuncionMetodoContext) AllCABRE() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserCABRE)
}

func (s *FuncionMetodoContext) CABRE(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserCABRE, i)
}

func (s *FuncionMetodoContext) PRINCIPAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRINCIPAL, 0)
}

func (s *FuncionMetodoContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *FuncionMetodoContext) STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTRING, 0)
}

func (s *FuncionMetodoContext) LLABRE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLLABRE, 0)
}

func (s *FuncionMetodoContext) LLACIE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLLACIE, 0)
}

func (s *FuncionMetodoContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *FuncionMetodoContext) METODO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMETODO, 0)
}

func (s *FuncionMetodoContext) CuerpoFunciones() ICuerpoFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuerpoFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuerpoFuncionesContext)
}

func (s *FuncionMetodoContext) AllCCIER() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserCCIER)
}

func (s *FuncionMetodoContext) CCIER(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserCCIER, i)
}

func (s *FuncionMetodoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionMetodoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionMetodoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFuncionMetodo(s)
	}
}

func (s *FuncionMetodoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFuncionMetodo(s)
	}
}

func (p *GramaticaParser) FuncionMetodo() (localctx IFuncionMetodoContext) {
	localctx = NewFuncionMetodoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_funcionMetodo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(GramaticaParserPUBLIC)
	}
	{
		p.SetState(28)
		p.Match(GramaticaParserMAIN)
	}
	{
		p.SetState(29)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(30)
		p.Match(GramaticaParserID)
	}
	{
		p.SetState(31)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(32)
		p.Match(GramaticaParserPRINCIPAL)
	}
	{
		p.SetState(33)
		p.Match(GramaticaParserPARA)
	}
	{
		p.SetState(34)
		p.Match(GramaticaParserSTRING)
	}
	{
		p.SetState(35)
		p.Match(GramaticaParserLLABRE)
	}
	{
		p.SetState(36)
		p.Match(GramaticaParserLLACIE)
	}
	{
		p.SetState(37)
		p.Match(GramaticaParserPARC)
	}
	{
		p.SetState(38)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(39)
		p.Match(GramaticaParserMETODO)
	}
	{
		p.SetState(40)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(41)
		p.cuerpoFunciones(0)
	}
	{
		p.SetState(42)
		p.Match(GramaticaParserCCIER)
	}
	{
		p.SetState(43)
		p.Match(GramaticaParserCCIER)
	}

	return localctx
}

// ICuerpoFuncionesContext is an interface to support dynamic dispatch.
type ICuerpoFuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCuerpoFuncionesContext differentiates from other interfaces.
	IsCuerpoFuncionesContext()
}

type CuerpoFuncionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCuerpoFuncionesContext() *CuerpoFuncionesContext {
	var p = new(CuerpoFuncionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_cuerpoFunciones
	return p
}

func (*CuerpoFuncionesContext) IsCuerpoFuncionesContext() {}

func NewCuerpoFuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CuerpoFuncionesContext {
	var p = new(CuerpoFuncionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_cuerpoFunciones

	return p
}

func (s *CuerpoFuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *CuerpoFuncionesContext) Declaraciones() IDeclaracionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionesContext)
}

func (s *CuerpoFuncionesContext) CuerpoFunciones() ICuerpoFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuerpoFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuerpoFuncionesContext)
}

func (s *CuerpoFuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CuerpoFuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CuerpoFuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterCuerpoFunciones(s)
	}
}

func (s *CuerpoFuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitCuerpoFunciones(s)
	}
}

func (p *GramaticaParser) CuerpoFunciones() (localctx ICuerpoFuncionesContext) {
	return p.cuerpoFunciones(0)
}

func (p *GramaticaParser) cuerpoFunciones(_p int) (localctx ICuerpoFuncionesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCuerpoFuncionesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICuerpoFuncionesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_cuerpoFunciones, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Declaraciones()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCuerpoFuncionesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_cuerpoFunciones)
			p.SetState(48)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(49)
				p.Declaraciones()
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionesContext() *DeclaracionesContext {
	var p = new(DeclaracionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_declaraciones
	return p
}

func (*DeclaracionesContext) IsDeclaracionesContext() {}

func NewDeclaracionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionesContext {
	var p = new(DeclaracionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_declaraciones

	return p
}

func (s *DeclaracionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionesContext) SentenceIf() ISentenceIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceIfContext)
}

func (s *DeclaracionesContext) IMPRIMIR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIMPRIMIR, 0)
}

func (s *DeclaracionesContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *DeclaracionesContext) ListaExpresiones() IListaExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *DeclaracionesContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *DeclaracionesContext) PCOMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPCOMA, 0)
}

func (s *DeclaracionesContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DeclaracionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDeclaraciones(s)
	}
}

func (p *GramaticaParser) Declaraciones() (localctx IDeclaracionesContext) {
	localctx = NewDeclaracionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_declaraciones)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.SentenceIf()
		}

	case GramaticaParserIMPRIMIR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(GramaticaParserIMPRIMIR)
		}
		{
			p.SetState(57)
			p.Match(GramaticaParserPARA)
		}
		{
			p.SetState(58)
			p.listaExpresiones(0)
		}
		{
			p.SetState(59)
			p.Match(GramaticaParserPARC)
		}
		{
			p.SetState(60)
			p.Match(GramaticaParserPCOMA)
		}

	case GramaticaParserDECLARAR, GramaticaParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListaExpresionesContext is an interface to support dynamic dispatch.
type IListaExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListaExpresionesContext differentiates from other interfaces.
	IsListaExpresionesContext()
}

type ListaExpresionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaExpresionesContext() *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_listaExpresiones
	return p
}

func (*ListaExpresionesContext) IsListaExpresionesContext() {}

func NewListaExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_listaExpresiones

	return p
}

func (s *ListaExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExpresionesContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ListaExpresionesContext) ListaExpresiones() IListaExpresionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaExpresionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *ListaExpresionesContext) COMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCOMA, 0)
}

func (s *ListaExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterListaExpresiones(s)
	}
}

func (s *ListaExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitListaExpresiones(s)
	}
}

func (p *GramaticaParser) ListaExpresiones() (localctx IListaExpresionesContext) {
	return p.listaExpresiones(0)
}

func (p *GramaticaParser) listaExpresiones(_p int) (localctx IListaExpresionesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaExpresionesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaExpresionesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, GramaticaParserRULE_listaExpresiones, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.expresion(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaExpresionesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_listaExpresiones)
			p.SetState(68)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(69)
				p.Match(GramaticaParserCOMA)
			}
			{
				p.SetState(70)
				p.expresion(0)
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ISentenceIfContext is an interface to support dynamic dispatch.
type ISentenceIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenceIfContext differentiates from other interfaces.
	IsSentenceIfContext()
}

type SentenceIfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenceIfContext() *SentenceIfContext {
	var p = new(SentenceIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_sentenceIf
	return p
}

func (*SentenceIfContext) IsSentenceIfContext() {}

func NewSentenceIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceIfContext {
	var p = new(SentenceIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_sentenceIf

	return p
}

func (s *SentenceIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceIfContext) IF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIF, 0)
}

func (s *SentenceIfContext) DPUNTOS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDPUNTOS, 0)
}

func (s *SentenceIfContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *SentenceIfContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentenceIfContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *SentenceIfContext) CABRE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCABRE, 0)
}

func (s *SentenceIfContext) CuerpoFunciones() ICuerpoFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuerpoFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuerpoFuncionesContext)
}

func (s *SentenceIfContext) CCIER() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCCIER, 0)
}

func (s *SentenceIfContext) SentenceElse() ISentenceElseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceElseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceElseContext)
}

func (s *SentenceIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSentenceIf(s)
	}
}

func (s *SentenceIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSentenceIf(s)
	}
}

func (p *GramaticaParser) SentenceIf() (localctx ISentenceIfContext) {
	localctx = NewSentenceIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_sentenceIf)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(GramaticaParserIF)
	}
	{
		p.SetState(77)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(78)
		p.Match(GramaticaParserPARA)
	}
	{
		p.SetState(79)
		p.expresion(0)
	}
	{
		p.SetState(80)
		p.Match(GramaticaParserPARC)
	}
	{
		p.SetState(81)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(82)
		p.cuerpoFunciones(0)
	}
	{
		p.SetState(83)
		p.Match(GramaticaParserCCIER)
	}
	{
		p.SetState(84)
		p.SentenceElse()
	}

	return localctx
}

// ISentenceElseContext is an interface to support dynamic dispatch.
type ISentenceElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenceElseContext differentiates from other interfaces.
	IsSentenceElseContext()
}

type SentenceElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenceElseContext() *SentenceElseContext {
	var p = new(SentenceElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_sentenceElse
	return p
}

func (*SentenceElseContext) IsSentenceElseContext() {}

func NewSentenceElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceElseContext {
	var p = new(SentenceElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_sentenceElse

	return p
}

func (s *SentenceElseContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserELSE, 0)
}

func (s *SentenceElseContext) CABRE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCABRE, 0)
}

func (s *SentenceElseContext) CuerpoFunciones() ICuerpoFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuerpoFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuerpoFuncionesContext)
}

func (s *SentenceElseContext) CCIER() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCCIER, 0)
}

func (s *SentenceElseContext) SentenceIf() ISentenceIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceIfContext)
}

func (s *SentenceElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSentenceElse(s)
	}
}

func (s *SentenceElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSentenceElse(s)
	}
}

func (p *GramaticaParser) SentenceElse() (localctx ISentenceElseContext) {
	localctx = NewSentenceElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_sentenceElse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(GramaticaParserELSE)
		}
		{
			p.SetState(87)
			p.Match(GramaticaParserCABRE)
		}
		{
			p.SetState(88)
			p.cuerpoFunciones(0)
		}
		{
			p.SetState(89)
			p.Match(GramaticaParserCCIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(GramaticaParserELSE)
		}
		{
			p.SetState(92)
			p.SentenceIf()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) DECLARAR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDECLARAR, 0)
}

func (s *VariableContext) Identificadores() IIdentificadoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentificadoresContext)
}

func (s *VariableContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VariableContext) PCOMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPCOMA, 0)
}

func (s *VariableContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIGUAL, 0)
}

func (s *VariableContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *GramaticaParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(GramaticaParserDECLARAR)
		}
		{
			p.SetState(97)
			p.identificadores(0)
		}
		{
			p.SetState(98)
			p.Tipo()
		}
		{
			p.SetState(99)
			p.Match(GramaticaParserPCOMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(GramaticaParserDECLARAR)
		}
		{
			p.SetState(102)
			p.identificadores(0)
		}
		{
			p.SetState(103)
			p.Tipo()
		}
		{
			p.SetState(104)
			p.Match(GramaticaParserIGUAL)
		}
		{
			p.SetState(105)
			p.expresion(0)
		}
		{
			p.SetState(106)
			p.Match(GramaticaParserPCOMA)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.identificadores(0)
		}
		{
			p.SetState(109)
			p.Match(GramaticaParserIGUAL)
		}
		{
			p.SetState(110)
			p.expresion(0)
		}
		{
			p.SetState(111)
			p.Match(GramaticaParserPCOMA)
		}

	}

	return localctx
}

// IIdentificadoresContext is an interface to support dynamic dispatch.
type IIdentificadoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentificadoresContext differentiates from other interfaces.
	IsIdentificadoresContext()
}

type IdentificadoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentificadoresContext() *IdentificadoresContext {
	var p = new(IdentificadoresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_identificadores
	return p
}

func (*IdentificadoresContext) IsIdentificadoresContext() {}

func NewIdentificadoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificadoresContext {
	var p = new(IdentificadoresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_identificadores

	return p
}

func (s *IdentificadoresContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificadoresContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *IdentificadoresContext) AllIdentificadores() []IIdentificadoresContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem())
	var tst = make([]IIdentificadoresContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentificadoresContext)
		}
	}

	return tst
}

func (s *IdentificadoresContext) Identificadores(i int) IIdentificadoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentificadoresContext)
}

func (s *IdentificadoresContext) COMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCOMA, 0)
}

func (s *IdentificadoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificadoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificadoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIdentificadores(s)
	}
}

func (s *IdentificadoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIdentificadores(s)
	}
}

func (p *GramaticaParser) Identificadores() (localctx IIdentificadoresContext) {
	return p.identificadores(0)
}

func (p *GramaticaParser) identificadores(_p int) (localctx IIdentificadoresContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIdentificadoresContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIdentificadoresContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, GramaticaParserRULE_identificadores, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(GramaticaParserID)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentificadoresContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_identificadores)
			p.SetState(118)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(119)
				p.Match(GramaticaParserCOMA)
			}
			{
				p.SetState(120)
				p.identificadores(3)
			}

		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserINT, 0)
}

func (s *TipoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserBOOLEAN, 0)
}

func (s *TipoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDOUBLE, 0)
}

func (s *TipoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCHAR, 0)
}

func (s *TipoContext) STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTRING, 0)
}

func (s *TipoContext) REAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserREAL, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *GramaticaParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_tipo)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserDOUBLE)|(1<<GramaticaParserSTRING)|(1<<GramaticaParserINT)|(1<<GramaticaParserREAL)|(1<<GramaticaParserBOOLEAN))) != 0) || _la == GramaticaParserCHAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) MENOS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMENOS, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *ExpresionContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ExpresionContext) Primitivos() IPrimitivosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivosContext)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(GramaticaParserAND, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserOR, 0)
}

func (s *ExpresionContext) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *ExpresionContext) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *ExpresionContext) DIVIDIR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIVIDIR, 0)
}

func (s *ExpresionContext) MOD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMOD, 0)
}

func (s *ExpresionContext) IGUALI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIGUALI, 0)
}

func (s *ExpresionContext) DIFERENCIA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIFERENCIA, 0)
}

func (s *ExpresionContext) MAYI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAYI, 0)
}

func (s *ExpresionContext) MENI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMENI, 0)
}

func (s *ExpresionContext) MAY() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAY, 0)
}

func (s *ExpresionContext) MEN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEN, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *GramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *GramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, GramaticaParserRULE_expresion, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserMENOS:
		{
			p.SetState(129)
			p.Match(GramaticaParserMENOS)
		}
		{
			p.SetState(130)
			p.expresion(17)
		}

	case GramaticaParserPARA:
		{
			p.SetState(131)
			p.Match(GramaticaParserPARA)
		}
		{
			p.SetState(132)
			p.expresion(0)
		}
		{
			p.SetState(133)
			p.Match(GramaticaParserPARC)
		}

	case GramaticaParserID:
		{
			p.SetState(135)
			p.Match(GramaticaParserID)
		}

	case GramaticaParserCADENA, GramaticaParserENTERO, GramaticaParserDECIMAL, GramaticaParserTRUE, GramaticaParserFALSE:
		{
			p.SetState(136)
			p.Primitivos()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(140)
					p.Match(GramaticaParserAND)
				}
				{
					p.SetState(141)
					p.expresion(17)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(143)
					p.Match(GramaticaParserOR)
				}
				{
					p.SetState(144)
					p.expresion(16)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(146)
					p.Match(GramaticaParserMAS)
				}
				{
					p.SetState(147)
					p.expresion(15)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(149)
					p.Match(GramaticaParserMENOS)
				}
				{
					p.SetState(150)
					p.expresion(14)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(152)
					p.Match(GramaticaParserPOR)
				}
				{
					p.SetState(153)
					p.expresion(13)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(155)
					p.Match(GramaticaParserDIVIDIR)
				}
				{
					p.SetState(156)
					p.expresion(12)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(158)
					p.Match(GramaticaParserMOD)
				}
				{
					p.SetState(159)
					p.expresion(11)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(161)
					p.Match(GramaticaParserIGUALI)
				}
				{
					p.SetState(162)
					p.expresion(9)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(164)
					p.Match(GramaticaParserDIFERENCIA)
				}
				{
					p.SetState(165)
					p.expresion(8)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(167)
					p.Match(GramaticaParserMAYI)
				}
				{
					p.SetState(168)
					p.expresion(7)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(170)
					p.Match(GramaticaParserMENI)
				}
				{
					p.SetState(171)
					p.expresion(6)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(173)
					p.Match(GramaticaParserMAY)
				}
				{
					p.SetState(174)
					p.expresion(5)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(176)
					p.Match(GramaticaParserMEN)
				}
				{
					p.SetState(177)
					p.expresion(4)
				}

			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivosContext is an interface to support dynamic dispatch.
type IPrimitivosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitivosContext differentiates from other interfaces.
	IsPrimitivosContext()
}

type PrimitivosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitivosContext() *PrimitivosContext {
	var p = new(PrimitivosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_primitivos
	return p
}

func (*PrimitivosContext) IsPrimitivosContext() {}

func NewPrimitivosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivosContext {
	var p = new(PrimitivosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_primitivos

	return p
}

func (s *PrimitivosContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivosContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserENTERO, 0)
}

func (s *PrimitivosContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDECIMAL, 0)
}

func (s *PrimitivosContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTRUE, 0)
}

func (s *PrimitivosContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFALSE, 0)
}

func (s *PrimitivosContext) CADENA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCADENA, 0)
}

func (s *PrimitivosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterPrimitivos(s)
	}
}

func (s *PrimitivosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitPrimitivos(s)
	}
}

func (p *GramaticaParser) Primitivos() (localctx IPrimitivosContext) {
	localctx = NewPrimitivosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramaticaParserRULE_primitivos)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(GramaticaParserCADENA-41))|(1<<(GramaticaParserENTERO-41))|(1<<(GramaticaParserDECIMAL-41))|(1<<(GramaticaParserTRUE-41))|(1<<(GramaticaParserFALSE-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *CuerpoFuncionesContext = nil
		if localctx != nil {
			t = localctx.(*CuerpoFuncionesContext)
		}
		return p.CuerpoFunciones_Sempred(t, predIndex)

	case 4:
		var t *ListaExpresionesContext = nil
		if localctx != nil {
			t = localctx.(*ListaExpresionesContext)
		}
		return p.ListaExpresiones_Sempred(t, predIndex)

	case 8:
		var t *IdentificadoresContext = nil
		if localctx != nil {
			t = localctx.(*IdentificadoresContext)
		}
		return p.Identificadores_Sempred(t, predIndex)

	case 10:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) CuerpoFunciones_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) ListaExpresiones_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Identificadores_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
