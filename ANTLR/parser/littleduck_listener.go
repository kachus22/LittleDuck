// Code generated from LittleDuck.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LittleDuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LittleDuckListener is a complete listener for a parse tree produced by LittleDuckParser.
type LittleDuckListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterPrograma_vars is called when entering the programa_vars production.
	EnterPrograma_vars(c *Programa_varsContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterVars_id is called when entering the vars_id production.
	EnterVars_id(c *Vars_idContext)

	// EnterVars_id_next is called when entering the vars_id_next production.
	EnterVars_id_next(c *Vars_id_nextContext)

	// EnterVars_more is called when entering the vars_more production.
	EnterVars_more(c *Vars_moreContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBloque_more is called when entering the bloque_more production.
	EnterBloque_more(c *Bloque_moreContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterEstatuto is called when entering the estatuto production.
	EnterEstatuto(c *EstatutoContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterCondicion is called when entering the condicion production.
	EnterCondicion(c *CondicionContext)

	// EnterCondicion_else is called when entering the condicion_else production.
	EnterCondicion_else(c *Condicion_elseContext)

	// EnterEscritura is called when entering the escritura production.
	EnterEscritura(c *EscrituraContext)

	// EnterEscritura_params is called when entering the escritura_params production.
	EnterEscritura_params(c *Escritura_paramsContext)

	// EnterEscritura_params_next is called when entering the escritura_params_next production.
	EnterEscritura_params_next(c *Escritura_params_nextContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterExpresion_next is called when entering the expresion_next production.
	EnterExpresion_next(c *Expresion_nextContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterExp_next is called when entering the exp_next production.
	EnterExp_next(c *Exp_nextContext)

	// EnterTermino is called when entering the termino production.
	EnterTermino(c *TerminoContext)

	// EnterTermino_next is called when entering the termino_next production.
	EnterTermino_next(c *Termino_nextContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFactor_sign is called when entering the factor_sign production.
	EnterFactor_sign(c *Factor_signContext)

	// EnterCte is called when entering the cte production.
	EnterCte(c *CteContext)

	// EnterEmpty is called when entering the empty production.
	EnterEmpty(c *EmptyContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitPrograma_vars is called when exiting the programa_vars production.
	ExitPrograma_vars(c *Programa_varsContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitVars_id is called when exiting the vars_id production.
	ExitVars_id(c *Vars_idContext)

	// ExitVars_id_next is called when exiting the vars_id_next production.
	ExitVars_id_next(c *Vars_id_nextContext)

	// ExitVars_more is called when exiting the vars_more production.
	ExitVars_more(c *Vars_moreContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBloque_more is called when exiting the bloque_more production.
	ExitBloque_more(c *Bloque_moreContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitEstatuto is called when exiting the estatuto production.
	ExitEstatuto(c *EstatutoContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitCondicion is called when exiting the condicion production.
	ExitCondicion(c *CondicionContext)

	// ExitCondicion_else is called when exiting the condicion_else production.
	ExitCondicion_else(c *Condicion_elseContext)

	// ExitEscritura is called when exiting the escritura production.
	ExitEscritura(c *EscrituraContext)

	// ExitEscritura_params is called when exiting the escritura_params production.
	ExitEscritura_params(c *Escritura_paramsContext)

	// ExitEscritura_params_next is called when exiting the escritura_params_next production.
	ExitEscritura_params_next(c *Escritura_params_nextContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitExpresion_next is called when exiting the expresion_next production.
	ExitExpresion_next(c *Expresion_nextContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitExp_next is called when exiting the exp_next production.
	ExitExp_next(c *Exp_nextContext)

	// ExitTermino is called when exiting the termino production.
	ExitTermino(c *TerminoContext)

	// ExitTermino_next is called when exiting the termino_next production.
	ExitTermino_next(c *Termino_nextContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFactor_sign is called when exiting the factor_sign production.
	ExitFactor_sign(c *Factor_signContext)

	// ExitCte is called when exiting the cte production.
	ExitCte(c *CteContext)

	// ExitEmpty is called when exiting the empty production.
	ExitEmpty(c *EmptyContext)
}
