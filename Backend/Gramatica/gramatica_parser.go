// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica // Gramatica
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2_TAREA1/AST/Entornos"
import "OLC2_TAREA1/AST/Interfaces"
import "OLC2_TAREA1/AST/Expresion"
import "OLC2_TAREA1/AST/Funcion"
import arrayList "github.com/colegno/arraylist"
import "OLC2_TAREA1/AST/Instruccion"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 228,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7,
	3, 49, 10, 3, 12, 3, 14, 3, 52, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 69, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 93,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 115, 10, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 125, 10, 8, 12, 8,
	14, 8, 128, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	138, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 149, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 163, 10, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 174, 10, 12, 12, 12,
	14, 12, 177, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 192, 10, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 209, 10, 13, 12, 13, 14, 13, 212, 11, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 5, 14, 226, 10, 14, 3, 14, 2, 5, 14, 22, 24, 15, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 2, 6, 3, 2, 28, 29, 3, 2, 22, 27, 3, 2, 32,
	33, 3, 2, 34, 35, 2, 240, 2, 28, 3, 2, 2, 2, 4, 50, 3, 2, 2, 2, 6, 68,
	3, 2, 2, 2, 8, 70, 3, 2, 2, 2, 10, 92, 3, 2, 2, 2, 12, 114, 3, 2, 2, 2,
	14, 116, 3, 2, 2, 2, 16, 137, 3, 2, 2, 2, 18, 148, 3, 2, 2, 2, 20, 162,
	3, 2, 2, 2, 22, 164, 3, 2, 2, 2, 24, 191, 3, 2, 2, 2, 26, 225, 3, 2, 2,
	2, 28, 29, 7, 7, 2, 2, 29, 30, 7, 6, 2, 2, 30, 31, 7, 45, 2, 2, 31, 32,
	7, 21, 2, 2, 32, 33, 7, 43, 2, 2, 33, 34, 7, 14, 2, 2, 34, 35, 7, 39, 2,
	2, 35, 36, 7, 10, 2, 2, 36, 37, 7, 41, 2, 2, 37, 38, 7, 42, 2, 2, 38, 39,
	7, 40, 2, 2, 39, 40, 7, 45, 2, 2, 40, 41, 7, 15, 2, 2, 41, 42, 7, 43, 2,
	2, 42, 43, 5, 4, 3, 2, 43, 44, 7, 44, 2, 2, 44, 45, 7, 44, 2, 2, 45, 46,
	8, 2, 1, 2, 46, 3, 3, 2, 2, 2, 47, 49, 5, 6, 4, 2, 48, 47, 3, 2, 2, 2,
	49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 53, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 8, 3, 1, 2, 54, 5, 3, 2, 2, 2, 55,
	56, 5, 8, 5, 2, 56, 57, 8, 4, 1, 2, 57, 69, 3, 2, 2, 2, 58, 59, 7, 5, 2,
	2, 59, 60, 7, 39, 2, 2, 60, 61, 5, 18, 10, 2, 61, 62, 7, 40, 2, 2, 62,
	63, 7, 46, 2, 2, 63, 64, 8, 4, 1, 2, 64, 69, 3, 2, 2, 2, 65, 66, 5, 12,
	7, 2, 66, 67, 8, 4, 1, 2, 67, 69, 3, 2, 2, 2, 68, 55, 3, 2, 2, 2, 68, 58,
	3, 2, 2, 2, 68, 65, 3, 2, 2, 2, 69, 7, 3, 2, 2, 2, 70, 71, 7, 3, 2, 2,
	71, 72, 7, 45, 2, 2, 72, 73, 7, 39, 2, 2, 73, 74, 5, 18, 10, 2, 74, 75,
	7, 40, 2, 2, 75, 76, 7, 43, 2, 2, 76, 77, 5, 4, 3, 2, 77, 78, 7, 44, 2,
	2, 78, 79, 5, 10, 6, 2, 79, 80, 8, 5, 1, 2, 80, 9, 3, 2, 2, 2, 81, 82,
	7, 4, 2, 2, 82, 83, 7, 43, 2, 2, 83, 84, 5, 4, 3, 2, 84, 85, 7, 44, 2,
	2, 85, 86, 8, 6, 1, 2, 86, 93, 3, 2, 2, 2, 87, 88, 7, 4, 2, 2, 88, 89,
	5, 8, 5, 2, 89, 90, 8, 6, 1, 2, 90, 93, 3, 2, 2, 2, 91, 93, 8, 6, 1, 2,
	92, 81, 3, 2, 2, 2, 92, 87, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 11, 3,
	2, 2, 2, 94, 95, 7, 9, 2, 2, 95, 96, 5, 14, 8, 2, 96, 97, 5, 16, 9, 2,
	97, 98, 7, 46, 2, 2, 98, 99, 8, 7, 1, 2, 99, 115, 3, 2, 2, 2, 100, 101,
	7, 9, 2, 2, 101, 102, 5, 14, 8, 2, 102, 103, 5, 16, 9, 2, 103, 104, 7,
	36, 2, 2, 104, 105, 5, 18, 10, 2, 105, 106, 7, 46, 2, 2, 106, 107, 8, 7,
	1, 2, 107, 115, 3, 2, 2, 2, 108, 109, 5, 14, 8, 2, 109, 110, 7, 36, 2,
	2, 110, 111, 5, 18, 10, 2, 111, 112, 7, 46, 2, 2, 112, 113, 8, 7, 1, 2,
	113, 115, 3, 2, 2, 2, 114, 94, 3, 2, 2, 2, 114, 100, 3, 2, 2, 2, 114, 108,
	3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 117, 8, 8, 1, 2, 117, 118, 7, 21,
	2, 2, 118, 119, 8, 8, 1, 2, 119, 126, 3, 2, 2, 2, 120, 121, 12, 4, 2, 2,
	121, 122, 7, 47, 2, 2, 122, 123, 7, 21, 2, 2, 123, 125, 8, 8, 1, 2, 124,
	120, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 15, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 11,
	2, 2, 130, 138, 8, 9, 1, 2, 131, 132, 7, 13, 2, 2, 132, 138, 8, 9, 1, 2,
	133, 134, 7, 10, 2, 2, 134, 138, 8, 9, 1, 2, 135, 136, 7, 12, 2, 2, 136,
	138, 8, 9, 1, 2, 137, 129, 3, 2, 2, 2, 137, 131, 3, 2, 2, 2, 137, 133,
	3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 17, 3, 2, 2, 2, 139, 140, 5, 20,
	11, 2, 140, 141, 8, 10, 1, 2, 141, 149, 3, 2, 2, 2, 142, 143, 5, 22, 12,
	2, 143, 144, 8, 10, 1, 2, 144, 149, 3, 2, 2, 2, 145, 146, 5, 24, 13, 2,
	146, 147, 8, 10, 1, 2, 147, 149, 3, 2, 2, 2, 148, 139, 3, 2, 2, 2, 148,
	142, 3, 2, 2, 2, 148, 145, 3, 2, 2, 2, 149, 19, 3, 2, 2, 2, 150, 151, 7,
	30, 2, 2, 151, 152, 5, 22, 12, 2, 152, 153, 8, 11, 1, 2, 153, 163, 3, 2,
	2, 2, 154, 155, 5, 22, 12, 2, 155, 156, 9, 2, 2, 2, 156, 157, 5, 22, 12,
	2, 157, 158, 8, 11, 1, 2, 158, 163, 3, 2, 2, 2, 159, 160, 5, 22, 12, 2,
	160, 161, 8, 11, 1, 2, 161, 163, 3, 2, 2, 2, 162, 150, 3, 2, 2, 2, 162,
	154, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 163, 21, 3, 2, 2, 2, 164, 165, 8,
	12, 1, 2, 165, 166, 5, 24, 13, 2, 166, 167, 8, 12, 1, 2, 167, 175, 3, 2,
	2, 2, 168, 169, 12, 4, 2, 2, 169, 170, 9, 3, 2, 2, 170, 171, 5, 22, 12,
	5, 171, 172, 8, 12, 1, 2, 172, 174, 3, 2, 2, 2, 173, 168, 3, 2, 2, 2, 174,
	177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 23, 3,
	2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 8, 13, 1, 2, 179, 180, 7, 35,
	2, 2, 180, 181, 5, 24, 13, 8, 181, 182, 8, 13, 1, 2, 182, 192, 3, 2, 2,
	2, 183, 184, 5, 26, 14, 2, 184, 185, 8, 13, 1, 2, 185, 192, 3, 2, 2, 2,
	186, 187, 7, 39, 2, 2, 187, 188, 5, 18, 10, 2, 188, 189, 7, 40, 2, 2, 189,
	190, 8, 13, 1, 2, 190, 192, 3, 2, 2, 2, 191, 178, 3, 2, 2, 2, 191, 183,
	3, 2, 2, 2, 191, 186, 3, 2, 2, 2, 192, 210, 3, 2, 2, 2, 193, 194, 12, 7,
	2, 2, 194, 195, 7, 31, 2, 2, 195, 196, 5, 24, 13, 8, 196, 197, 8, 13, 1,
	2, 197, 209, 3, 2, 2, 2, 198, 199, 12, 6, 2, 2, 199, 200, 9, 4, 2, 2, 200,
	201, 5, 24, 13, 7, 201, 202, 8, 13, 1, 2, 202, 209, 3, 2, 2, 2, 203, 204,
	12, 5, 2, 2, 204, 205, 9, 5, 2, 2, 205, 206, 5, 24, 13, 6, 206, 207, 8,
	13, 1, 2, 207, 209, 3, 2, 2, 2, 208, 193, 3, 2, 2, 2, 208, 198, 3, 2, 2,
	2, 208, 203, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210,
	211, 3, 2, 2, 2, 211, 25, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 7,
	17, 2, 2, 214, 226, 8, 14, 1, 2, 215, 216, 7, 18, 2, 2, 216, 226, 8, 14,
	1, 2, 217, 218, 7, 19, 2, 2, 218, 226, 8, 14, 1, 2, 219, 220, 7, 20, 2,
	2, 220, 226, 8, 14, 1, 2, 221, 222, 7, 16, 2, 2, 222, 226, 8, 14, 1, 2,
	223, 224, 7, 21, 2, 2, 224, 226, 8, 14, 1, 2, 225, 213, 3, 2, 2, 2, 225,
	215, 3, 2, 2, 2, 225, 217, 3, 2, 2, 2, 225, 219, 3, 2, 2, 2, 225, 221,
	3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 27, 3, 2, 2, 2, 15, 50, 68, 92,
	114, 126, 137, 148, 162, 175, 191, 208, 210, 225,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "'<='", "'>='", "'<'", "'>'", "'=='", "'!='", "'&&'", "'||'", "'!'",
	"'%'", "'*'", "'/'", "'+'", "'-'", "'='", "'_'", "'.'", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "':'", "';'", "','",
}
var symbolicNames = []string{
	"", "IF", "ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", "STRING",
	"INT", "REAL", "BOOLEAN", "PRINCIPAL", "METODO", "CADENA", "ENTERO", "DECIMAL",
	"TRUE", "FALSE", "ID", "MENI", "MAYI", "MEN", "MAY", "IGUALI", "DIFERENCIA",
	"AND", "OR", "NOT", "MOD", "POR", "DIVIDIR", "MAS", "MENOS", "IGUAL", "GUIONB",
	"PUNTO", "PARA", "PARC", "LLABRE", "LLACIE", "CABRE", "CCIER", "DPUNTOS",
	"PCOMA", "COMA", "ESPACIOS", "COMENTARIO_MUL", "COMENTARIO_LIN",
}

