// Code generated from LittleDuck.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LittleDuck

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 190,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 79,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 85, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 95, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 5, 11, 102, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 120,
	10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 133, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 138, 10,
	17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	5, 19, 150, 10, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 5, 21, 160, 10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 5, 23, 170, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 5, 24, 179, 10, 24, 3, 25, 3, 25, 3, 25, 5, 25, 184, 10, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 2, 2, 28, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 2, 4, 3,
	2, 5, 6, 4, 2, 25, 25, 27, 28, 2, 182, 2, 54, 3, 2, 2, 2, 4, 62, 3, 2,
	2, 2, 6, 64, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 78, 3, 2, 2, 2, 12, 84,
	3, 2, 2, 2, 14, 86, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2,
	20, 101, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 108, 3, 2, 2, 2, 26, 119,
	3, 2, 2, 2, 28, 121, 3, 2, 2, 2, 30, 132, 3, 2, 2, 2, 32, 137, 3, 2, 2,
	2, 34, 139, 3, 2, 2, 2, 36, 149, 3, 2, 2, 2, 38, 151, 3, 2, 2, 2, 40, 159,
	3, 2, 2, 2, 42, 161, 3, 2, 2, 2, 44, 169, 3, 2, 2, 2, 46, 178, 3, 2, 2,
	2, 48, 183, 3, 2, 2, 2, 50, 185, 3, 2, 2, 2, 52, 187, 3, 2, 2, 2, 54, 55,
	7, 3, 2, 2, 55, 56, 7, 25, 2, 2, 56, 57, 7, 12, 2, 2, 57, 58, 5, 4, 3,
	2, 58, 59, 5, 14, 8, 2, 59, 3, 3, 2, 2, 2, 60, 63, 5, 6, 4, 2, 61, 63,
	5, 52, 27, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 5, 3, 2, 2, 2,
	64, 65, 7, 4, 2, 2, 65, 66, 5, 8, 5, 2, 66, 67, 5, 12, 7, 2, 67, 7, 3,
	2, 2, 2, 68, 69, 7, 25, 2, 2, 69, 70, 5, 10, 6, 2, 70, 71, 7, 11, 2, 2,
	71, 72, 5, 18, 10, 2, 72, 73, 7, 12, 2, 2, 73, 9, 3, 2, 2, 2, 74, 75, 7,
	10, 2, 2, 75, 76, 7, 25, 2, 2, 76, 79, 5, 10, 6, 2, 77, 79, 5, 52, 27,
	2, 78, 74, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 11, 3, 2, 2, 2, 80, 81,
	5, 8, 5, 2, 81, 82, 5, 12, 7, 2, 82, 85, 3, 2, 2, 2, 83, 85, 5, 52, 27,
	2, 84, 80, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 13, 3, 2, 2, 2, 86, 87,
	7, 13, 2, 2, 87, 88, 5, 16, 9, 2, 88, 89, 7, 14, 2, 2, 89, 15, 3, 2, 2,
	2, 90, 91, 5, 20, 11, 2, 91, 92, 5, 16, 9, 2, 92, 95, 3, 2, 2, 2, 93, 95,
	5, 52, 27, 2, 94, 90, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 17, 3, 2, 2,
	2, 96, 97, 9, 2, 2, 2, 97, 19, 3, 2, 2, 2, 98, 102, 5, 22, 12, 2, 99, 102,
	5, 24, 13, 2, 100, 102, 5, 28, 15, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3,
	2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 21, 3, 2, 2, 2, 103, 104, 7, 25, 2,
	2, 104, 105, 7, 21, 2, 2, 105, 106, 5, 34, 18, 2, 106, 107, 7, 12, 2, 2,
	107, 23, 3, 2, 2, 2, 108, 109, 7, 8, 2, 2, 109, 110, 7, 15, 2, 2, 110,
	111, 5, 34, 18, 2, 111, 112, 7, 16, 2, 2, 112, 113, 5, 14, 8, 2, 113, 114,
	5, 26, 14, 2, 114, 115, 7, 12, 2, 2, 115, 25, 3, 2, 2, 2, 116, 117, 7,
	9, 2, 2, 117, 120, 5, 14, 8, 2, 118, 120, 5, 52, 27, 2, 119, 116, 3, 2,
	2, 2, 119, 118, 3, 2, 2, 2, 120, 27, 3, 2, 2, 2, 121, 122, 7, 7, 2, 2,
	122, 123, 7, 15, 2, 2, 123, 124, 5, 30, 16, 2, 124, 125, 7, 16, 2, 2, 125,
	126, 7, 12, 2, 2, 126, 29, 3, 2, 2, 2, 127, 128, 7, 26, 2, 2, 128, 133,
	5, 32, 17, 2, 129, 130, 5, 34, 18, 2, 130, 131, 5, 32, 17, 2, 131, 133,
	3, 2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 133, 31, 3, 2,
	2, 2, 134, 135, 7, 10, 2, 2, 135, 138, 5, 30, 16, 2, 136, 138, 5, 52, 27,
	2, 137, 134, 3, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 33, 3, 2, 2, 2, 139,
	140, 5, 38, 20, 2, 140, 141, 5, 36, 19, 2, 141, 35, 3, 2, 2, 2, 142, 143,
	7, 23, 2, 2, 143, 150, 5, 38, 20, 2, 144, 145, 7, 22, 2, 2, 145, 150, 5,
	38, 20, 2, 146, 147, 7, 24, 2, 2, 147, 150, 5, 38, 20, 2, 148, 150, 5,
	52, 27, 2, 149, 142, 3, 2, 2, 2, 149, 144, 3, 2, 2, 2, 149, 146, 3, 2,
	2, 2, 149, 148, 3, 2, 2, 2, 150, 37, 3, 2, 2, 2, 151, 152, 5, 42, 22, 2,
	152, 153, 5, 40, 21, 2, 153, 39, 3, 2, 2, 2, 154, 155, 7, 18, 2, 2, 155,
	160, 5, 38, 20, 2, 156, 157, 7, 19, 2, 2, 157, 160, 5, 38, 20, 2, 158,
	160, 5, 52, 27, 2, 159, 154, 3, 2, 2, 2, 159, 156, 3, 2, 2, 2, 159, 158,
	3, 2, 2, 2, 160, 41, 3, 2, 2, 2, 161, 162, 5, 46, 24, 2, 162, 163, 5, 44,
	23, 2, 163, 43, 3, 2, 2, 2, 164, 165, 7, 17, 2, 2, 165, 170, 5, 42, 22,
	2, 166, 167, 7, 20, 2, 2, 167, 170, 5, 42, 22, 2, 168, 170, 5, 52, 27,
	2, 169, 164, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170,
	45, 3, 2, 2, 2, 171, 172, 7, 15, 2, 2, 172, 173, 5, 34, 18, 2, 173, 174,
	7, 16, 2, 2, 174, 179, 3, 2, 2, 2, 175, 176, 5, 48, 25, 2, 176, 177, 5,
	50, 26, 2, 177, 179, 3, 2, 2, 2, 178, 171, 3, 2, 2, 2, 178, 175, 3, 2,
	2, 2, 179, 47, 3, 2, 2, 2, 180, 184, 7, 18, 2, 2, 181, 184, 7, 19, 2, 2,
	182, 184, 5, 52, 27, 2, 183, 180, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183,
	182, 3, 2, 2, 2, 184, 49, 3, 2, 2, 2, 185, 186, 9, 3, 2, 2, 186, 51, 3,
	2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 53, 3, 2, 2, 2, 15, 62, 78, 84, 94,
	101, 119, 132, 137, 149, 159, 169, 178, 183,
}
var literalNames = []string{
	"", "'programa'", "'var'", "'int'", "'float'", "'print'", "'if'", "'else'",
	"','", "':'", "';'", "'{'", "'}'", "'('", "')'", "'*'", "'+'", "'-'", "'/'",
	"'='", "'<'", "'>'", "'<>'",
}
var symbolicNames = []string{
	"", "PROGRAM", "VAR", "INT", "FLOAT", "PRINT", "IF", "ELSE", "COMMA", "COLON",
	"SEMICOLON", "LEFT_CURLY", "RIGHT_CURLY", "LEFT_ROUND", "RIGHT_ROUND",
	"STAR", "PLUS", "MINUS", "DIV", "EQUAL", "LESS_THAN", "GREATER_THAN", "NOT_EQUAL",
	"LET_ID", "LET_STRING", "LET_INT", "LET_FLOAT", "WHITESPACE",
}

