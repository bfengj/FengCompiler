// Code generated from ./g4/C0.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C0

import "github.com/antlr/antlr4/runtime/Go/antlr"

// C0Listener is a complete listener for a parse tree produced by C0Parser.
type C0Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterFuncDeclStmt is called when entering the funcDeclStmt production.
	EnterFuncDeclStmt(c *FuncDeclStmtContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterParaList is called when entering the paraList production.
	EnterParaList(c *ParaListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterVarDelcId is called when entering the varDelcId production.
	EnterVarDelcId(c *VarDelcIdContext)

	// EnterFuncBody is called when entering the funcBody production.
	EnterFuncBody(c *FuncBodyContext)

	// EnterVarDeclStmt is called when entering the varDeclStmt production.
	EnterVarDeclStmt(c *VarDeclStmtContext)

	// EnterVarDef is called when entering the varDef production.
	EnterVarDef(c *VarDefContext)

	// EnterConstDeclStmt is called when entering the constDeclStmt production.
	EnterConstDeclStmt(c *ConstDeclStmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterFuncCallStmt is called when entering the funcCallStmt production.
	EnterFuncCallStmt(c *FuncCallStmtContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterIntliteral is called when entering the intliteral production.
	EnterIntliteral(c *IntliteralContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitFuncDeclStmt is called when exiting the funcDeclStmt production.
	ExitFuncDeclStmt(c *FuncDeclStmtContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitParaList is called when exiting the paraList production.
	ExitParaList(c *ParaListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitVarDelcId is called when exiting the varDelcId production.
	ExitVarDelcId(c *VarDelcIdContext)

	// ExitFuncBody is called when exiting the funcBody production.
	ExitFuncBody(c *FuncBodyContext)

	// ExitVarDeclStmt is called when exiting the varDeclStmt production.
	ExitVarDeclStmt(c *VarDeclStmtContext)

	// ExitVarDef is called when exiting the varDef production.
	ExitVarDef(c *VarDefContext)

	// ExitConstDeclStmt is called when exiting the constDeclStmt production.
	ExitConstDeclStmt(c *ConstDeclStmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitFuncCallStmt is called when exiting the funcCallStmt production.
	ExitFuncCallStmt(c *FuncCallStmtContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitIntliteral is called when exiting the intliteral production.
	ExitIntliteral(c *IntliteralContext)
}