var ruleNames = []string{
	"ini", "instrucciones", "declaraciones", "sentenceIf", "sentenceElse",
	"variable", "identificadores", "tipo", "expresion", "logicas", "relacionales",
	"aritmeticas", "primitivos",
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
	GramaticaParserEOF            = antlr.TokenEOF
	GramaticaParserIF             = 1
	GramaticaParserELSE           = 2
	GramaticaParserIMPRIMIR       = 3
	GramaticaParserMAIN           = 4
	GramaticaParserPUBLIC         = 5
	GramaticaParserCLASS          = 6
	GramaticaParserDECLARAR       = 7
	GramaticaParserSTRING         = 8
	GramaticaParserINT            = 9
	GramaticaParserREAL           = 10
	GramaticaParserBOOLEAN        = 11
	GramaticaParserPRINCIPAL      = 12
	GramaticaParserMETODO         = 13
	GramaticaParserCADENA         = 14
	GramaticaParserENTERO         = 15
	GramaticaParserDECIMAL        = 16
	GramaticaParserTRUE           = 17
	GramaticaParserFALSE          = 18
	GramaticaParserID             = 19
	GramaticaParserMENI           = 20
	GramaticaParserMAYI           = 21
	GramaticaParserMEN            = 22
	GramaticaParserMAY            = 23
	GramaticaParserIGUALI         = 24
	GramaticaParserDIFERENCIA     = 25
	GramaticaParserAND            = 26
	GramaticaParserOR             = 27
	GramaticaParserNOT            = 28
	GramaticaParserMOD            = 29
	GramaticaParserPOR            = 30
	GramaticaParserDIVIDIR        = 31
	GramaticaParserMAS            = 32
	GramaticaParserMENOS          = 33
	GramaticaParserIGUAL          = 34
	GramaticaParserGUIONB         = 35
	GramaticaParserPUNTO          = 36
	GramaticaParserPARA           = 37
	GramaticaParserPARC           = 38
	GramaticaParserLLABRE         = 39
	GramaticaParserLLACIE         = 40
	GramaticaParserCABRE          = 41
	GramaticaParserCCIER          = 42
	GramaticaParserDPUNTOS        = 43
	GramaticaParserPCOMA          = 44
	GramaticaParserCOMA           = 45
	GramaticaParserESPACIOS       = 46
	GramaticaParserCOMENTARIO_MUL = 47
	GramaticaParserCOMENTARIO_LIN = 48
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_ini             = 0
	GramaticaParserRULE_instrucciones   = 1
	GramaticaParserRULE_declaraciones   = 2
	GramaticaParserRULE_sentenceIf      = 3
	GramaticaParserRULE_sentenceElse    = 4
	GramaticaParserRULE_variable        = 5
	GramaticaParserRULE_identificadores = 6
	GramaticaParserRULE_tipo            = 7
	GramaticaParserRULE_expresion       = 8
	GramaticaParserRULE_logicas         = 9
	GramaticaParserRULE_relacionales    = 10
	GramaticaParserRULE_aritmeticas     = 11
	GramaticaParserRULE_primitivos      = 12
)