var ruleNames = []string{
	"programa", "programa_vars", "vars", "vars_id", "vars_id_next", "vars_more",
	"bloque", "bloque_more", "tipo", "estatuto", "asignacion", "condicion",
	"condicion_else", "escritura", "escritura_params", "escritura_params_next",
	"expresion", "expresion_next", "exp", "exp_next", "termino", "termino_next",
	"factor", "factor_sign", "cte", "empty",
}

type LittleDuckParser struct {
	*antlr.BaseParser
}

// NewLittleDuckParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LittleDuckParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLittleDuckParser(input antlr.TokenStream) *LittleDuckParser {
	this := new(LittleDuckParser)
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
	this.GrammarFileName = "LittleDuck.g4"

	return this
}

// LittleDuckParser tokens.
const (
	LittleDuckParserEOF          = antlr.TokenEOF
	LittleDuckParserPROGRAM      = 1
	LittleDuckParserVAR          = 2
	LittleDuckParserINT          = 3
	LittleDuckParserFLOAT        = 4
	LittleDuckParserPRINT        = 5
	LittleDuckParserIF           = 6
	LittleDuckParserELSE         = 7
	LittleDuckParserCOMMA        = 8
	LittleDuckParserCOLON        = 9
	LittleDuckParserSEMICOLON    = 10
	LittleDuckParserLEFT_CURLY   = 11
	LittleDuckParserRIGHT_CURLY  = 12
	LittleDuckParserLEFT_ROUND   = 13
	LittleDuckParserRIGHT_ROUND  = 14
	LittleDuckParserSTAR         = 15
	LittleDuckParserPLUS         = 16
	LittleDuckParserMINUS        = 17
	LittleDuckParserDIV          = 18
	LittleDuckParserEQUAL        = 19
	LittleDuckParserLESS_THAN    = 20
	LittleDuckParserGREATER_THAN = 21
	LittleDuckParserNOT_EQUAL    = 22
	LittleDuckParserLET_ID       = 23
	LittleDuckParserLET_STRING   = 24
	LittleDuckParserLET_INT      = 25
	LittleDuckParserLET_FLOAT    = 26
	LittleDuckParserWHITESPACE   = 27
)

