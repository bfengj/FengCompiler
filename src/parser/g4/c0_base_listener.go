// Code generated from ./g4/C0.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C0

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseC0Listener is a complete listener for a parse tree produced by C0Parser.
type BaseC0Listener struct{}

var _ C0Listener = &BaseC0Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseC0Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseC0Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseC0Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseC0Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseC0Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseC0Listener) ExitProg(ctx *ProgContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BaseC0Listener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BaseC0Listener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseC0Listener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseC0Listener) ExitStmt(ctx *StmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseC0Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseC0Listener) ExitBlock(ctx *BlockContext) {}

// EnterFuncDeclStmt is called when production funcDeclStmt is entered.
func (s *BaseC0Listener) EnterFuncDeclStmt(ctx *FuncDeclStmtContext) {}

// ExitFuncDeclStmt is called when production funcDeclStmt is exited.
func (s *BaseC0Listener) ExitFuncDeclStmt(ctx *FuncDeclStmtContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BaseC0Listener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BaseC0Listener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseC0Listener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseC0Listener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseC0Listener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseC0Listener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseC0Listener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseC0Listener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterParaList is called when production paraList is entered.
func (s *BaseC0Listener) EnterParaList(ctx *ParaListContext) {}

// ExitParaList is called when production paraList is exited.
func (s *BaseC0Listener) ExitParaList(ctx *ParaListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseC0Listener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseC0Listener) ExitParameter(ctx *ParameterContext) {}

// EnterVarDelcId is called when production varDelcId is entered.
func (s *BaseC0Listener) EnterVarDelcId(ctx *VarDelcIdContext) {}

// ExitVarDelcId is called when production varDelcId is exited.
func (s *BaseC0Listener) ExitVarDelcId(ctx *VarDelcIdContext) {}

// EnterFuncBody is called when production funcBody is entered.
func (s *BaseC0Listener) EnterFuncBody(ctx *FuncBodyContext) {}

// ExitFuncBody is called when production funcBody is exited.
func (s *BaseC0Listener) ExitFuncBody(ctx *FuncBodyContext) {}

// EnterVarDeclStmt is called when production varDeclStmt is entered.
func (s *BaseC0Listener) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production varDeclStmt is exited.
func (s *BaseC0Listener) ExitVarDeclStmt(ctx *VarDeclStmtContext) {}

// EnterVarDef is called when production varDef is entered.
func (s *BaseC0Listener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production varDef is exited.
func (s *BaseC0Listener) ExitVarDef(ctx *VarDefContext) {}

// EnterConstDeclStmt is called when production constDeclStmt is entered.
func (s *BaseC0Listener) EnterConstDeclStmt(ctx *ConstDeclStmtContext) {}

// ExitConstDeclStmt is called when production constDeclStmt is exited.
func (s *BaseC0Listener) ExitConstDeclStmt(ctx *ConstDeclStmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseC0Listener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseC0Listener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseC0Listener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseC0Listener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseC0Listener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseC0Listener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterFuncCallStmt is called when production funcCallStmt is entered.
func (s *BaseC0Listener) EnterFuncCallStmt(ctx *FuncCallStmtContext) {}

// ExitFuncCallStmt is called when production funcCallStmt is exited.
func (s *BaseC0Listener) ExitFuncCallStmt(ctx *FuncCallStmtContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseC0Listener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseC0Listener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseC0Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseC0Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseC0Listener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseC0Listener) ExitPrimary(ctx *PrimaryContext) {}

// EnterIntliteral is called when production intliteral is entered.
func (s *BaseC0Listener) EnterIntliteral(ctx *IntliteralContext) {}

// ExitIntliteral is called when production intliteral is exited.
func (s *BaseC0Listener) ExitIntliteral(ctx *IntliteralContext) {}