// IIniContext is an interface to support dynamic dispatch.
type IIniContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsIniContext differentiates from other interfaces.
	IsIniContext()
}

type IniContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
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

func (s *IniContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *IniContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *IniContext) GetLista() *arrayList.List { return s.lista }

func (s *IniContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *IniContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPUBLIC, 0)
}

func (s *IniContext) MAIN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAIN, 0)
}

func (s *IniContext) AllDPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserDPUNTOS)
}

func (s *IniContext) DPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserDPUNTOS, i)
}

func (s *IniContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *IniContext) AllCABRE() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserCABRE)
}

func (s *IniContext) CABRE(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserCABRE, i)
}

func (s *IniContext) PRINCIPAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRINCIPAL, 0)
}

func (s *IniContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *IniContext) STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTRING, 0)
}

func (s *IniContext) LLABRE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLLABRE, 0)
}

func (s *IniContext) LLACIE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLLACIE, 0)
}

func (s *IniContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *IniContext) METODO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMETODO, 0)
}

func (s *IniContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *IniContext) AllCCIER() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserCCIER)
}

func (s *IniContext) CCIER(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserCCIER, i)
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
		p.SetState(26)
		p.Match(GramaticaParserPUBLIC)
	}
	{
		p.SetState(27)
		p.Match(GramaticaParserMAIN)
	}
	{
		p.SetState(28)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(29)
		p.Match(GramaticaParserID)
	}
	{
		p.SetState(30)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(31)
		p.Match(GramaticaParserPRINCIPAL)
	}
	{
		p.SetState(32)
		p.Match(GramaticaParserPARA)
	}
	{
		p.SetState(33)
		p.Match(GramaticaParserSTRING)
	}
	{
		p.SetState(34)
		p.Match(GramaticaParserLLABRE)
	}
	{
		p.SetState(35)
		p.Match(GramaticaParserLLACIE)
	}
	{
		p.SetState(36)
		p.Match(GramaticaParserPARC)
	}
	{
		p.SetState(37)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(38)
		p.Match(GramaticaParserMETODO)
	}
	{
		p.SetState(39)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(40)

		var _x = p.Instrucciones()

		localctx.(*IniContext)._instrucciones = _x
	}
	{
		p.SetState(41)
		p.Match(GramaticaParserCCIER)
	}
	{
		p.SetState(42)
		p.Match(GramaticaParserCCIER)
	}
	localctx.(*IniContext).lista = localctx.(*IniContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_declaraciones returns the _declaraciones rule contexts.
	Get_declaraciones() IDeclaracionesContext

	// Set_declaraciones sets the _declaraciones rule contexts.
	Set_declaraciones(IDeclaracionesContext)

	// GetE returns the e rule context list.
	GetE() []IDeclaracionesContext

	// SetE sets the e rule context list.
	SetE([]IDeclaracionesContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	l              *arrayList.List
	_declaraciones IDeclaracionesContext
	e              []IDeclaracionesContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_declaraciones() IDeclaracionesContext { return s._declaraciones }

func (s *InstruccionesContext) Set_declaraciones(v IDeclaracionesContext) { s._declaraciones = v }

func (s *InstruccionesContext) GetE() []IDeclaracionesContext { return s.e }

func (s *InstruccionesContext) SetE(v []IDeclaracionesContext) { s.e = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllDeclaraciones() []IDeclaracionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclaracionesContext)(nil)).Elem())
	var tst = make([]IDeclaracionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclaracionesContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Declaraciones(i int) IDeclaracionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionesContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *GramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserIF)|(1<<GramaticaParserIMPRIMIR)|(1<<GramaticaParserDECLARAR)|(1<<GramaticaParserID))) != 0 {
		{
			p.SetState(45)

			var _x = p.Declaraciones()

			localctx.(*InstruccionesContext)._declaraciones = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._declaraciones)

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_sentenceIf returns the _sentenceIf rule contexts.
	Get_sentenceIf() ISentenceIfContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_variable returns the _variable rule contexts.
	Get_variable() IVariableContext

	// Set_sentenceIf sets the _sentenceIf rule contexts.
	Set_sentenceIf(ISentenceIfContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_variable sets the _variable rule contexts.
	Set_variable(IVariableContext)

	// GetInstr returns the instr attribute.
	GetInstr() Interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(Interfaces.Instruccion)

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Interfaces.Instruccion
	_sentenceIf ISentenceIfContext
	_expresion  IExpresionContext
	_variable   IVariableContext
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

func (s *DeclaracionesContext) Get_sentenceIf() ISentenceIfContext { return s._sentenceIf }

func (s *DeclaracionesContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionesContext) Get_variable() IVariableContext { return s._variable }

func (s *DeclaracionesContext) Set_sentenceIf(v ISentenceIfContext) { s._sentenceIf = v }

func (s *DeclaracionesContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionesContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *DeclaracionesContext) GetInstr() Interfaces.Instruccion { return s.instr }

func (s *DeclaracionesContext) SetInstr(v Interfaces.Instruccion) { s.instr = v }

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

func (s *DeclaracionesContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
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
	p.EnterRule(localctx, 4, GramaticaParserRULE_declaraciones)

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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)

			var _x = p.SentenceIf()

			localctx.(*DeclaracionesContext)._sentenceIf = _x
		}
		localctx.(*DeclaracionesContext).instr = localctx.(*DeclaracionesContext).Get_sentenceIf().GetInstr()

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

			var _x = p.Expresion()

			localctx.(*DeclaracionesContext)._expresion = _x
		}
		{
			p.SetState(59)
			p.Match(GramaticaParserPARC)
		}
		{
			p.SetState(60)
			p.Match(GramaticaParserPCOMA)
		}
		localctx.(*DeclaracionesContext).instr = Funcion.NewImprimir(localctx.(*DeclaracionesContext).Get_expresion().GetP())

	case GramaticaParserDECLARAR, GramaticaParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)

			var _x = p.Variable()

			localctx.(*DeclaracionesContext)._variable = _x
		}
		localctx.(*DeclaracionesContext).instr = localctx.(*DeclaracionesContext).Get_variable().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISentenceIfContext is an interface to support dynamic dispatch.
type ISentenceIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_sentenceElse returns the _sentenceElse rule contexts.
	Get_sentenceElse() ISentenceElseContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_sentenceElse sets the _sentenceElse rule contexts.
	Set_sentenceElse(ISentenceElseContext)

	// GetInstr returns the instr attribute.
	GetInstr() Interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(Interfaces.Instruccion)

	// IsSentenceIfContext differentiates from other interfaces.
	IsSentenceIfContext()
}

type SentenceIfContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          Interfaces.Instruccion
	_IF            antlr.Token
	_expresion     IExpresionContext
	_instrucciones IInstruccionesContext
	_sentenceElse  ISentenceElseContext
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

func (s *SentenceIfContext) Get_IF() antlr.Token { return s._IF }

func (s *SentenceIfContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *SentenceIfContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *SentenceIfContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *SentenceIfContext) Get_sentenceElse() ISentenceElseContext { return s._sentenceElse }