// LittleDuckParser rules.
const (
	LittleDuckParserRULE_programa              = 0
	LittleDuckParserRULE_programa_vars         = 1
	LittleDuckParserRULE_vars                  = 2
	LittleDuckParserRULE_vars_id               = 3
	LittleDuckParserRULE_vars_id_next          = 4
	LittleDuckParserRULE_vars_more             = 5
	LittleDuckParserRULE_bloque                = 6
	LittleDuckParserRULE_bloque_more           = 7
	LittleDuckParserRULE_tipo                  = 8
	LittleDuckParserRULE_estatuto              = 9
	LittleDuckParserRULE_asignacion            = 10
	LittleDuckParserRULE_condicion             = 11
	LittleDuckParserRULE_condicion_else        = 12
	LittleDuckParserRULE_escritura             = 13
	LittleDuckParserRULE_escritura_params      = 14
	LittleDuckParserRULE_escritura_params_next = 15
	LittleDuckParserRULE_expresion             = 16
	LittleDuckParserRULE_expresion_next        = 17
	LittleDuckParserRULE_exp                   = 18
	LittleDuckParserRULE_exp_next              = 19
	LittleDuckParserRULE_termino               = 20
	LittleDuckParserRULE_termino_next          = 21
	LittleDuckParserRULE_factor                = 22
	LittleDuckParserRULE_factor_sign           = 23
	LittleDuckParserRULE_cte                   = 24
	LittleDuckParserRULE_empty                 = 25
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_programa
	return p
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserPROGRAM, 0)
}

func (s *ProgramaContext) LET_ID() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_ID, 0)
}

func (s *ProgramaContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSEMICOLON, 0)
}

func (s *ProgramaContext) Programa_vars() IPrograma_varsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrograma_varsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrograma_varsContext)
}

func (s *ProgramaContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *LittleDuckParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LittleDuckParserRULE_programa)

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
		p.SetState(52)
		p.Match(LittleDuckParserPROGRAM)
	}
	{
		p.SetState(53)
		p.Match(LittleDuckParserLET_ID)
	}
	{
		p.SetState(54)
		p.Match(LittleDuckParserSEMICOLON)
	}
	{
		p.SetState(55)
		p.Programa_vars()
	}
	{
		p.SetState(56)
		p.Bloque()
	}

	return localctx
}

// IPrograma_varsContext is an interface to support dynamic dispatch.
type IPrograma_varsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrograma_varsContext differentiates from other interfaces.
	IsPrograma_varsContext()
}

type Programa_varsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrograma_varsContext() *Programa_varsContext {
	var p = new(Programa_varsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_programa_vars
	return p
}

func (*Programa_varsContext) IsPrograma_varsContext() {}

func NewPrograma_varsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Programa_varsContext {
	var p = new(Programa_varsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_programa_vars

	return p
}

func (s *Programa_varsContext) GetParser() antlr.Parser { return s.parser }

func (s *Programa_varsContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *Programa_varsContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Programa_varsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Programa_varsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Programa_varsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterPrograma_vars(s)
	}
}

func (s *Programa_varsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitPrograma_vars(s)
	}
}

func (p *LittleDuckParser) Programa_vars() (localctx IPrograma_varsContext) {
	localctx = NewPrograma_varsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LittleDuckParserRULE_programa_vars)

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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Vars()
		}

	case LittleDuckParserLEFT_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) VAR() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserVAR, 0)
}

func (s *VarsContext) Vars_id() IVars_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_idContext)
}

func (s *VarsContext) Vars_more() IVars_moreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_moreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_moreContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *LittleDuckParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LittleDuckParserRULE_vars)

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
		p.SetState(62)
		p.Match(LittleDuckParserVAR)
	}
	{
		p.SetState(63)
		p.Vars_id()
	}
	{
		p.SetState(64)
		p.Vars_more()
	}

	return localctx
}

// IVars_idContext is an interface to support dynamic dispatch.
type IVars_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_idContext differentiates from other interfaces.
	IsVars_idContext()
}

type Vars_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_idContext() *Vars_idContext {
	var p = new(Vars_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_vars_id
	return p
}

func (*Vars_idContext) IsVars_idContext() {}

func NewVars_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_idContext {
	var p = new(Vars_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_vars_id

	return p
}

func (s *Vars_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_idContext) LET_ID() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_ID, 0)
}

