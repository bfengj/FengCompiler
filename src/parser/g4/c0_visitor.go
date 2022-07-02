// Code generated from ./g4/C0.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C0

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by C0Parser.
type C0Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by C0Parser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by C0Parser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by C0Parser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by C0Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by C0Parser#funcDeclStmt.
	VisitFuncDeclStmt(ctx *FuncDeclStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{}

	// Visit a parse tree produced by C0Parser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by C0Parser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by C0Parser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by C0Parser#paraList.
	VisitParaList(ctx *ParaListContext) interface{}

	// Visit a parse tree produced by C0Parser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by C0Parser#varDelcId.
	VisitVarDelcId(ctx *VarDelcIdContext) interface{}

	// Visit a parse tree produced by C0Parser#funcBody.
	VisitFuncBody(ctx *FuncBodyContext) interface{}

	// Visit a parse tree produced by C0Parser#varDeclStmt.
	VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#varDef.
	VisitVarDef(ctx *VarDefContext) interface{}

	// Visit a parse tree produced by C0Parser#constDeclStmt.
	VisitConstDeclStmt(ctx *ConstDeclStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#funcCallStmt.
	VisitFuncCallStmt(ctx *FuncCallStmtContext) interface{}

	// Visit a parse tree produced by C0Parser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by C0Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by C0Parser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by C0Parser#intliteral.
	VisitIntliteral(ctx *IntliteralContext) interface{}
}