func (s *SentenceIfContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *SentenceIfContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *SentenceIfContext) Set_sentenceElse(v ISentenceElseContext) { s._sentenceElse = v }

func (s *SentenceIfContext) GetInstr() Interfaces.Instruccion { return s.instr }

func (s *SentenceIfContext) SetInstr(v Interfaces.Instruccion) { s.instr = v }

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

func (s *SentenceIfContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
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
	p.EnterRule(localctx, 6, GramaticaParserRULE_sentenceIf)

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
		p.SetState(68)

		var _m = p.Match(GramaticaParserIF)

		localctx.(*SentenceIfContext)._IF = _m
	}
	{
		p.SetState(69)
		p.Match(GramaticaParserDPUNTOS)
	}
	{
		p.SetState(70)
		p.Match(GramaticaParserPARA)
	}
	{
		p.SetState(71)

		var _x = p.Expresion()

		localctx.(*SentenceIfContext)._expresion = _x
	}
	{
		p.SetState(72)
		p.Match(GramaticaParserPARC)
	}
	{
		p.SetState(73)
		p.Match(GramaticaParserCABRE)
	}
	{
		p.SetState(74)

		var _x = p.Instrucciones()

		localctx.(*SentenceIfContext)._instrucciones = _x
	}
	{
		p.SetState(75)
		p.Match(GramaticaParserCCIER)
	}
	{
		p.SetState(76)

		var _x = p.SentenceElse()

		localctx.(*SentenceIfContext)._sentenceElse = _x
	}
	localctx.(*SentenceIfContext).instr = Funcion.NewIf((func() int {
		if localctx.(*SentenceIfContext).Get_IF() == nil {
			return 0
		} else {
			return localctx.(*SentenceIfContext).Get_IF().GetLine()
		}
	}()), localctx.(*SentenceIfContext).Get_IF().GetColumn(), localctx.(*SentenceIfContext).Get_expresion().GetP(), localctx.(*SentenceIfContext).Get_instrucciones().GetL(), localctx.(*SentenceIfContext).Get_sentenceElse().GetL())

	return localctx
}

// ISentenceElseContext is an interface to support dynamic dispatch.
type ISentenceElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_sentenceIf returns the _sentenceIf rule contexts.
	Get_sentenceIf() ISentenceIfContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_sentenceIf sets the _sentenceIf rule contexts.
	Set_sentenceIf(ISentenceIfContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsSentenceElseContext differentiates from other interfaces.
	IsSentenceElseContext()
}

type SentenceElseContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	l              *arrayList.List
	_instrucciones IInstruccionesContext
	_sentenceIf    ISentenceIfContext
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

func (s *SentenceElseContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *SentenceElseContext) Get_sentenceIf() ISentenceIfContext { return s._sentenceIf }

func (s *SentenceElseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *SentenceElseContext) Set_sentenceIf(v ISentenceIfContext) { s._sentenceIf = v }

func (s *SentenceElseContext) GetL() *arrayList.List { return s.l }

func (s *SentenceElseContext) SetL(v *arrayList.List) { s.l = v }

func (s *SentenceElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserELSE, 0)
}

func (s *SentenceElseContext) CABRE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCABRE, 0)
}

func (s *SentenceElseContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
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
	p.EnterRule(localctx, 8, GramaticaParserRULE_sentenceElse)
	localctx.(*SentenceElseContext).l = arrayList.New()

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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(GramaticaParserELSE)
		}
		{
			p.SetState(80)
			p.Match(GramaticaParserCABRE)
		}
		{
			p.SetState(81)

			var _x = p.Instrucciones()

			localctx.(*SentenceElseContext)._instrucciones = _x
		}
		{
			p.SetState(82)
			p.Match(GramaticaParserCCIER)
		}
		localctx.(*SentenceElseContext).l = localctx.(*SentenceElseContext).Get_instrucciones().GetL()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(GramaticaParserELSE)
		}
		{
			p.SetState(86)

			var _x = p.SentenceIf()

			localctx.(*SentenceElseContext)._sentenceIf = _x
		}
		localctx.(*SentenceElseContext).l.Add(localctx.(*SentenceElseContext).Get_sentenceIf().GetInstr())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		localctx.(*SentenceElseContext).l = nil

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_DECLARAR returns the _DECLARAR token.
	Get_DECLARAR() antlr.Token

	// Get_IGUAL returns the _IGUAL token.
	Get_IGUAL() antlr.Token

	// Set_DECLARAR sets the _DECLARAR token.
	Set_DECLARAR(antlr.Token)

	// Set_IGUAL sets the _IGUAL token.
	Set_IGUAL(antlr.Token)

	// Get_identificadores returns the _identificadores rule contexts.
	Get_identificadores() IIdentificadoresContext

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_identificadores sets the _identificadores rule contexts.
	Set_identificadores(IIdentificadoresContext)

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(Interfaces.Instruccion)

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            Interfaces.Instruccion
	_DECLARAR        antlr.Token
	_identificadores IIdentificadoresContext
	_tipo            ITipoContext
	_IGUAL           antlr.Token
	_expresion       IExpresionContext
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

func (s *VariableContext) Get_DECLARAR() antlr.Token { return s._DECLARAR }

func (s *VariableContext) Get_IGUAL() antlr.Token { return s._IGUAL }

func (s *VariableContext) Set_DECLARAR(v antlr.Token) { s._DECLARAR = v }

func (s *VariableContext) Set_IGUAL(v antlr.Token) { s._IGUAL = v }

func (s *VariableContext) Get_identificadores() IIdentificadoresContext { return s._identificadores }

func (s *VariableContext) Get_tipo() ITipoContext { return s._tipo }

func (s *VariableContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *VariableContext) Set_identificadores(v IIdentificadoresContext) { s._identificadores = v }

func (s *VariableContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *VariableContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *VariableContext) GetInstr() Interfaces.Instruccion { return s.instr }