func (s *Vars_idContext) Vars_id_next() IVars_id_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_id_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_id_nextContext)
}

func (s *Vars_idContext) COLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserCOLON, 0)
}

func (s *Vars_idContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Vars_idContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSEMICOLON, 0)
}

func (s *Vars_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterVars_id(s)
	}
}

func (s *Vars_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitVars_id(s)
	}
}

func (p *LittleDuckParser) Vars_id() (localctx IVars_idContext) {
	localctx = NewVars_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LittleDuckParserRULE_vars_id)

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
		p.SetState(66)
		p.Match(LittleDuckParserLET_ID)
	}
	{
		p.SetState(67)
		p.Vars_id_next()
	}
	{
		p.SetState(68)
		p.Match(LittleDuckParserCOLON)
	}
	{
		p.SetState(69)
		p.Tipo()
	}
	{
		p.SetState(70)
		p.Match(LittleDuckParserSEMICOLON)
	}

	return localctx
}

// IVars_id_nextContext is an interface to support dynamic dispatch.
type IVars_id_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_id_nextContext differentiates from other interfaces.
	IsVars_id_nextContext()
}

type Vars_id_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_id_nextContext() *Vars_id_nextContext {
	var p = new(Vars_id_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_vars_id_next
	return p
}

func (*Vars_id_nextContext) IsVars_id_nextContext() {}

func NewVars_id_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_id_nextContext {
	var p = new(Vars_id_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_vars_id_next

	return p
}

func (s *Vars_id_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_id_nextContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserCOMMA, 0)
}

func (s *Vars_id_nextContext) LET_ID() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_ID, 0)
}

func (s *Vars_id_nextContext) Vars_id_next() IVars_id_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_id_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_id_nextContext)
}

func (s *Vars_id_nextContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Vars_id_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_id_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_id_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterVars_id_next(s)
	}
}

func (s *Vars_id_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitVars_id_next(s)
	}
}

func (p *LittleDuckParser) Vars_id_next() (localctx IVars_id_nextContext) {
	localctx = NewVars_id_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LittleDuckParserRULE_vars_id_next)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(LittleDuckParserCOMMA)
		}
		{
			p.SetState(73)
			p.Match(LittleDuckParserLET_ID)
		}
		{
			p.SetState(74)
			p.Vars_id_next()
		}

	case LittleDuckParserCOLON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVars_moreContext is an interface to support dynamic dispatch.
type IVars_moreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_moreContext differentiates from other interfaces.
	IsVars_moreContext()
}

type Vars_moreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_moreContext() *Vars_moreContext {
	var p = new(Vars_moreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_vars_more
	return p
}

func (*Vars_moreContext) IsVars_moreContext() {}

func NewVars_moreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_moreContext {
	var p = new(Vars_moreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_vars_more

	return p
}

func (s *Vars_moreContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_moreContext) Vars_id() IVars_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_idContext)
}

func (s *Vars_moreContext) Vars_more() IVars_moreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_moreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_moreContext)
}

func (s *Vars_moreContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Vars_moreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_moreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_moreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterVars_more(s)
	}
}

func (s *Vars_moreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitVars_more(s)
	}
}

func (p *LittleDuckParser) Vars_more() (localctx IVars_moreContext) {
	localctx = NewVars_moreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LittleDuckParserRULE_vars_more)

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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserLET_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Vars_id()
		}
		{
			p.SetState(79)
			p.Vars_more()
		}

	case LittleDuckParserLEFT_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) LEFT_CURLY() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLEFT_CURLY, 0)
}

func (s *BloqueContext) Bloque_more() IBloque_moreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_moreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_moreContext)
}

func (s *BloqueContext) RIGHT_CURLY() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserRIGHT_CURLY, 0)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *LittleDuckParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LittleDuckParserRULE_bloque)

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
		p.SetState(84)
		p.Match(LittleDuckParserLEFT_CURLY)
	}
	{
		p.SetState(85)
		p.Bloque_more()
	}
	{
		p.SetState(86)
		p.Match(LittleDuckParserRIGHT_CURLY)
	}

	return localctx
}

// IBloque_moreContext is an interface to support dynamic dispatch.
type IBloque_moreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloque_moreContext differentiates from other interfaces.
	IsBloque_moreContext()
}

type Bloque_moreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloque_moreContext() *Bloque_moreContext {
	var p = new(Bloque_moreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_bloque_more
	return p
}

func (*Bloque_moreContext) IsBloque_moreContext() {}

func NewBloque_moreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloque_moreContext {
	var p = new(Bloque_moreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_bloque_more

	return p
}

func (s *Bloque_moreContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloque_moreContext) Estatuto() IEstatutoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEstatutoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEstatutoContext)
}

func (s *Bloque_moreContext) Bloque_more() IBloque_moreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_moreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_moreContext)
}

func (s *Bloque_moreContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Bloque_moreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloque_moreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloque_moreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterBloque_more(s)
	}
}

func (s *Bloque_moreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitBloque_more(s)
	}
}

