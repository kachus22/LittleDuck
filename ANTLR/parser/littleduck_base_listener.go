// Code generated from LittleDuck.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LittleDuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLittleDuckListener is a complete listener for a parse tree produced by LittleDuckParser.
type BaseLittleDuckListener struct{}

var _ LittleDuckListener = &BaseLittleDuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLittleDuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLittleDuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLittleDuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLittleDuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseLittleDuckListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseLittleDuckListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterPrograma_vars is called when production programa_vars is entered.
func (s *BaseLittleDuckListener) EnterPrograma_vars(ctx *Programa_varsContext) {}

// ExitPrograma_vars is called when production programa_vars is exited.
func (s *BaseLittleDuckListener) ExitPrograma_vars(ctx *Programa_varsContext) {}

// EnterVars is called when production vars is entered.
func (s *BaseLittleDuckListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseLittleDuckListener) ExitVars(ctx *VarsContext) {}

// EnterVars_id is called when production vars_id is entered.
func (s *BaseLittleDuckListener) EnterVars_id(ctx *Vars_idContext) {}

// ExitVars_id is called when production vars_id is exited.
func (s *BaseLittleDuckListener) ExitVars_id(ctx *Vars_idContext) {}

// EnterVars_id_next is called when production vars_id_next is entered.
func (s *BaseLittleDuckListener) EnterVars_id_next(ctx *Vars_id_nextContext) {}

// ExitVars_id_next is called when production vars_id_next is exited.
func (s *BaseLittleDuckListener) ExitVars_id_next(ctx *Vars_id_nextContext) {}

// EnterVars_more is called when production vars_more is entered.
func (s *BaseLittleDuckListener) EnterVars_more(ctx *Vars_moreContext) {}

// ExitVars_more is called when production vars_more is exited.
func (s *BaseLittleDuckListener) ExitVars_more(ctx *Vars_moreContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseLittleDuckListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseLittleDuckListener) ExitBloque(ctx *BloqueContext) {}

// EnterBloque_more is called when production bloque_more is entered.
func (s *BaseLittleDuckListener) EnterBloque_more(ctx *Bloque_moreContext) {}

// ExitBloque_more is called when production bloque_more is exited.
func (s *BaseLittleDuckListener) ExitBloque_more(ctx *Bloque_moreContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseLittleDuckListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseLittleDuckListener) ExitTipo(ctx *TipoContext) {}

// EnterEstatuto is called when production estatuto is entered.
func (s *BaseLittleDuckListener) EnterEstatuto(ctx *EstatutoContext) {}

// ExitEstatuto is called when production estatuto is exited.
func (s *BaseLittleDuckListener) ExitEstatuto(ctx *EstatutoContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseLittleDuckListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseLittleDuckListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterCondicion is called when production condicion is entered.
func (s *BaseLittleDuckListener) EnterCondicion(ctx *CondicionContext) {}

// ExitCondicion is called when production condicion is exited.
func (s *BaseLittleDuckListener) ExitCondicion(ctx *CondicionContext) {}

// EnterCondicion_else is called when production condicion_else is entered.
func (s *BaseLittleDuckListener) EnterCondicion_else(ctx *Condicion_elseContext) {}

// ExitCondicion_else is called when production condicion_else is exited.
func (s *BaseLittleDuckListener) ExitCondicion_else(ctx *Condicion_elseContext) {}

// EnterEscritura is called when production escritura is entered.
func (s *BaseLittleDuckListener) EnterEscritura(ctx *EscrituraContext) {}

// ExitEscritura is called when production escritura is exited.
func (s *BaseLittleDuckListener) ExitEscritura(ctx *EscrituraContext) {}

// EnterEscritura_params is called when production escritura_params is entered.
func (s *BaseLittleDuckListener) EnterEscritura_params(ctx *Escritura_paramsContext) {}

// ExitEscritura_params is called when production escritura_params is exited.
func (s *BaseLittleDuckListener) ExitEscritura_params(ctx *Escritura_paramsContext) {}

// EnterEscritura_params_next is called when production escritura_params_next is entered.
func (s *BaseLittleDuckListener) EnterEscritura_params_next(ctx *Escritura_params_nextContext) {}

// ExitEscritura_params_next is called when production escritura_params_next is exited.
func (s *BaseLittleDuckListener) ExitEscritura_params_next(ctx *Escritura_params_nextContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseLittleDuckListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseLittleDuckListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterExpresion_next is called when production expresion_next is entered.
func (s *BaseLittleDuckListener) EnterExpresion_next(ctx *Expresion_nextContext) {}

// ExitExpresion_next is called when production expresion_next is exited.
func (s *BaseLittleDuckListener) ExitExpresion_next(ctx *Expresion_nextContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseLittleDuckListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseLittleDuckListener) ExitExp(ctx *ExpContext) {}

// EnterExp_next is called when production exp_next is entered.
func (s *BaseLittleDuckListener) EnterExp_next(ctx *Exp_nextContext) {}

// ExitExp_next is called when production exp_next is exited.
func (s *BaseLittleDuckListener) ExitExp_next(ctx *Exp_nextContext) {}

// EnterTermino is called when production termino is entered.
func (s *BaseLittleDuckListener) EnterTermino(ctx *TerminoContext) {}

// ExitTermino is called when production termino is exited.
func (s *BaseLittleDuckListener) ExitTermino(ctx *TerminoContext) {}

// EnterTermino_next is called when production termino_next is entered.
func (s *BaseLittleDuckListener) EnterTermino_next(ctx *Termino_nextContext) {}

// ExitTermino_next is called when production termino_next is exited.
func (s *BaseLittleDuckListener) ExitTermino_next(ctx *Termino_nextContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseLittleDuckListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseLittleDuckListener) ExitFactor(ctx *FactorContext) {}

// EnterFactor_sign is called when production factor_sign is entered.
func (s *BaseLittleDuckListener) EnterFactor_sign(ctx *Factor_signContext) {}

// ExitFactor_sign is called when production factor_sign is exited.
func (s *BaseLittleDuckListener) ExitFactor_sign(ctx *Factor_signContext) {}

// EnterCte is called when production cte is entered.
func (s *BaseLittleDuckListener) EnterCte(ctx *CteContext) {}

// ExitCte is called when production cte is exited.
func (s *BaseLittleDuckListener) ExitCte(ctx *CteContext) {}

// EnterEmpty is called when production empty is entered.
func (s *BaseLittleDuckListener) EnterEmpty(ctx *EmptyContext) {}

// ExitEmpty is called when production empty is exited.
func (s *BaseLittleDuckListener) ExitEmpty(ctx *EmptyContext) {}