func (s *VariableContext) SetInstr(v Interfaces.Instruccion) { s.instr = v }

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
	p.EnterRule(localctx, 10, GramaticaParserRULE_variable)

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)

			var _m = p.Match(GramaticaParserDECLARAR)

			localctx.(*VariableContext)._DECLARAR = _m
		}
		{
			p.SetState(93)

			var _x = p.identificadores(0)

			localctx.(*VariableContext)._identificadores = _x
		}
		{
			p.SetState(94)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(95)
			p.Match(GramaticaParserPCOMA)
		}
		localctx.(*VariableContext).instr = Instruccion.NuevaDeclaracion((func() int {
			if localctx.(*VariableContext).Get_DECLARAR() == nil {
				return 0
			} else {
				return localctx.(*VariableContext).Get_DECLARAR().GetLine()
			}
		}()), localctx.(*VariableContext).Get_DECLARAR().GetColumn(), localctx.(*VariableContext).Get_identificadores().GetLista(), localctx.(*VariableContext).Get_tipo().GetT())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(GramaticaParserDECLARAR)
		}
		{
			p.SetState(99)

			var _x = p.identificadores(0)

			localctx.(*VariableContext)._identificadores = _x
		}
		{
			p.SetState(100)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(101)

			var _m = p.Match(GramaticaParserIGUAL)

			localctx.(*VariableContext)._IGUAL = _m
		}
		{
			p.SetState(102)

			var _x = p.Expresion()

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(103)
			p.Match(GramaticaParserPCOMA)
		}
		localctx.(*VariableContext).instr = Instruccion.NuevaDeclaracionInicio((func() int {
			if localctx.(*VariableContext).Get_IGUAL() == nil {
				return 0
			} else {
				return localctx.(*VariableContext).Get_IGUAL().GetLine()
			}
		}()), localctx.(*VariableContext).Get_IGUAL().GetColumn(), localctx.(*VariableContext).Get_identificadores().GetLista(), localctx.(*VariableContext).Get_tipo().GetT(), localctx.(*VariableContext).Get_expresion().GetP())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)

			var _x = p.identificadores(0)

			localctx.(*VariableContext)._identificadores = _x
		}
		{
			p.SetState(107)

			var _m = p.Match(GramaticaParserIGUAL)

			localctx.(*VariableContext)._IGUAL = _m
		}
		{
			p.SetState(108)

			var _x = p.Expresion()

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(109)
			p.Match(GramaticaParserPCOMA)
		}
		localctx.(*VariableContext).instr = Instruccion.NuevaDeclaracionInicio((func() int {
			if localctx.(*VariableContext).Get_IGUAL() == nil {
				return 0
			} else {
				return localctx.(*VariableContext).Get_IGUAL().GetLine()
			}
		}()), localctx.(*VariableContext).Get_IGUAL().GetColumn(), localctx.(*VariableContext).Get_identificadores().GetLista(), Entornos.NULL, localctx.(*VariableContext).Get_expresion().GetP())

	}

	return localctx
}

// IIdentificadoresContext is an interface to support dynamic dispatch.
type IIdentificadoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetSub returns the sub rule contexts.
	GetSub() IIdentificadoresContext

	// SetSub sets the sub rule contexts.
	SetSub(IIdentificadoresContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsIdentificadoresContext differentiates from other interfaces.
	IsIdentificadoresContext()
}

type IdentificadoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayList.List
	sub    IIdentificadoresContext
	_ID    antlr.Token
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

func (s *IdentificadoresContext) Get_ID() antlr.Token { return s._ID }

func (s *IdentificadoresContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *IdentificadoresContext) GetSub() IIdentificadoresContext { return s.sub }

func (s *IdentificadoresContext) SetSub(v IIdentificadoresContext) { s.sub = v }

func (s *IdentificadoresContext) GetLista() *arrayList.List { return s.lista }

func (s *IdentificadoresContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *IdentificadoresContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *IdentificadoresContext) COMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCOMA, 0)
}

func (s *IdentificadoresContext) Identificadores() IIdentificadoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentificadoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentificadoresContext)
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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, GramaticaParserRULE_identificadores, _p)

	localctx.(*IdentificadoresContext).lista = arrayList.New()

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
		p.SetState(115)

		var _m = p.Match(GramaticaParserID)

		localctx.(*IdentificadoresContext)._ID = _m
	}
	localctx.(*IdentificadoresContext).lista.Add(Expresion.NuevoIdentificador((func() int {
		if localctx.(*IdentificadoresContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*IdentificadoresContext).Get_ID().GetLine()
		}
	}()), localctx.(*IdentificadoresContext).Get_ID().GetColumn(), (func() string {
		if localctx.(*IdentificadoresContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*IdentificadoresContext).Get_ID().GetText()
		}
	}())))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentificadoresContext(p, _parentctx, _parentState)
			localctx.(*IdentificadoresContext).sub = _prevctx
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

				var _m = p.Match(GramaticaParserID)

				localctx.(*IdentificadoresContext)._ID = _m
			}

			localctx.(*IdentificadoresContext).GetSub().GetLista().Add(Expresion.NuevoIdentificador((func() int {
				if localctx.(*IdentificadoresContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*IdentificadoresContext).Get_ID().GetLine()
				}
			}()), localctx.(*IdentificadoresContext).Get_ID().GetColumn(), (func() string {
				if localctx.(*IdentificadoresContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*IdentificadoresContext).Get_ID().GetText()
				}
			}())))
			localctx.(*IdentificadoresContext).lista = localctx.(*IdentificadoresContext).GetSub().GetLista()

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the T attribute.
	GetT() Entornos.TipoDato

	// SetT sets the T attribute.
	SetT(Entornos.TipoDato)

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	T      Entornos.TipoDato
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

func (s *TipoContext) GetT() Entornos.TipoDato { return s.T }

func (s *TipoContext) SetT(v Entornos.TipoDato) { s.T = v }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserINT, 0)
}

func (s *TipoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserBOOLEAN, 0)
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
	p.EnterRule(localctx, 14, GramaticaParserRULE_tipo)

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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(GramaticaParserINT)
		}
		localctx.(*TipoContext).T = Entornos.INTEGER

	case GramaticaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(GramaticaParserBOOLEAN)
		}
		localctx.(*TipoContext).T = Entornos.BOOLEAN

	case GramaticaParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(GramaticaParserSTRING)
		}
		localctx.(*TipoContext).T = Entornos.STRING

	case GramaticaParserREAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.Match(GramaticaParserREAL)
		}
		localctx.(*TipoContext).T = Entornos.REAL

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_logicas returns the _logicas rule contexts.
	Get_logicas() ILogicasContext

	// Get_relacionales returns the _relacionales rule contexts.
	Get_relacionales() IRelacionalesContext

	// Get_aritmeticas returns the _aritmeticas rule contexts.
	Get_aritmeticas() IAritmeticasContext

	// Set_logicas sets the _logicas rule contexts.
	Set_logicas(ILogicasContext)

	// Set_relacionales sets the _relacionales rule contexts.
	Set_relacionales(IRelacionalesContext)

	// Set_aritmeticas sets the _aritmeticas rule contexts.
	Set_aritmeticas(IAritmeticasContext)

	// GetP returns the p attribute.
	GetP() Interfaces.Expresion

	// SetP sets the p attribute.
	SetP(Interfaces.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	p             Interfaces.Expresion
	_logicas      ILogicasContext
	_relacionales IRelacionalesContext
	_aritmeticas  IAritmeticasContext
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

func (s *ExpresionContext) Get_logicas() ILogicasContext { return s._logicas }

func (s *ExpresionContext) Get_relacionales() IRelacionalesContext { return s._relacionales }

func (s *ExpresionContext) Get_aritmeticas() IAritmeticasContext { return s._aritmeticas }

func (s *ExpresionContext) Set_logicas(v ILogicasContext) { s._logicas = v }

func (s *ExpresionContext) Set_relacionales(v IRelacionalesContext) { s._relacionales = v }

func (s *ExpresionContext) Set_aritmeticas(v IAritmeticasContext) { s._aritmeticas = v }

func (s *ExpresionContext) GetP() Interfaces.Expresion { return s.p }

func (s *ExpresionContext) SetP(v Interfaces.Expresion) { s.p = v }

func (s *ExpresionContext) Logicas() ILogicasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicasContext)
}

func (s *ExpresionContext) Relacionales() IRelacionalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelacionalesContext)
}