func (p *LittleDuckParser) Bloque_more() (localctx IBloque_moreContext) {
	localctx = NewBloque_moreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LittleDuckParserRULE_bloque_more)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserPRINT, LittleDuckParserIF, LittleDuckParserLET_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Estatuto()
		}
		{
			p.SetState(89)
			p.Bloque_more()
		}

	case LittleDuckParserRIGHT_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = LittleDuckParserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserINT, 0)
}

func (s *TipoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserFLOAT, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *LittleDuckParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LittleDuckParserRULE_tipo)
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
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LittleDuckParserINT || _la == LittleDuckParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEstatutoContext is an interface to support dynamic dispatch.
type IEstatutoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEstatutoContext differentiates from other interfaces.
	IsEstatutoContext()
}

type EstatutoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstatutoContext() *EstatutoContext {
	var p = new(EstatutoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_estatuto
	return p
}

func (*EstatutoContext) IsEstatutoContext() {}

func NewEstatutoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstatutoContext {
	var p = new(EstatutoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_estatuto

	return p
}

func (s *EstatutoContext) GetParser() antlr.Parser { return s.parser }

func (s *EstatutoContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *EstatutoContext) Condicion() ICondicionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondicionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *EstatutoContext) Escritura() IEscrituraContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscrituraContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscrituraContext)
}

func (s *EstatutoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstatutoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstatutoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterEstatuto(s)
	}
}

func (s *EstatutoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitEstatuto(s)
	}
}

func (p *LittleDuckParser) Estatuto() (localctx IEstatutoContext) {
	localctx = NewEstatutoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LittleDuckParserRULE_estatuto)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserLET_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Asignacion()
		}

	case LittleDuckParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Condicion()
		}

	case LittleDuckParserPRINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Escritura()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) LET_ID() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_ID, 0)
}

func (s *AsignacionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserEQUAL, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSEMICOLON, 0)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *LittleDuckParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LittleDuckParserRULE_asignacion)

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
		p.SetState(101)
		p.Match(LittleDuckParserLET_ID)
	}
	{
		p.SetState(102)
		p.Match(LittleDuckParserEQUAL)
	}
	{
		p.SetState(103)
		p.Expresion()
	}
	{
		p.SetState(104)
		p.Match(LittleDuckParserSEMICOLON)
	}

	return localctx
}

// ICondicionContext is an interface to support dynamic dispatch.
type ICondicionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondicionContext differentiates from other interfaces.
	IsCondicionContext()
}

type CondicionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondicionContext() *CondicionContext {
	var p = new(CondicionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_condicion
	return p
}

func (*CondicionContext) IsCondicionContext() {}

func NewCondicionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondicionContext {
	var p = new(CondicionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_condicion

	return p
}

func (s *CondicionContext) GetParser() antlr.Parser { return s.parser }

func (s *CondicionContext) IF() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserIF, 0)
}

func (s *CondicionContext) LEFT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLEFT_ROUND, 0)
}

func (s *CondicionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CondicionContext) RIGHT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserRIGHT_ROUND, 0)
}

func (s *CondicionContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *CondicionContext) Condicion_else() ICondicion_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondicion_elseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondicion_elseContext)
}

func (s *CondicionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSEMICOLON, 0)
}

func (s *CondicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondicionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterCondicion(s)
	}
}

func (s *CondicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitCondicion(s)
	}
}

func (p *LittleDuckParser) Condicion() (localctx ICondicionContext) {
	localctx = NewCondicionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LittleDuckParserRULE_condicion)

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
		p.SetState(106)
		p.Match(LittleDuckParserIF)
	}
	{
		p.SetState(107)
		p.Match(LittleDuckParserLEFT_ROUND)
	}
	{
		p.SetState(108)
		p.Expresion()
	}
	{
		p.SetState(109)
		p.Match(LittleDuckParserRIGHT_ROUND)
	}
	{
		p.SetState(110)
		p.Bloque()
	}
	{
		p.SetState(111)
		p.Condicion_else()
	}
	{
		p.SetState(112)
		p.Match(LittleDuckParserSEMICOLON)
	}

	return localctx
}

// ICondicion_elseContext is an interface to support dynamic dispatch.
type ICondicion_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondicion_elseContext differentiates from other interfaces.
	IsCondicion_elseContext()
}

type Condicion_elseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondicion_elseContext() *Condicion_elseContext {
	var p = new(Condicion_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_condicion_else
	return p
}

func (*Condicion_elseContext) IsCondicion_elseContext() {}

func NewCondicion_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condicion_elseContext {
	var p = new(Condicion_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_condicion_else

	return p
}

func (s *Condicion_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Condicion_elseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserELSE, 0)
}

func (s *Condicion_elseContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Condicion_elseContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Condicion_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condicion_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condicion_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterCondicion_else(s)
	}
}

func (s *Condicion_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitCondicion_else(s)
	}
}

func (p *LittleDuckParser) Condicion_else() (localctx ICondicion_elseContext) {
	localctx = NewCondicion_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LittleDuckParserRULE_condicion_else)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserELSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(LittleDuckParserELSE)
		}
		{
			p.SetState(115)
			p.Bloque()
		}

	case LittleDuckParserSEMICOLON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEscrituraContext is an interface to support dynamic dispatch.
type IEscrituraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscrituraContext differentiates from other interfaces.
	IsEscrituraContext()
}

type EscrituraContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscrituraContext() *EscrituraContext {
	var p = new(EscrituraContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_escritura
	return p
}

func (*EscrituraContext) IsEscrituraContext() {}

func NewEscrituraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EscrituraContext {
	var p = new(EscrituraContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_escritura

	return p
}

func (s *EscrituraContext) GetParser() antlr.Parser { return s.parser }

func (s *EscrituraContext) PRINT() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserPRINT, 0)
}

func (s *EscrituraContext) LEFT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLEFT_ROUND, 0)
}

func (s *EscrituraContext) Escritura_params() IEscritura_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscritura_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscritura_paramsContext)
}

func (s *EscrituraContext) RIGHT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserRIGHT_ROUND, 0)
}

func (s *EscrituraContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSEMICOLON, 0)
}

func (s *EscrituraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscrituraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EscrituraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterEscritura(s)
	}
}

func (s *EscrituraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitEscritura(s)
	}
}

func (p *LittleDuckParser) Escritura() (localctx IEscrituraContext) {
	localctx = NewEscrituraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LittleDuckParserRULE_escritura)

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
		p.SetState(119)
		p.Match(LittleDuckParserPRINT)
	}
	{
		p.SetState(120)
		p.Match(LittleDuckParserLEFT_ROUND)
	}
	{
		p.SetState(121)
		p.Escritura_params()
	}
	{
		p.SetState(122)
		p.Match(LittleDuckParserRIGHT_ROUND)
	}
	{
		p.SetState(123)
		p.Match(LittleDuckParserSEMICOLON)
	}

	return localctx
}

// IEscritura_paramsContext is an interface to support dynamic dispatch.
type IEscritura_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscritura_paramsContext differentiates from other interfaces.
	IsEscritura_paramsContext()
}

type Escritura_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscritura_paramsContext() *Escritura_paramsContext {
	var p = new(Escritura_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_escritura_params
	return p
}

func (*Escritura_paramsContext) IsEscritura_paramsContext() {}

func NewEscritura_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Escritura_paramsContext {
	var p = new(Escritura_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_escritura_params

	return p
}

func (s *Escritura_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Escritura_paramsContext) LET_STRING() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_STRING, 0)
}

func (s *Escritura_paramsContext) Escritura_params_next() IEscritura_params_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscritura_params_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscritura_params_nextContext)
}

func (s *Escritura_paramsContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Escritura_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Escritura_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Escritura_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterEscritura_params(s)
	}
}

func (s *Escritura_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitEscritura_params(s)
	}
}

func (p *LittleDuckParser) Escritura_params() (localctx IEscritura_paramsContext) {
	localctx = NewEscritura_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LittleDuckParserRULE_escritura_params)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserLET_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(LittleDuckParserLET_STRING)
		}
		{
			p.SetState(126)
			p.Escritura_params_next()
		}

	case LittleDuckParserLEFT_ROUND, LittleDuckParserPLUS, LittleDuckParserMINUS, LittleDuckParserLET_ID, LittleDuckParserLET_INT, LittleDuckParserLET_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Expresion()
		}
		{
			p.SetState(128)
			p.Escritura_params_next()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEscritura_params_nextContext is an interface to support dynamic dispatch.
type IEscritura_params_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscritura_params_nextContext differentiates from other interfaces.
	IsEscritura_params_nextContext()
}

type Escritura_params_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscritura_params_nextContext() *Escritura_params_nextContext {
	var p = new(Escritura_params_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_escritura_params_next
	return p
}

func (*Escritura_params_nextContext) IsEscritura_params_nextContext() {}

func NewEscritura_params_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Escritura_params_nextContext {
	var p = new(Escritura_params_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_escritura_params_next

	return p
}

func (s *Escritura_params_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Escritura_params_nextContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserCOMMA, 0)
}

func (s *Escritura_params_nextContext) Escritura_params() IEscritura_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscritura_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscritura_paramsContext)
}

func (s *Escritura_params_nextContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Escritura_params_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Escritura_params_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Escritura_params_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterEscritura_params_next(s)
	}
}

func (s *Escritura_params_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitEscritura_params_next(s)
	}
}

func (p *LittleDuckParser) Escritura_params_next() (localctx IEscritura_params_nextContext) {
	localctx = NewEscritura_params_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LittleDuckParserRULE_escritura_params_next)

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
	case LittleDuckParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(LittleDuckParserCOMMA)
		}
		{
			p.SetState(133)
			p.Escritura_params()
		}

	case LittleDuckParserRIGHT_ROUND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Empty()
		}

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
	p.RuleIndex = LittleDuckParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpresionContext) Expresion_next() IExpresion_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresion_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresion_nextContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *LittleDuckParser) Expresion() (localctx IExpresionContext) {
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LittleDuckParserRULE_expresion)

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
		p.SetState(137)
		p.Exp()
	}
	{
		p.SetState(138)
		p.Expresion_next()
	}

	return localctx
}