func (s *ExpresionContext) Aritmeticas() IAritmeticasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAritmeticasContext)
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
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_expresion)

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)

			var _x = p.Logicas()

			localctx.(*ExpresionContext)._logicas = _x
		}
		localctx.(*ExpresionContext).p = localctx.(*ExpresionContext).Get_logicas().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)

			var _x = p.relacionales(0)

			localctx.(*ExpresionContext)._relacionales = _x
		}
		localctx.(*ExpresionContext).p = localctx.(*ExpresionContext).Get_relacionales().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)

			var _x = p.aritmeticas(0)

			localctx.(*ExpresionContext)._aritmeticas = _x
		}
		localctx.(*ExpresionContext).p = localctx.(*ExpresionContext).Get_aritmeticas().GetP()

	}

	return localctx
}

// ILogicasContext is an interface to support dynamic dispatch.
type ILogicasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IRelacionalesContext

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IRelacionalesContext

	// Get_relacionales returns the _relacionales rule contexts.
	Get_relacionales() IRelacionalesContext

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IRelacionalesContext)

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IRelacionalesContext)

	// Set_relacionales sets the _relacionales rule contexts.
	Set_relacionales(IRelacionalesContext)

	// GetP returns the p attribute.
	GetP() Interfaces.Expresion

	// SetP sets the p attribute.
	SetP(Interfaces.Expresion)

	// IsLogicasContext differentiates from other interfaces.
	IsLogicasContext()
}

type LogicasContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	p             Interfaces.Expresion
	op            antlr.Token
	opDe          IRelacionalesContext
	opIz          IRelacionalesContext
	_relacionales IRelacionalesContext
}

func NewEmptyLogicasContext() *LogicasContext {
	var p = new(LogicasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_logicas
	return p
}

func (*LogicasContext) IsLogicasContext() {}

func NewLogicasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicasContext {
	var p = new(LogicasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_logicas

	return p
}

func (s *LogicasContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicasContext) GetOp() antlr.Token { return s.op }

func (s *LogicasContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicasContext) GetOpDe() IRelacionalesContext { return s.opDe }

func (s *LogicasContext) GetOpIz() IRelacionalesContext { return s.opIz }

func (s *LogicasContext) Get_relacionales() IRelacionalesContext { return s._relacionales }

func (s *LogicasContext) SetOpDe(v IRelacionalesContext) { s.opDe = v }

func (s *LogicasContext) SetOpIz(v IRelacionalesContext) { s.opIz = v }

func (s *LogicasContext) Set_relacionales(v IRelacionalesContext) { s._relacionales = v }

func (s *LogicasContext) GetP() Interfaces.Expresion { return s.p }

func (s *LogicasContext) SetP(v Interfaces.Expresion) { s.p = v }

func (s *LogicasContext) NOT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNOT, 0)
}

func (s *LogicasContext) AllRelacionales() []IRelacionalesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelacionalesContext)(nil)).Elem())
	var tst = make([]IRelacionalesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelacionalesContext)
		}
	}

	return tst
}

func (s *LogicasContext) Relacionales(i int) IRelacionalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelacionalesContext)
}

func (s *LogicasContext) AND() antlr.TerminalNode {
	return s.GetToken(GramaticaParserAND, 0)
}

func (s *LogicasContext) OR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserOR, 0)
}

func (s *LogicasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterLogicas(s)
	}
}

func (s *LogicasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitLogicas(s)
	}
}

func (p *GramaticaParser) Logicas() (localctx ILogicasContext) {
	localctx = NewLogicasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_logicas)
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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)

			var _m = p.Match(GramaticaParserNOT)

			localctx.(*LogicasContext).op = _m
		}
		{
			p.SetState(149)

			var _x = p.relacionales(0)

			localctx.(*LogicasContext).opDe = _x
		}
		localctx.(*LogicasContext).p = Expresion.NewLogico((func() int {
			if localctx.(*LogicasContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*LogicasContext).GetOp().GetLine()
			}
		}()), localctx.(*LogicasContext).GetOp().GetColumn(), nil, (func() string {
			if localctx.(*LogicasContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*LogicasContext).GetOp().GetText()
			}
		}()), localctx.(*LogicasContext).GetOpDe().GetP(), false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)

			var _x = p.relacionales(0)

			localctx.(*LogicasContext).opIz = _x
		}
		{
			p.SetState(153)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LogicasContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramaticaParserAND || _la == GramaticaParserOR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LogicasContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(154)

			var _x = p.relacionales(0)

			localctx.(*LogicasContext).opDe = _x
		}
		localctx.(*LogicasContext).p = Expresion.NewLogico((func() int {
			if localctx.(*LogicasContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*LogicasContext).GetOp().GetLine()
			}
		}()), localctx.(*LogicasContext).GetOp().GetColumn(), localctx.(*LogicasContext).GetOpIz().GetP(), (func() string {
			if localctx.(*LogicasContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*LogicasContext).GetOp().GetText()
			}
		}()), localctx.(*LogicasContext).GetOpDe().GetP(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)

			var _x = p.relacionales(0)

			localctx.(*LogicasContext)._relacionales = _x
		}
		localctx.(*LogicasContext).p = localctx.(*LogicasContext).Get_relacionales().GetP()

	}

	return localctx
}

// IRelacionalesContext is an interface to support dynamic dispatch.
type IRelacionalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IRelacionalesContext

	// Get_aritmeticas returns the _aritmeticas rule contexts.
	Get_aritmeticas() IAritmeticasContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IRelacionalesContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IRelacionalesContext)

	// Set_aritmeticas sets the _aritmeticas rule contexts.
	Set_aritmeticas(IAritmeticasContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IRelacionalesContext)

	// GetP returns the p attribute.
	GetP() Interfaces.Expresion

	// SetP sets the p attribute.
	SetP(Interfaces.Expresion)

	// IsRelacionalesContext differentiates from other interfaces.
	IsRelacionalesContext()
}

type RelacionalesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            Interfaces.Expresion
	opIz         IRelacionalesContext
	_aritmeticas IAritmeticasContext
	op           antlr.Token
	opDe         IRelacionalesContext
}

func NewEmptyRelacionalesContext() *RelacionalesContext {
	var p = new(RelacionalesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_relacionales
	return p
}

func (*RelacionalesContext) IsRelacionalesContext() {}

func NewRelacionalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelacionalesContext {
	var p = new(RelacionalesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_relacionales

	return p
}

func (s *RelacionalesContext) GetParser() antlr.Parser { return s.parser }

func (s *RelacionalesContext) GetOp() antlr.Token { return s.op }

func (s *RelacionalesContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelacionalesContext) GetOpIz() IRelacionalesContext { return s.opIz }

func (s *RelacionalesContext) Get_aritmeticas() IAritmeticasContext { return s._aritmeticas }

func (s *RelacionalesContext) GetOpDe() IRelacionalesContext { return s.opDe }

func (s *RelacionalesContext) SetOpIz(v IRelacionalesContext) { s.opIz = v }

func (s *RelacionalesContext) Set_aritmeticas(v IAritmeticasContext) { s._aritmeticas = v }

func (s *RelacionalesContext) SetOpDe(v IRelacionalesContext) { s.opDe = v }

func (s *RelacionalesContext) GetP() Interfaces.Expresion { return s.p }

func (s *RelacionalesContext) SetP(v Interfaces.Expresion) { s.p = v }

func (s *RelacionalesContext) Aritmeticas() IAritmeticasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAritmeticasContext)
}

func (s *RelacionalesContext) AllRelacionales() []IRelacionalesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelacionalesContext)(nil)).Elem())
	var tst = make([]IRelacionalesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelacionalesContext)
		}
	}

	return tst
}

func (s *RelacionalesContext) Relacionales(i int) IRelacionalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelacionalesContext)
}

func (s *RelacionalesContext) MENI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMENI, 0)
}

func (s *RelacionalesContext) MAYI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAYI, 0)
}

func (s *RelacionalesContext) MEN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEN, 0)
}

func (s *RelacionalesContext) MAY() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAY, 0)
}

func (s *RelacionalesContext) IGUALI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIGUALI, 0)
}

func (s *RelacionalesContext) DIFERENCIA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIFERENCIA, 0)
}

func (s *RelacionalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelacionalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelacionalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterRelacionales(s)
	}
}

func (s *RelacionalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitRelacionales(s)
	}
}

func (p *GramaticaParser) Relacionales() (localctx IRelacionalesContext) {
	return p.relacionales(0)
}

func (p *GramaticaParser) relacionales(_p int) (localctx IRelacionalesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelacionalesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelacionalesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, GramaticaParserRULE_relacionales, _p)
	var _la int

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
		p.SetState(163)

		var _x = p.aritmeticas(0)

		localctx.(*RelacionalesContext)._aritmeticas = _x
	}
	localctx.(*RelacionalesContext).p = localctx.(*RelacionalesContext).Get_aritmeticas().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelacionalesContext(p, _parentctx, _parentState)
			localctx.(*RelacionalesContext).opIz = _prevctx
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_relacionales)
			p.SetState(166)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(167)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RelacionalesContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserMENI)|(1<<GramaticaParserMAYI)|(1<<GramaticaParserMEN)|(1<<GramaticaParserMAY)|(1<<GramaticaParserIGUALI)|(1<<GramaticaParserDIFERENCIA))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RelacionalesContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(168)

				var _x = p.relacionales(3)

				localctx.(*RelacionalesContext).opDe = _x
			}
			localctx.(*RelacionalesContext).p = Expresion.NewRelacional((func() int {
				if localctx.(*RelacionalesContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*RelacionalesContext).GetOp().GetLine()
				}
			}()), localctx.(*RelacionalesContext).GetOp().GetColumn(), localctx.(*RelacionalesContext).GetOpIz().GetP(), (func() string {
				if localctx.(*RelacionalesContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*RelacionalesContext).GetOp().GetText()
				}
			}()), localctx.(*RelacionalesContext).GetOpDe().GetP(), false)

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IAritmeticasContext is an interface to support dynamic dispatch.
type IAritmeticasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IAritmeticasContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IAritmeticasContext

	// Get_primitivos returns the _primitivos rule contexts.
	Get_primitivos() IPrimitivosContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IAritmeticasContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IAritmeticasContext)

	// Set_primitivos sets the _primitivos rule contexts.
	Set_primitivos(IPrimitivosContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetP returns the p attribute.
	GetP() Interfaces.Expresion

	// SetP sets the p attribute.
	SetP(Interfaces.Expresion)

	// IsAritmeticasContext differentiates from other interfaces.
	IsAritmeticasContext()
}

type AritmeticasContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Interfaces.Expresion
	opIz        IAritmeticasContext
	op          antlr.Token
	opDe        IAritmeticasContext
	_primitivos IPrimitivosContext
	_expresion  IExpresionContext
}

func NewEmptyAritmeticasContext() *AritmeticasContext {
	var p = new(AritmeticasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_aritmeticas
	return p
}

func (*AritmeticasContext) IsAritmeticasContext() {}

func NewAritmeticasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AritmeticasContext {
	var p = new(AritmeticasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_aritmeticas

	return p
}

func (s *AritmeticasContext) GetParser() antlr.Parser { return s.parser }

func (s *AritmeticasContext) GetOp() antlr.Token { return s.op }

func (s *AritmeticasContext) SetOp(v antlr.Token) { s.op = v }

func (s *AritmeticasContext) GetOpIz() IAritmeticasContext { return s.opIz }

func (s *AritmeticasContext) GetOpDe() IAritmeticasContext { return s.opDe }

func (s *AritmeticasContext) Get_primitivos() IPrimitivosContext { return s._primitivos }

func (s *AritmeticasContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AritmeticasContext) SetOpIz(v IAritmeticasContext) { s.opIz = v }

func (s *AritmeticasContext) SetOpDe(v IAritmeticasContext) { s.opDe = v }

func (s *AritmeticasContext) Set_primitivos(v IPrimitivosContext) { s._primitivos = v }

func (s *AritmeticasContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AritmeticasContext) GetP() Interfaces.Expresion { return s.p }

func (s *AritmeticasContext) SetP(v Interfaces.Expresion) { s.p = v }

func (s *AritmeticasContext) MENOS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMENOS, 0)
}

func (s *AritmeticasContext) AllAritmeticas() []IAritmeticasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAritmeticasContext)(nil)).Elem())
	var tst = make([]IAritmeticasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAritmeticasContext)
		}
	}

	return tst
}

func (s *AritmeticasContext) Aritmeticas(i int) IAritmeticasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAritmeticasContext)
}

func (s *AritmeticasContext) Primitivos() IPrimitivosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivosContext)
}

func (s *AritmeticasContext) PARA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARA, 0)
}

func (s *AritmeticasContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AritmeticasContext) PARC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARC, 0)
}

func (s *AritmeticasContext) MOD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMOD, 0)
}

func (s *AritmeticasContext) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *AritmeticasContext) DIVIDIR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIVIDIR, 0)
}

func (s *AritmeticasContext) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *AritmeticasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AritmeticasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AritmeticasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAritmeticas(s)
	}
}

func (s *AritmeticasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAritmeticas(s)
	}
}

func (p *GramaticaParser) Aritmeticas() (localctx IAritmeticasContext) {
	return p.aritmeticas(0)
}