// IExpresion_nextContext is an interface to support dynamic dispatch.
type IExpresion_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresion_nextContext differentiates from other interfaces.
	IsExpresion_nextContext()
}

type Expresion_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresion_nextContext() *Expresion_nextContext {
	var p = new(Expresion_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_expresion_next
	return p
}

func (*Expresion_nextContext) IsExpresion_nextContext() {}

func NewExpresion_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expresion_nextContext {
	var p = new(Expresion_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_expresion_next

	return p
}

func (s *Expresion_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Expresion_nextContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserGREATER_THAN, 0)
}

func (s *Expresion_nextContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Expresion_nextContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLESS_THAN, 0)
}

func (s *Expresion_nextContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserNOT_EQUAL, 0)
}

func (s *Expresion_nextContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Expresion_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expresion_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterExpresion_next(s)
	}
}

func (s *Expresion_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitExpresion_next(s)
	}
}

func (p *LittleDuckParser) Expresion_next() (localctx IExpresion_nextContext) {
	localctx = NewExpresion_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LittleDuckParserRULE_expresion_next)

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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserGREATER_THAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(LittleDuckParserGREATER_THAN)
		}
		{
			p.SetState(141)
			p.Exp()
		}

	case LittleDuckParserLESS_THAN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(LittleDuckParserLESS_THAN)
		}
		{
			p.SetState(143)
			p.Exp()
		}

	case LittleDuckParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(LittleDuckParserNOT_EQUAL)
		}
		{
			p.SetState(145)
			p.Exp()
		}

	case LittleDuckParserCOMMA, LittleDuckParserSEMICOLON, LittleDuckParserRIGHT_ROUND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Termino() ITerminoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminoContext)
}

func (s *ExpContext) Exp_next() IExp_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp_nextContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *LittleDuckParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LittleDuckParserRULE_exp)

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
		p.SetState(149)
		p.Termino()
	}
	{
		p.SetState(150)
		p.Exp_next()
	}

	return localctx
}

// IExp_nextContext is an interface to support dynamic dispatch.
type IExp_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp_nextContext differentiates from other interfaces.
	IsExp_nextContext()
}

type Exp_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_nextContext() *Exp_nextContext {
	var p = new(Exp_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_exp_next
	return p
}

func (*Exp_nextContext) IsExp_nextContext() {}

func NewExp_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_nextContext {
	var p = new(Exp_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_exp_next

	return p
}

func (s *Exp_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_nextContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserPLUS, 0)
}

func (s *Exp_nextContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Exp_nextContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserMINUS, 0)
}

func (s *Exp_nextContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Exp_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterExp_next(s)
	}
}

func (s *Exp_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitExp_next(s)
	}
}

func (p *LittleDuckParser) Exp_next() (localctx IExp_nextContext) {
	localctx = NewExp_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LittleDuckParserRULE_exp_next)

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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(LittleDuckParserPLUS)
		}
		{
			p.SetState(153)
			p.Exp()
		}

	case LittleDuckParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(LittleDuckParserMINUS)
		}
		{
			p.SetState(155)
			p.Exp()
		}

	case LittleDuckParserCOMMA, LittleDuckParserSEMICOLON, LittleDuckParserRIGHT_ROUND, LittleDuckParserLESS_THAN, LittleDuckParserGREATER_THAN, LittleDuckParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITerminoContext is an interface to support dynamic dispatch.
type ITerminoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerminoContext differentiates from other interfaces.
	IsTerminoContext()
}

type TerminoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminoContext() *TerminoContext {
	var p = new(TerminoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_termino
	return p
}

func (*TerminoContext) IsTerminoContext() {}

func NewTerminoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminoContext {
	var p = new(TerminoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_termino

	return p
}

func (s *TerminoContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminoContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TerminoContext) Termino_next() ITermino_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermino_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermino_nextContext)
}

func (s *TerminoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterTermino(s)
	}
}

func (s *TerminoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitTermino(s)
	}
}

func (p *LittleDuckParser) Termino() (localctx ITerminoContext) {
	localctx = NewTerminoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LittleDuckParserRULE_termino)

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
		p.SetState(159)
		p.Factor()
	}
	{
		p.SetState(160)
		p.Termino_next()
	}

	return localctx
}

// ITermino_nextContext is an interface to support dynamic dispatch.
type ITermino_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermino_nextContext differentiates from other interfaces.
	IsTermino_nextContext()
}

type Termino_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermino_nextContext() *Termino_nextContext {
	var p = new(Termino_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_termino_next
	return p
}

func (*Termino_nextContext) IsTermino_nextContext() {}

func NewTermino_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Termino_nextContext {
	var p = new(Termino_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_termino_next

	return p
}

func (s *Termino_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Termino_nextContext) STAR() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserSTAR, 0)
}