func (p *GramaticaParser) aritmeticas(_p int) (localctx IAritmeticasContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAritmeticasContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAritmeticasContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, GramaticaParserRULE_aritmeticas, _p)
	var _la int

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
	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserMENOS:
		{
			p.SetState(177)

			var _m = p.Match(GramaticaParserMENOS)

			localctx.(*AritmeticasContext).op = _m
		}
		{
			p.SetState(178)

			var _x = p.aritmeticas(6)

			localctx.(*AritmeticasContext).opDe = _x
		}
		localctx.(*AritmeticasContext).p = Expresion.NewAritmetica((func() int {
			if localctx.(*AritmeticasContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*AritmeticasContext).GetOp().GetLine()
			}
		}()), localctx.(*AritmeticasContext).GetOp().GetColumn(), nil, (func() string {
			if localctx.(*AritmeticasContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*AritmeticasContext).GetOp().GetText()
			}
		}()), localctx.(*AritmeticasContext).GetOpDe().GetP(), true)

	case GramaticaParserCADENA, GramaticaParserENTERO, GramaticaParserDECIMAL, GramaticaParserTRUE, GramaticaParserFALSE, GramaticaParserID:
		{
			p.SetState(181)

			var _x = p.Primitivos()

			localctx.(*AritmeticasContext)._primitivos = _x
		}
		localctx.(*AritmeticasContext).p = localctx.(*AritmeticasContext).Get_primitivos().GetP()

	case GramaticaParserPARA:
		{
			p.SetState(184)
			p.Match(GramaticaParserPARA)
		}
		{
			p.SetState(185)

			var _x = p.Expresion()

			localctx.(*AritmeticasContext)._expresion = _x
		}
		{
			p.SetState(186)
			p.Match(GramaticaParserPARC)
		}
		localctx.(*AritmeticasContext).p = localctx.(*AritmeticasContext).Get_expresion().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticasContext(p, _parentctx, _parentState)
				localctx.(*AritmeticasContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_aritmeticas)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(192)

					var _m = p.Match(GramaticaParserMOD)

					localctx.(*AritmeticasContext).op = _m
				}
				{
					p.SetState(193)

					var _x = p.aritmeticas(6)

					localctx.(*AritmeticasContext).opDe = _x
				}
				localctx.(*AritmeticasContext).p = Expresion.NewAritmetica((func() int {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetLine()
					}
				}()), localctx.(*AritmeticasContext).GetOp().GetColumn(), localctx.(*AritmeticasContext).GetOpIz().GetP(), (func() string {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetText()
					}
				}()), localctx.(*AritmeticasContext).GetOpDe().GetP(), false)

			case 2:
				localctx = NewAritmeticasContext(p, _parentctx, _parentState)
				localctx.(*AritmeticasContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_aritmeticas)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(197)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticasContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserPOR || _la == GramaticaParserDIVIDIR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticasContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(198)

					var _x = p.aritmeticas(5)

					localctx.(*AritmeticasContext).opDe = _x
				}
				localctx.(*AritmeticasContext).p = Expresion.NewAritmetica((func() int {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetLine()
					}
				}()), localctx.(*AritmeticasContext).GetOp().GetColumn(), localctx.(*AritmeticasContext).GetOpIz().GetP(), (func() string {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetText()
					}
				}()), localctx.(*AritmeticasContext).GetOpDe().GetP(), false)

			case 3:
				localctx = NewAritmeticasContext(p, _parentctx, _parentState)
				localctx.(*AritmeticasContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_aritmeticas)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(202)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticasContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserMAS || _la == GramaticaParserMENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticasContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(203)

					var _x = p.aritmeticas(4)

					localctx.(*AritmeticasContext).opDe = _x
				}
				localctx.(*AritmeticasContext).p = Expresion.NewAritmetica((func() int {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetLine()
					}
				}()), localctx.(*AritmeticasContext).GetOp().GetColumn(), localctx.(*AritmeticasContext).GetOpIz().GetP(), (func() string {
					if localctx.(*AritmeticasContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AritmeticasContext).GetOp().GetText()
					}
				}()), localctx.(*AritmeticasContext).GetOpDe().GetP(), false)

			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivosContext is an interface to support dynamic dispatch.
type IPrimitivosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetP returns the p attribute.
	GetP() Interfaces.Expresion

	// SetP sets the p attribute.
	SetP(Interfaces.Expresion)

	// IsPrimitivosContext differentiates from other interfaces.
	IsPrimitivosContext()
}

type PrimitivosContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        Interfaces.Expresion
	_ENTERO  antlr.Token
	_DECIMAL antlr.Token
	_TRUE    antlr.Token
	_FALSE   antlr.Token
	_CADENA  antlr.Token
	_ID      antlr.Token
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

func (s *PrimitivosContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *PrimitivosContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *PrimitivosContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *PrimitivosContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *PrimitivosContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrimitivosContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivosContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *PrimitivosContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *PrimitivosContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *PrimitivosContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *PrimitivosContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivosContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivosContext) GetP() Interfaces.Expresion { return s.p }

func (s *PrimitivosContext) SetP(v Interfaces.Expresion) { s.p = v }

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

func (s *PrimitivosContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
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
	p.EnterRule(localctx, 24, GramaticaParserRULE_primitivos)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)

			var _m = p.Match(GramaticaParserENTERO)

			localctx.(*PrimitivosContext)._ENTERO = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivosContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_ENTERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivosContext).p = Expresion.NewPrimitivo((func() int {
			if localctx.(*PrimitivosContext).Get_ENTERO() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_ENTERO().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_ENTERO().GetColumn(), num, Entornos.INTEGER)

	case GramaticaParserDECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)

			var _m = p.Match(GramaticaParserDECIMAL)

			localctx.(*PrimitivosContext)._DECIMAL = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivosContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivosContext).p = Expresion.NewPrimitivo((func() int {
			if localctx.(*PrimitivosContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_DECIMAL().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_DECIMAL().GetColumn(), num, Entornos.REAL)

	case GramaticaParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(215)

			var _m = p.Match(GramaticaParserTRUE)

			localctx.(*PrimitivosContext)._TRUE = _m
		}

		localctx.(*PrimitivosContext).p = Expresion.NewPrimitivo((func() int {
			if localctx.(*PrimitivosContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_TRUE().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_TRUE().GetColumn(), (func() string {
			if localctx.(*PrimitivosContext).Get_TRUE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_TRUE().GetText()
			}
		}()), Entornos.BOOLEAN)

	case GramaticaParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)

			var _m = p.Match(GramaticaParserFALSE)

			localctx.(*PrimitivosContext)._FALSE = _m
		}

		localctx.(*PrimitivosContext).p = Expresion.NewPrimitivo((func() int {
			if localctx.(*PrimitivosContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_FALSE().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_FALSE().GetColumn(), (func() string {
			if localctx.(*PrimitivosContext).Get_FALSE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_FALSE().GetText()
			}
		}()), Entornos.BOOLEAN)

	case GramaticaParserCADENA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(219)

			var _m = p.Match(GramaticaParserCADENA)

			localctx.(*PrimitivosContext)._CADENA = _m
		}

		str := (func() string {
			if localctx.(*PrimitivosContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivosContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_CADENA().GetText()
			}
		}()))-1]
		localctx.(*PrimitivosContext).p = Expresion.NewPrimitivo((func() int {
			if localctx.(*PrimitivosContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_CADENA().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_CADENA().GetColumn(), str, Entornos.STRING)

	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(221)

			var _m = p.Match(GramaticaParserID)

			localctx.(*PrimitivosContext)._ID = _m
		}
		localctx.(*PrimitivosContext).p = Expresion.NewAcceso((func() int {
			if localctx.(*PrimitivosContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitivosContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitivosContext).Get_ID().GetColumn(), (func() string {
			if localctx.(*PrimitivosContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivosContext).Get_ID().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *IdentificadoresContext = nil
		if localctx != nil {
			t = localctx.(*IdentificadoresContext)
		}
		return p.Identificadores_Sempred(t, predIndex)

	case 10:
		var t *RelacionalesContext = nil
		if localctx != nil {
			t = localctx.(*RelacionalesContext)
		}
		return p.Relacionales_Sempred(t, predIndex)

	case 11:
		var t *AritmeticasContext = nil
		if localctx != nil {
			t = localctx.(*AritmeticasContext)
		}
		return p.Aritmeticas_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Identificadores_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Relacionales_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Aritmeticas_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