func (s *Termino_nextContext) Termino() ITerminoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminoContext)
}

func (s *Termino_nextContext) DIV() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserDIV, 0)
}

func (s *Termino_nextContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Termino_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Termino_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Termino_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterTermino_next(s)
	}
}

func (s *Termino_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitTermino_next(s)
	}
}

func (p *LittleDuckParser) Termino_next() (localctx ITermino_nextContext) {
	localctx = NewTermino_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LittleDuckParserRULE_termino_next)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(LittleDuckParserSTAR)
		}
		{
			p.SetState(163)
			p.Termino()
		}

	case LittleDuckParserDIV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(LittleDuckParserDIV)
		}
		{
			p.SetState(165)
			p.Termino()
		}

	case LittleDuckParserCOMMA, LittleDuckParserSEMICOLON, LittleDuckParserRIGHT_ROUND, LittleDuckParserPLUS, LittleDuckParserMINUS, LittleDuckParserLESS_THAN, LittleDuckParserGREATER_THAN, LittleDuckParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) LEFT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLEFT_ROUND, 0)
}

func (s *FactorContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FactorContext) RIGHT_ROUND() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserRIGHT_ROUND, 0)
}

func (s *FactorContext) Factor_sign() IFactor_signContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactor_signContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactor_signContext)
}

func (s *FactorContext) Cte() ICteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICteContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *LittleDuckParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LittleDuckParserRULE_factor)

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

	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserLEFT_ROUND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(LittleDuckParserLEFT_ROUND)
		}
		{
			p.SetState(170)
			p.Expresion()
		}
		{
			p.SetState(171)
			p.Match(LittleDuckParserRIGHT_ROUND)
		}

	case LittleDuckParserPLUS, LittleDuckParserMINUS, LittleDuckParserLET_ID, LittleDuckParserLET_INT, LittleDuckParserLET_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Factor_sign()
		}
		{
			p.SetState(174)
			p.Cte()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFactor_signContext is an interface to support dynamic dispatch.
type IFactor_signContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactor_signContext differentiates from other interfaces.
	IsFactor_signContext()
}

type Factor_signContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactor_signContext() *Factor_signContext {
	var p = new(Factor_signContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_factor_sign
	return p
}

func (*Factor_signContext) IsFactor_signContext() {}

func NewFactor_signContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Factor_signContext {
	var p = new(Factor_signContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_factor_sign

	return p
}

func (s *Factor_signContext) GetParser() antlr.Parser { return s.parser }

func (s *Factor_signContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserPLUS, 0)
}

func (s *Factor_signContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserMINUS, 0)
}

func (s *Factor_signContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *Factor_signContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Factor_signContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Factor_signContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterFactor_sign(s)
	}
}

func (s *Factor_signContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitFactor_sign(s)
	}
}

func (p *LittleDuckParser) Factor_sign() (localctx IFactor_signContext) {
	localctx = NewFactor_signContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LittleDuckParserRULE_factor_sign)

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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LittleDuckParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Match(LittleDuckParserPLUS)
		}

	case LittleDuckParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(LittleDuckParserMINUS)
		}

	case LittleDuckParserLET_ID, LittleDuckParserLET_INT, LittleDuckParserLET_FLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.Empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICteContext is an interface to support dynamic dispatch.
type ICteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCteContext differentiates from other interfaces.
	IsCteContext()
}

type CteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCteContext() *CteContext {
	var p = new(CteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_cte
	return p
}

func (*CteContext) IsCteContext() {}

func NewCteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CteContext {
	var p = new(CteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_cte

	return p
}

func (s *CteContext) GetParser() antlr.Parser { return s.parser }

func (s *CteContext) LET_ID() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_ID, 0)
}

func (s *CteContext) LET_INT() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_INT, 0)
}

func (s *CteContext) LET_FLOAT() antlr.TerminalNode {
	return s.GetToken(LittleDuckParserLET_FLOAT, 0)
}

func (s *CteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterCte(s)
	}
}

func (s *CteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitCte(s)
	}
}

func (p *LittleDuckParser) Cte() (localctx ICteContext) {
	localctx = NewCteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LittleDuckParserRULE_cte)
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

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LittleDuckParserLET_ID)|(1<<LittleDuckParserLET_INT)|(1<<LittleDuckParserLET_FLOAT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEmptyContext is an interface to support dynamic dispatch.
type IEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyContext differentiates from other interfaces.
	IsEmptyContext()
}

type EmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyContext() *EmptyContext {
	var p = new(EmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LittleDuckParserRULE_empty
	return p
}

func (*EmptyContext) IsEmptyContext() {}

func NewEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyContext {
	var p = new(EmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LittleDuckParserRULE_empty

	return p
}

func (s *EmptyContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LittleDuckListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (p *LittleDuckParser) Empty() (localctx IEmptyContext) {
	localctx = NewEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LittleDuckParserRULE_empty)

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

	return localctx
}
