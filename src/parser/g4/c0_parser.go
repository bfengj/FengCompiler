// Code generated from ./g4/C0.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C0

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type C0Parser struct {
	*antlr.BaseParser
}

var c0ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func c0ParserInit() {
	staticData := &c0ParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "'int'", "'char'", "'if'", "'else'", "'while'", "'true'",
		"'false'", "'void'", "'return'", "'const'", "'break'", "'continue'",
		"", "", "", "", "", "','", "';'", "':'", "'{'", "'}'", "'('", "')'",
		"'['", "']'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'++'", "'--'",
		"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'!'", "'&&'", "'||'",
		"'//'", "'/*'", "'*/'",
	}
	staticData.symbolicNames = []string{
		"", "COMMENT", "WS", "NEWLINE", "INT", "CHAR", "IF", "ELSE", "WHILE",
		"TRUE", "FALSE", "VOID", "RETURN", "CONST", "BREAK", "CONTINUE", "DECIMAL",
		"OCTAL", "HEXADECIMAL", "HEXADECIMALPREFIX", "IDENTIFIER", "COMMA",
		"SEMICOLON", "COLON", "LBRACE", "RBRACE", "LPARENTHESIS", "RPARENTHESIS",
		"LSQUAREBRACKET", "RSQUAREBRACKET", "ASSIGN", "PLUS", "MINUS", "MUL",
		"DIV", "MOD", "MULPLUS", "MULMINUS", "EQ", "NE", "LT", "GT", "LE", "GE",
		"NOT", "AND", "OR", "LINECOMMENT", "BEGINCOMMENT", "ENDCOMMENT",
	}
	staticData.ruleNames = []string{
		"prog", "stmtList", "stmt", "block", "funcDeclStmt", "typeTypeOrVoid",
		"typeType", "baseType", "formalParameters", "paraList", "parameter",
		"varDelcId", "funcBody", "varDeclStmt", "varDef", "constDeclStmt", "assignStmt",
		"ifStmt", "whileStmt", "funcCallStmt", "expressionList", "expression",
		"primary", "intliteral",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 293, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 5, 1, 52, 8, 1,
		10, 1, 12, 1, 55, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 71, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 79, 8, 2, 1, 3, 1, 3, 5, 3, 83, 8, 3, 10, 3, 12, 3,
		86, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 97,
		8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 102, 8, 6, 10, 6, 12, 6, 105, 9, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 3, 8, 111, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9,
		118, 8, 9, 10, 9, 12, 9, 121, 9, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 134, 8, 11, 10, 11, 12,
		11, 137, 9, 11, 3, 11, 139, 8, 11, 1, 12, 1, 12, 3, 12, 143, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 5, 13, 149, 8, 13, 10, 13, 12, 13, 152, 9, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 161, 8, 14, 10,
		14, 12, 14, 164, 9, 14, 1, 14, 3, 14, 167, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 5, 15, 174, 8, 15, 10, 15, 12, 15, 177, 9, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 186, 8, 16, 10, 16, 12, 16,
		189, 9, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 201, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 3, 19, 212, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 5, 20, 219, 8, 20, 10, 20, 12, 20, 222, 9, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 234, 8, 21, 10,
		21, 12, 21, 237, 9, 21, 3, 21, 239, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 247, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 272, 8, 21, 11,
		21, 12, 21, 273, 1, 21, 1, 21, 5, 21, 278, 8, 21, 10, 21, 12, 21, 281,
		9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 289, 8, 22, 1,
		23, 1, 23, 1, 23, 0, 1, 42, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 0, 7, 2, 0, 31, 32,
		36, 37, 1, 0, 33, 35, 1, 0, 31, 32, 1, 0, 40, 43, 1, 0, 38, 39, 1, 0, 36,
		37, 1, 0, 16, 18, 315, 0, 48, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 78, 1,
		0, 0, 0, 6, 80, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12,
		98, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 108, 1, 0, 0, 0, 18, 114, 1, 0,
		0, 0, 20, 122, 1, 0, 0, 0, 22, 138, 1, 0, 0, 0, 24, 142, 1, 0, 0, 0, 26,
		144, 1, 0, 0, 0, 28, 166, 1, 0, 0, 0, 30, 168, 1, 0, 0, 0, 32, 180, 1,
		0, 0, 0, 34, 193, 1, 0, 0, 0, 36, 202, 1, 0, 0, 0, 38, 208, 1, 0, 0, 0,
		40, 215, 1, 0, 0, 0, 42, 246, 1, 0, 0, 0, 44, 288, 1, 0, 0, 0, 46, 290,
		1, 0, 0, 0, 48, 49, 3, 2, 1, 0, 49, 1, 1, 0, 0, 0, 50, 52, 3, 4, 2, 0,
		51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1,
		0, 0, 0, 54, 3, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 79, 3, 26, 13, 0, 57,
		79, 3, 30, 15, 0, 58, 59, 3, 32, 16, 0, 59, 60, 5, 22, 0, 0, 60, 79, 1,
		0, 0, 0, 61, 79, 3, 8, 4, 0, 62, 63, 3, 38, 19, 0, 63, 64, 5, 22, 0, 0,
		64, 79, 1, 0, 0, 0, 65, 79, 3, 6, 3, 0, 66, 79, 3, 34, 17, 0, 67, 79, 3,
		36, 18, 0, 68, 70, 5, 12, 0, 0, 69, 71, 3, 42, 21, 0, 70, 69, 1, 0, 0,
		0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 79, 5, 22, 0, 0, 73, 74,
		5, 14, 0, 0, 74, 79, 5, 22, 0, 0, 75, 76, 5, 15, 0, 0, 76, 79, 5, 22, 0,
		0, 77, 79, 5, 22, 0, 0, 78, 56, 1, 0, 0, 0, 78, 57, 1, 0, 0, 0, 78, 58,
		1, 0, 0, 0, 78, 61, 1, 0, 0, 0, 78, 62, 1, 0, 0, 0, 78, 65, 1, 0, 0, 0,
		78, 66, 1, 0, 0, 0, 78, 67, 1, 0, 0, 0, 78, 68, 1, 0, 0, 0, 78, 73, 1,
		0, 0, 0, 78, 75, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 5, 1, 0, 0, 0, 80,
		84, 5, 24, 0, 0, 81, 83, 3, 4, 2, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 87, 88, 5, 25, 0, 0, 88, 7, 1, 0, 0, 0, 89, 90, 3, 10, 5, 0,
		90, 91, 5, 20, 0, 0, 91, 92, 3, 16, 8, 0, 92, 93, 3, 24, 12, 0, 93, 9,
		1, 0, 0, 0, 94, 97, 3, 12, 6, 0, 95, 97, 5, 11, 0, 0, 96, 94, 1, 0, 0,
		0, 96, 95, 1, 0, 0, 0, 97, 11, 1, 0, 0, 0, 98, 103, 3, 14, 7, 0, 99, 100,
		5, 28, 0, 0, 100, 102, 5, 29, 0, 0, 101, 99, 1, 0, 0, 0, 102, 105, 1, 0,
		0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 13, 1, 0, 0, 0,
		105, 103, 1, 0, 0, 0, 106, 107, 5, 4, 0, 0, 107, 15, 1, 0, 0, 0, 108, 110,
		5, 26, 0, 0, 109, 111, 3, 18, 9, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1,
		0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 27, 0, 0, 113, 17, 1, 0, 0,
		0, 114, 119, 3, 20, 10, 0, 115, 116, 5, 21, 0, 0, 116, 118, 3, 20, 10,
		0, 117, 115, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 19, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 3,
		14, 7, 0, 123, 124, 3, 22, 11, 0, 124, 21, 1, 0, 0, 0, 125, 139, 5, 20,
		0, 0, 126, 127, 5, 20, 0, 0, 127, 128, 5, 28, 0, 0, 128, 135, 5, 29, 0,
		0, 129, 130, 5, 28, 0, 0, 130, 131, 3, 42, 21, 0, 131, 132, 5, 29, 0, 0,
		132, 134, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135,
		133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 138, 125, 1, 0, 0, 0, 138, 126, 1, 0, 0, 0, 139, 23, 1, 0,
		0, 0, 140, 143, 3, 6, 3, 0, 141, 143, 5, 22, 0, 0, 142, 140, 1, 0, 0, 0,
		142, 141, 1, 0, 0, 0, 143, 25, 1, 0, 0, 0, 144, 145, 3, 14, 7, 0, 145,
		150, 3, 28, 14, 0, 146, 147, 5, 21, 0, 0, 147, 149, 3, 28, 14, 0, 148,
		146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 22,
		0, 0, 154, 27, 1, 0, 0, 0, 155, 162, 5, 20, 0, 0, 156, 157, 5, 28, 0, 0,
		157, 158, 3, 42, 21, 0, 158, 159, 5, 29, 0, 0, 159, 161, 1, 0, 0, 0, 160,
		156, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163,
		1, 0, 0, 0, 163, 167, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 167, 3, 32,
		16, 0, 166, 155, 1, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 29, 1, 0, 0, 0,
		168, 169, 5, 13, 0, 0, 169, 170, 3, 14, 7, 0, 170, 175, 3, 32, 16, 0, 171,
		172, 5, 21, 0, 0, 172, 174, 3, 32, 16, 0, 173, 171, 1, 0, 0, 0, 174, 177,
		1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0,
		0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 5, 22, 0, 0, 179, 31, 1, 0, 0, 0,
		180, 187, 5, 20, 0, 0, 181, 182, 5, 28, 0, 0, 182, 183, 3, 42, 21, 0, 183,
		184, 5, 29, 0, 0, 184, 186, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 186, 189,
		1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 190, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 30, 0, 0, 191, 192, 3, 42, 21,
		0, 192, 33, 1, 0, 0, 0, 193, 194, 5, 6, 0, 0, 194, 195, 5, 26, 0, 0, 195,
		196, 3, 42, 21, 0, 196, 197, 5, 27, 0, 0, 197, 200, 3, 4, 2, 0, 198, 199,
		5, 7, 0, 0, 199, 201, 3, 4, 2, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0,
		0, 0, 201, 35, 1, 0, 0, 0, 202, 203, 5, 8, 0, 0, 203, 204, 5, 26, 0, 0,
		204, 205, 3, 42, 21, 0, 205, 206, 5, 27, 0, 0, 206, 207, 3, 4, 2, 0, 207,
		37, 1, 0, 0, 0, 208, 209, 5, 20, 0, 0, 209, 211, 5, 26, 0, 0, 210, 212,
		3, 40, 20, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1,
		0, 0, 0, 213, 214, 5, 27, 0, 0, 214, 39, 1, 0, 0, 0, 215, 220, 3, 42, 21,
		0, 216, 217, 5, 21, 0, 0, 217, 219, 3, 42, 21, 0, 218, 216, 1, 0, 0, 0,
		219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		41, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223, 224, 6, 21, -1, 0, 224, 225,
		5, 26, 0, 0, 225, 226, 3, 42, 21, 0, 226, 227, 5, 27, 0, 0, 227, 247, 1,
		0, 0, 0, 228, 247, 3, 44, 22, 0, 229, 238, 5, 24, 0, 0, 230, 235, 3, 42,
		21, 0, 231, 232, 5, 21, 0, 0, 232, 234, 3, 42, 21, 0, 233, 231, 1, 0, 0,
		0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236,
		239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 230, 1, 0, 0, 0, 238, 239,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 247, 5, 25, 0, 0, 241, 242, 7, 0,
		0, 0, 242, 247, 3, 42, 21, 9, 243, 244, 5, 44, 0, 0, 244, 247, 3, 42, 21,
		8, 245, 247, 3, 38, 19, 0, 246, 223, 1, 0, 0, 0, 246, 228, 1, 0, 0, 0,
		246, 229, 1, 0, 0, 0, 246, 241, 1, 0, 0, 0, 246, 243, 1, 0, 0, 0, 246,
		245, 1, 0, 0, 0, 247, 279, 1, 0, 0, 0, 248, 249, 10, 7, 0, 0, 249, 250,
		7, 1, 0, 0, 250, 278, 3, 42, 21, 8, 251, 252, 10, 6, 0, 0, 252, 253, 7,
		2, 0, 0, 253, 278, 3, 42, 21, 7, 254, 255, 10, 5, 0, 0, 255, 256, 7, 3,
		0, 0, 256, 278, 3, 42, 21, 6, 257, 258, 10, 4, 0, 0, 258, 259, 7, 4, 0,
		0, 259, 278, 3, 42, 21, 5, 260, 261, 10, 3, 0, 0, 261, 262, 5, 45, 0, 0,
		262, 278, 3, 42, 21, 4, 263, 264, 10, 2, 0, 0, 264, 265, 5, 46, 0, 0, 265,
		278, 3, 42, 21, 3, 266, 271, 10, 12, 0, 0, 267, 268, 5, 28, 0, 0, 268,
		269, 3, 42, 21, 0, 269, 270, 5, 29, 0, 0, 270, 272, 1, 0, 0, 0, 271, 267,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0,
		0, 0, 274, 278, 1, 0, 0, 0, 275, 276, 10, 10, 0, 0, 276, 278, 7, 5, 0,
		0, 277, 248, 1, 0, 0, 0, 277, 251, 1, 0, 0, 0, 277, 254, 1, 0, 0, 0, 277,
		257, 1, 0, 0, 0, 277, 260, 1, 0, 0, 0, 277, 263, 1, 0, 0, 0, 277, 266,
		1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0,
		0, 0, 279, 280, 1, 0, 0, 0, 280, 43, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0,
		282, 283, 5, 26, 0, 0, 283, 284, 3, 42, 21, 0, 284, 285, 5, 27, 0, 0, 285,
		289, 1, 0, 0, 0, 286, 289, 5, 20, 0, 0, 287, 289, 3, 46, 23, 0, 288, 282,
		1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 287, 1, 0, 0, 0, 289, 45, 1, 0,
		0, 0, 290, 291, 7, 6, 0, 0, 291, 47, 1, 0, 0, 0, 26, 53, 70, 78, 84, 96,
		103, 110, 119, 135, 138, 142, 150, 162, 166, 175, 187, 200, 211, 220, 235,
		238, 246, 273, 277, 279, 288,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// C0ParserInit initializes any static state used to implement C0Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewC0Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func C0ParserInit() {
	staticData := &c0ParserStaticData
	staticData.once.Do(c0ParserInit)
}

// NewC0Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewC0Parser(input antlr.TokenStream) *C0Parser {
	C0ParserInit()
	this := new(C0Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &c0ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "C0.g4"

	return this
}

// C0Parser tokens.
const (
	C0ParserEOF               = antlr.TokenEOF
	C0ParserCOMMENT           = 1
	C0ParserWS                = 2
	C0ParserNEWLINE           = 3
	C0ParserINT               = 4
	C0ParserCHAR              = 5
	C0ParserIF                = 6
	C0ParserELSE              = 7
	C0ParserWHILE             = 8
	C0ParserTRUE              = 9
	C0ParserFALSE             = 10
	C0ParserVOID              = 11
	C0ParserRETURN            = 12
	C0ParserCONST             = 13
	C0ParserBREAK             = 14
	C0ParserCONTINUE          = 15
	C0ParserDECIMAL           = 16
	C0ParserOCTAL             = 17
	C0ParserHEXADECIMAL       = 18
	C0ParserHEXADECIMALPREFIX = 19
	C0ParserIDENTIFIER        = 20
	C0ParserCOMMA             = 21
	C0ParserSEMICOLON         = 22
	C0ParserCOLON             = 23
	C0ParserLBRACE            = 24
	C0ParserRBRACE            = 25
	C0ParserLPARENTHESIS      = 26
	C0ParserRPARENTHESIS      = 27
	C0ParserLSQUAREBRACKET    = 28
	C0ParserRSQUAREBRACKET    = 29
	C0ParserASSIGN            = 30
	C0ParserPLUS              = 31
	C0ParserMINUS             = 32
	C0ParserMUL               = 33
	C0ParserDIV               = 34
	C0ParserMOD               = 35
	C0ParserMULPLUS           = 36
	C0ParserMULMINUS          = 37
	C0ParserEQ                = 38
	C0ParserNE                = 39
	C0ParserLT                = 40
	C0ParserGT                = 41
	C0ParserLE                = 42
	C0ParserGE                = 43
	C0ParserNOT               = 44
	C0ParserAND               = 45
	C0ParserOR                = 46
	C0ParserLINECOMMENT       = 47
	C0ParserBEGINCOMMENT      = 48
	C0ParserENDCOMMENT        = 49
)

// C0Parser rules.
const (
	C0ParserRULE_prog             = 0
	C0ParserRULE_stmtList         = 1
	C0ParserRULE_stmt             = 2
	C0ParserRULE_block            = 3
	C0ParserRULE_funcDeclStmt     = 4
	C0ParserRULE_typeTypeOrVoid   = 5
	C0ParserRULE_typeType         = 6
	C0ParserRULE_baseType         = 7
	C0ParserRULE_formalParameters = 8
	C0ParserRULE_paraList         = 9
	C0ParserRULE_parameter        = 10
	C0ParserRULE_varDelcId        = 11
	C0ParserRULE_funcBody         = 12
	C0ParserRULE_varDeclStmt      = 13
	C0ParserRULE_varDef           = 14
	C0ParserRULE_constDeclStmt    = 15
	C0ParserRULE_assignStmt       = 16
	C0ParserRULE_ifStmt           = 17
	C0ParserRULE_whileStmt        = 18
	C0ParserRULE_funcCallStmt     = 19
	C0ParserRULE_expressionList   = 20
	C0ParserRULE_expression       = 21
	C0ParserRULE_primary          = 22
	C0ParserRULE_intliteral       = 23
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, C0ParserRULE_prog)

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
		p.SetState(48)
		p.StmtList()
	}

	return localctx
}

// IStmtListContext is an interface to support dynamic dispatch.
type IStmtListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtListContext differentiates from other interfaces.
	IsStmtListContext()
}

type StmtListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtListContext() *StmtListContext {
	var p = new(StmtListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_stmtList
	return p
}

func (*StmtListContext) IsStmtListContext() {}

func NewStmtListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtListContext {
	var p = new(StmtListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_stmtList

	return p
}

func (s *StmtListContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtListContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtListContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterStmtList(s)
	}
}

func (s *StmtListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitStmtList(s)
	}
}

func (s *StmtListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitStmtList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) StmtList() (localctx IStmtListContext) {
	this := p
	_ = this

	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, C0ParserRULE_stmtList)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<C0ParserINT)|(1<<C0ParserIF)|(1<<C0ParserWHILE)|(1<<C0ParserVOID)|(1<<C0ParserRETURN)|(1<<C0ParserCONST)|(1<<C0ParserBREAK)|(1<<C0ParserCONTINUE)|(1<<C0ParserIDENTIFIER)|(1<<C0ParserSEMICOLON)|(1<<C0ParserLBRACE))) != 0 {
		{
			p.SetState(50)
			p.Stmt()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) VarDeclStmt() IVarDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclStmtContext)
}

func (s *StmtContext) ConstDeclStmt() IConstDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclStmtContext)
}

func (s *StmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(C0ParserSEMICOLON, 0)
}

func (s *StmtContext) FuncDeclStmt() IFuncDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclStmtContext)
}

func (s *StmtContext) FuncCallStmt() IFuncCallStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallStmtContext)
}

func (s *StmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(C0ParserRETURN, 0)
}

func (s *StmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(C0ParserBREAK, 0)
}

func (s *StmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(C0ParserCONTINUE, 0)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, C0ParserRULE_stmt)
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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.VarDeclStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.ConstDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.AssignStmt()
		}
		{
			p.SetState(59)
			p.Match(C0ParserSEMICOLON)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.FuncDeclStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.FuncCallStmt()
		}
		{
			p.SetState(63)
			p.Match(C0ParserSEMICOLON)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(66)
			p.IfStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(67)
			p.WhileStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(68)
			p.Match(C0ParserRETURN)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(C0ParserDECIMAL-16))|(1<<(C0ParserOCTAL-16))|(1<<(C0ParserHEXADECIMAL-16))|(1<<(C0ParserIDENTIFIER-16))|(1<<(C0ParserLBRACE-16))|(1<<(C0ParserLPARENTHESIS-16))|(1<<(C0ParserPLUS-16))|(1<<(C0ParserMINUS-16))|(1<<(C0ParserMULPLUS-16))|(1<<(C0ParserMULMINUS-16))|(1<<(C0ParserNOT-16)))) != 0 {
			{
				p.SetState(69)
				p.expression(0)
			}

		}
		{
			p.SetState(72)
			p.Match(C0ParserSEMICOLON)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(73)
			p.Match(C0ParserBREAK)
		}
		{
			p.SetState(74)
			p.Match(C0ParserSEMICOLON)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(75)
			p.Match(C0ParserCONTINUE)
		}
		{
			p.SetState(76)
			p.Match(C0ParserSEMICOLON)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(77)
			p.Match(C0ParserSEMICOLON)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(C0ParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(C0ParserRBRACE, 0)
}

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, C0ParserRULE_block)
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
		p.SetState(80)
		p.Match(C0ParserLBRACE)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<C0ParserINT)|(1<<C0ParserIF)|(1<<C0ParserWHILE)|(1<<C0ParserVOID)|(1<<C0ParserRETURN)|(1<<C0ParserCONST)|(1<<C0ParserBREAK)|(1<<C0ParserCONTINUE)|(1<<C0ParserIDENTIFIER)|(1<<C0ParserSEMICOLON)|(1<<C0ParserLBRACE))) != 0 {
		{
			p.SetState(81)
			p.Stmt()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(C0ParserRBRACE)
	}

	return localctx
}

// IFuncDeclStmtContext is an interface to support dynamic dispatch.
type IFuncDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclStmtContext differentiates from other interfaces.
	IsFuncDeclStmtContext()
}

type FuncDeclStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclStmtContext() *FuncDeclStmtContext {
	var p = new(FuncDeclStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_funcDeclStmt
	return p
}

func (*FuncDeclStmtContext) IsFuncDeclStmtContext() {}

func NewFuncDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclStmtContext {
	var p = new(FuncDeclStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_funcDeclStmt

	return p
}

func (s *FuncDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclStmtContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeOrVoidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FuncDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *FuncDeclStmtContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FuncDeclStmtContext) FuncBody() IFuncBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterFuncDeclStmt(s)
	}
}

func (s *FuncDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitFuncDeclStmt(s)
	}
}

func (s *FuncDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitFuncDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) FuncDeclStmt() (localctx IFuncDeclStmtContext) {
	this := p
	_ = this

	localctx = NewFuncDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, C0ParserRULE_funcDeclStmt)

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
		p.SetState(89)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(90)
		p.Match(C0ParserIDENTIFIER)
	}
	{
		p.SetState(91)
		p.FormalParameters()
	}
	{
		p.SetState(92)
		p.FuncBody()
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(C0ParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	this := p
	_ = this

	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, C0ParserRULE_typeTypeOrVoid)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case C0ParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.TypeType()
		}

	case C0ParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(C0ParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeTypeContext) AllLSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserLSQUAREBRACKET)
}

func (s *TypeTypeContext) LSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserLSQUAREBRACKET, i)
}

func (s *TypeTypeContext) AllRSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserRSQUAREBRACKET)
}

func (s *TypeTypeContext) RSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserRSQUAREBRACKET, i)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterTypeType(s)
	}
}

func (s *TypeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitTypeType(s)
	}
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) TypeType() (localctx ITypeTypeContext) {
	this := p
	_ = this

	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, C0ParserRULE_typeType)
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
		p.SetState(98)
		p.BaseType()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserLSQUAREBRACKET {
		{
			p.SetState(99)
			p.Match(C0ParserLSQUAREBRACKET)
		}
		{
			p.SetState(100)
			p.Match(C0ParserRSQUAREBRACKET)
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(C0ParserINT, 0)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) BaseType() (localctx IBaseTypeContext) {
	this := p
	_ = this

	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, C0ParserRULE_baseType)

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
		p.Match(C0ParserINT)
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *FormalParametersContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *FormalParametersContext) ParaList() IParaListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParaListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParaListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) FormalParameters() (localctx IFormalParametersContext) {
	this := p
	_ = this

	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, C0ParserRULE_formalParameters)
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
		p.SetState(108)
		p.Match(C0ParserLPARENTHESIS)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == C0ParserINT {
		{
			p.SetState(109)
			p.ParaList()
		}

	}
	{
		p.SetState(112)
		p.Match(C0ParserRPARENTHESIS)
	}

	return localctx
}

// IParaListContext is an interface to support dynamic dispatch.
type IParaListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaListContext differentiates from other interfaces.
	IsParaListContext()
}

type ParaListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaListContext() *ParaListContext {
	var p = new(ParaListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_paraList
	return p
}

func (*ParaListContext) IsParaListContext() {}

func NewParaListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaListContext {
	var p = new(ParaListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_paraList

	return p
}

func (s *ParaListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParaListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParaListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(C0ParserCOMMA)
}

func (s *ParaListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserCOMMA, i)
}

func (s *ParaListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterParaList(s)
	}
}

func (s *ParaListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitParaList(s)
	}
}

func (s *ParaListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitParaList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) ParaList() (localctx IParaListContext) {
	this := p
	_ = this

	localctx = NewParaListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, C0ParserRULE_paraList)
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
		p.SetState(114)
		p.Parameter()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserCOMMA {
		{
			p.SetState(115)
			p.Match(C0ParserCOMMA)
		}
		{
			p.SetState(116)
			p.Parameter()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *ParameterContext) VarDelcId() IVarDelcIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDelcIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDelcIdContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, C0ParserRULE_parameter)

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
		p.SetState(122)
		p.BaseType()
	}
	{
		p.SetState(123)
		p.VarDelcId()
	}

	return localctx
}

// IVarDelcIdContext is an interface to support dynamic dispatch.
type IVarDelcIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDelcIdContext differentiates from other interfaces.
	IsVarDelcIdContext()
}

type VarDelcIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDelcIdContext() *VarDelcIdContext {
	var p = new(VarDelcIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_varDelcId
	return p
}

func (*VarDelcIdContext) IsVarDelcIdContext() {}

func NewVarDelcIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDelcIdContext {
	var p = new(VarDelcIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_varDelcId

	return p
}

func (s *VarDelcIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDelcIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *VarDelcIdContext) AllLSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserLSQUAREBRACKET)
}

func (s *VarDelcIdContext) LSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserLSQUAREBRACKET, i)
}

func (s *VarDelcIdContext) AllRSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserRSQUAREBRACKET)
}

func (s *VarDelcIdContext) RSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserRSQUAREBRACKET, i)
}

func (s *VarDelcIdContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VarDelcIdContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDelcIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDelcIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDelcIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterVarDelcId(s)
	}
}

func (s *VarDelcIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitVarDelcId(s)
	}
}

func (s *VarDelcIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitVarDelcId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) VarDelcId() (localctx IVarDelcIdContext) {
	this := p
	_ = this

	localctx = NewVarDelcIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, C0ParserRULE_varDelcId)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(C0ParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(C0ParserIDENTIFIER)
		}
		{
			p.SetState(127)
			p.Match(C0ParserLSQUAREBRACKET)
		}
		{
			p.SetState(128)
			p.Match(C0ParserRSQUAREBRACKET)
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == C0ParserLSQUAREBRACKET {
			{
				p.SetState(129)
				p.Match(C0ParserLSQUAREBRACKET)
			}
			{
				p.SetState(130)
				p.expression(0)
			}
			{
				p.SetState(131)
				p.Match(C0ParserRSQUAREBRACKET)
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFuncBodyContext is an interface to support dynamic dispatch.
type IFuncBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncBodyContext differentiates from other interfaces.
	IsFuncBodyContext()
}

type FuncBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBodyContext() *FuncBodyContext {
	var p = new(FuncBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_funcBody
	return p
}

func (*FuncBodyContext) IsFuncBodyContext() {}

func NewFuncBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBodyContext {
	var p = new(FuncBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_funcBody

	return p
}

func (s *FuncBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncBodyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(C0ParserSEMICOLON, 0)
}

func (s *FuncBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterFuncBody(s)
	}
}

func (s *FuncBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitFuncBody(s)
	}
}

func (s *FuncBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitFuncBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) FuncBody() (localctx IFuncBodyContext) {
	this := p
	_ = this

	localctx = NewFuncBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, C0ParserRULE_funcBody)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case C0ParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Block()
		}

	case C0ParserSEMICOLON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(C0ParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarDeclStmtContext is an interface to support dynamic dispatch.
type IVarDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclStmtContext differentiates from other interfaces.
	IsVarDeclStmtContext()
}

type VarDeclStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclStmtContext() *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_varDeclStmt
	return p
}

func (*VarDeclStmtContext) IsVarDeclStmtContext() {}

func NewVarDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_varDeclStmt

	return p
}

func (s *VarDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclStmtContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *VarDeclStmtContext) AllVarDef() []IVarDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDefContext); ok {
			len++
		}
	}

	tst := make([]IVarDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDefContext); ok {
			tst[i] = t.(IVarDefContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclStmtContext) VarDef(i int) IVarDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *VarDeclStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(C0ParserSEMICOLON, 0)
}

func (s *VarDeclStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(C0ParserCOMMA)
}

func (s *VarDeclStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserCOMMA, i)
}

func (s *VarDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitVarDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) VarDeclStmt() (localctx IVarDeclStmtContext) {
	this := p
	_ = this

	localctx = NewVarDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, C0ParserRULE_varDeclStmt)
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
		p.SetState(144)
		p.BaseType()
	}
	{
		p.SetState(145)
		p.VarDef()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserCOMMA {
		{
			p.SetState(146)
			p.Match(C0ParserCOMMA)
		}
		{
			p.SetState(147)
			p.VarDef()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(C0ParserSEMICOLON)
	}

	return localctx
}

// IVarDefContext is an interface to support dynamic dispatch.
type IVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDefContext differentiates from other interfaces.
	IsVarDefContext()
}

type VarDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDefContext() *VarDefContext {
	var p = new(VarDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_varDef
	return p
}

func (*VarDefContext) IsVarDefContext() {}

func NewVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefContext {
	var p = new(VarDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_varDef

	return p
}

func (s *VarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *VarDefContext) AllLSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserLSQUAREBRACKET)
}

func (s *VarDefContext) LSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserLSQUAREBRACKET, i)
}

func (s *VarDefContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VarDefContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDefContext) AllRSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserRSQUAREBRACKET)
}

func (s *VarDefContext) RSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserRSQUAREBRACKET, i)
}

func (s *VarDefContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterVarDef(s)
	}
}

func (s *VarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitVarDef(s)
	}
}

func (s *VarDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitVarDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) VarDef() (localctx IVarDefContext) {
	this := p
	_ = this

	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, C0ParserRULE_varDef)
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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(C0ParserIDENTIFIER)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == C0ParserLSQUAREBRACKET {
			{
				p.SetState(156)
				p.Match(C0ParserLSQUAREBRACKET)
			}
			{
				p.SetState(157)
				p.expression(0)
			}
			{
				p.SetState(158)
				p.Match(C0ParserRSQUAREBRACKET)
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.AssignStmt()
		}

	}

	return localctx
}

// IConstDeclStmtContext is an interface to support dynamic dispatch.
type IConstDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstDeclStmtContext differentiates from other interfaces.
	IsConstDeclStmtContext()
}

type ConstDeclStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclStmtContext() *ConstDeclStmtContext {
	var p = new(ConstDeclStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_constDeclStmt
	return p
}

func (*ConstDeclStmtContext) IsConstDeclStmtContext() {}

func NewConstDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclStmtContext {
	var p = new(ConstDeclStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_constDeclStmt

	return p
}

func (s *ConstDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclStmtContext) CONST() antlr.TerminalNode {
	return s.GetToken(C0ParserCONST, 0)
}

func (s *ConstDeclStmtContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *ConstDeclStmtContext) AllAssignStmt() []IAssignStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignStmtContext); ok {
			len++
		}
	}

	tst := make([]IAssignStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignStmtContext); ok {
			tst[i] = t.(IAssignStmtContext)
			i++
		}
	}

	return tst
}

func (s *ConstDeclStmtContext) AssignStmt(i int) IAssignStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ConstDeclStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(C0ParserSEMICOLON, 0)
}

func (s *ConstDeclStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(C0ParserCOMMA)
}

func (s *ConstDeclStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserCOMMA, i)
}

func (s *ConstDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterConstDeclStmt(s)
	}
}

func (s *ConstDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitConstDeclStmt(s)
	}
}

func (s *ConstDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitConstDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) ConstDeclStmt() (localctx IConstDeclStmtContext) {
	this := p
	_ = this

	localctx = NewConstDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, C0ParserRULE_constDeclStmt)
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
		p.SetState(168)
		p.Match(C0ParserCONST)
	}
	{
		p.SetState(169)
		p.BaseType()
	}
	{
		p.SetState(170)
		p.AssignStmt()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserCOMMA {
		{
			p.SetState(171)
			p.Match(C0ParserCOMMA)
		}
		{
			p.SetState(172)
			p.AssignStmt()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(C0ParserSEMICOLON)
	}

	return localctx
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_assignStmt
	return p
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *AssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(C0ParserASSIGN, 0)
}

func (s *AssignStmtContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignStmtContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignStmtContext) AllLSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserLSQUAREBRACKET)
}

func (s *AssignStmtContext) LSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserLSQUAREBRACKET, i)
}

func (s *AssignStmtContext) AllRSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserRSQUAREBRACKET)
}

func (s *AssignStmtContext) RSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserRSQUAREBRACKET, i)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) AssignStmt() (localctx IAssignStmtContext) {
	this := p
	_ = this

	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, C0ParserRULE_assignStmt)
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
		p.SetState(180)
		p.Match(C0ParserIDENTIFIER)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserLSQUAREBRACKET {
		{
			p.SetState(181)
			p.Match(C0ParserLSQUAREBRACKET)
		}
		{
			p.SetState(182)
			p.expression(0)
		}
		{
			p.SetState(183)
			p.Match(C0ParserRSQUAREBRACKET)
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(C0ParserASSIGN)
	}
	{
		p.SetState(191)
		p.expression(0)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(C0ParserIF, 0)
}

func (s *IfStmtContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *IfStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(C0ParserELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, C0ParserRULE_ifStmt)

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
		p.SetState(193)
		p.Match(C0ParserIF)
	}
	{
		p.SetState(194)
		p.Match(C0ParserLPARENTHESIS)
	}
	{
		p.SetState(195)
		p.expression(0)
	}
	{
		p.SetState(196)
		p.Match(C0ParserRPARENTHESIS)
	}
	{
		p.SetState(197)
		p.Stmt()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(198)
			p.Match(C0ParserELSE)
		}
		{
			p.SetState(199)
			p.Stmt()
		}

	}

	return localctx
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_whileStmt
	return p
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(C0ParserWHILE, 0)
}

func (s *WhileStmtContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *WhileStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStmtContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *WhileStmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) WhileStmt() (localctx IWhileStmtContext) {
	this := p
	_ = this

	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, C0ParserRULE_whileStmt)

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
		p.SetState(202)
		p.Match(C0ParserWHILE)
	}
	{
		p.SetState(203)
		p.Match(C0ParserLPARENTHESIS)
	}
	{
		p.SetState(204)
		p.expression(0)
	}
	{
		p.SetState(205)
		p.Match(C0ParserRPARENTHESIS)
	}
	{
		p.SetState(206)
		p.Stmt()
	}

	return localctx
}

// IFuncCallStmtContext is an interface to support dynamic dispatch.
type IFuncCallStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallStmtContext differentiates from other interfaces.
	IsFuncCallStmtContext()
}

type FuncCallStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallStmtContext() *FuncCallStmtContext {
	var p = new(FuncCallStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_funcCallStmt
	return p
}

func (*FuncCallStmtContext) IsFuncCallStmtContext() {}

func NewFuncCallStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallStmtContext {
	var p = new(FuncCallStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_funcCallStmt

	return p
}

func (s *FuncCallStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *FuncCallStmtContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *FuncCallStmtContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *FuncCallStmtContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FuncCallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterFuncCallStmt(s)
	}
}

func (s *FuncCallStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitFuncCallStmt(s)
	}
}

func (s *FuncCallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitFuncCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) FuncCallStmt() (localctx IFuncCallStmtContext) {
	this := p
	_ = this

	localctx = NewFuncCallStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, C0ParserRULE_funcCallStmt)
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
		p.SetState(208)
		p.Match(C0ParserIDENTIFIER)
	}
	{
		p.SetState(209)
		p.Match(C0ParserLPARENTHESIS)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(C0ParserDECIMAL-16))|(1<<(C0ParserOCTAL-16))|(1<<(C0ParserHEXADECIMAL-16))|(1<<(C0ParserIDENTIFIER-16))|(1<<(C0ParserLBRACE-16))|(1<<(C0ParserLPARENTHESIS-16))|(1<<(C0ParserPLUS-16))|(1<<(C0ParserMINUS-16))|(1<<(C0ParserMULPLUS-16))|(1<<(C0ParserMULMINUS-16))|(1<<(C0ParserNOT-16)))) != 0 {
		{
			p.SetState(210)
			p.ExpressionList()
		}

	}
	{
		p.SetState(213)
		p.Match(C0ParserRPARENTHESIS)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(C0ParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, C0ParserRULE_expressionList)
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
		p.SetState(215)
		p.expression(0)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == C0ParserCOMMA {
		{
			p.SetState(216)
			p.Match(C0ParserCOMMA)
		}
		{
			p.SetState(217)
			p.expression(0)
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// GetPostfix returns the postfix token.
	GetPostfix() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// SetPostfix sets the postfix token.
	SetPostfix(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) GetPostfix() antlr.Token { return s.postfix }

func (s *ExpressionContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *ExpressionContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(C0ParserLBRACE, 0)
}

func (s *ExpressionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(C0ParserRBRACE, 0)
}

func (s *ExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(C0ParserCOMMA)
}

func (s *ExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserCOMMA, i)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(C0ParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(C0ParserMINUS, 0)
}

func (s *ExpressionContext) MULPLUS() antlr.TerminalNode {
	return s.GetToken(C0ParserMULPLUS, 0)
}

func (s *ExpressionContext) MULMINUS() antlr.TerminalNode {
	return s.GetToken(C0ParserMULMINUS, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(C0ParserNOT, 0)
}

func (s *ExpressionContext) FuncCallStmt() IFuncCallStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallStmtContext)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(C0ParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(C0ParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(C0ParserMOD, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(C0ParserLE, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(C0ParserGE, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(C0ParserGT, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(C0ParserLT, 0)
}

func (s *ExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(C0ParserEQ, 0)
}

func (s *ExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(C0ParserNE, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(C0ParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(C0ParserOR, 0)
}

func (s *ExpressionContext) AllLSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserLSQUAREBRACKET)
}

func (s *ExpressionContext) LSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserLSQUAREBRACKET, i)
}

func (s *ExpressionContext) AllRSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(C0ParserRSQUAREBRACKET)
}

func (s *ExpressionContext) RSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(C0ParserRSQUAREBRACKET, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *C0Parser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, C0ParserRULE_expression, _p)
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
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(224)
			p.Match(C0ParserLPARENTHESIS)
		}
		{
			p.SetState(225)
			p.expression(0)
		}
		{
			p.SetState(226)
			p.Match(C0ParserRPARENTHESIS)
		}

	case 2:
		{
			p.SetState(228)
			p.Primary()
		}

	case 3:
		{
			p.SetState(229)
			p.Match(C0ParserLBRACE)
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(C0ParserDECIMAL-16))|(1<<(C0ParserOCTAL-16))|(1<<(C0ParserHEXADECIMAL-16))|(1<<(C0ParserIDENTIFIER-16))|(1<<(C0ParserLBRACE-16))|(1<<(C0ParserLPARENTHESIS-16))|(1<<(C0ParserPLUS-16))|(1<<(C0ParserMINUS-16))|(1<<(C0ParserMULPLUS-16))|(1<<(C0ParserMULMINUS-16))|(1<<(C0ParserNOT-16)))) != 0 {
			{
				p.SetState(230)
				p.expression(0)
			}
			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == C0ParserCOMMA {
				{
					p.SetState(231)
					p.Match(C0ParserCOMMA)
				}
				{
					p.SetState(232)
					p.expression(0)
				}

				p.SetState(237)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(240)
			p.Match(C0ParserRBRACE)
		}

	case 4:
		{
			p.SetState(241)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(C0ParserPLUS-31))|(1<<(C0ParserMINUS-31))|(1<<(C0ParserMULPLUS-31))|(1<<(C0ParserMULMINUS-31)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(242)
			p.expression(9)
		}

	case 5:
		{
			p.SetState(243)

			var _m = p.Match(C0ParserNOT)

			localctx.(*ExpressionContext).prefix = _m
		}
		{
			p.SetState(244)
			p.expression(8)
		}

	case 6:
		{
			p.SetState(245)
			p.FuncCallStmt()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(249)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(C0ParserMUL-33))|(1<<(C0ParserDIV-33))|(1<<(C0ParserMOD-33)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(250)
					p.expression(8)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(252)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == C0ParserPLUS || _la == C0ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(253)
					p.expression(7)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(255)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(C0ParserLT-40))|(1<<(C0ParserGT-40))|(1<<(C0ParserLE-40))|(1<<(C0ParserGE-40)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(256)
					p.expression(6)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(258)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == C0ParserEQ || _la == C0ParserNE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(259)
					p.expression(5)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(261)

					var _m = p.Match(C0ParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(262)
					p.expression(4)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(264)

					var _m = p.Match(C0ParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(265)
					p.expression(3)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(267)
							p.Match(C0ParserLSQUAREBRACKET)
						}
						{
							p.SetState(268)
							p.expression(0)
						}
						{
							p.SetState(269)
							p.Match(C0ParserRSQUAREBRACKET)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(273)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, C0ParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(276)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == C0ParserMULPLUS || _la == C0ParserMULMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserLPARENTHESIS, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(C0ParserRPARENTHESIS, 0)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(C0ParserIDENTIFIER, 0)
}

func (s *PrimaryContext) Intliteral() IIntliteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntliteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntliteralContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, C0ParserRULE_primary)

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

	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case C0ParserLPARENTHESIS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Match(C0ParserLPARENTHESIS)
		}
		{
			p.SetState(283)
			p.expression(0)
		}
		{
			p.SetState(284)
			p.Match(C0ParserRPARENTHESIS)
		}

	case C0ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(C0ParserIDENTIFIER)
		}

	case C0ParserDECIMAL, C0ParserOCTAL, C0ParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.Intliteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntliteralContext is an interface to support dynamic dispatch.
type IIntliteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntliteralContext differentiates from other interfaces.
	IsIntliteralContext()
}

type IntliteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntliteralContext() *IntliteralContext {
	var p = new(IntliteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = C0ParserRULE_intliteral
	return p
}

func (*IntliteralContext) IsIntliteralContext() {}

func NewIntliteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntliteralContext {
	var p = new(IntliteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = C0ParserRULE_intliteral

	return p
}

func (s *IntliteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntliteralContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(C0ParserDECIMAL, 0)
}

func (s *IntliteralContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(C0ParserOCTAL, 0)
}

func (s *IntliteralContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(C0ParserHEXADECIMAL, 0)
}

func (s *IntliteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntliteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntliteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.EnterIntliteral(s)
	}
}

func (s *IntliteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C0Listener); ok {
		listenerT.ExitIntliteral(s)
	}
}

func (s *IntliteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case C0Visitor:
		return t.VisitIntliteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *C0Parser) Intliteral() (localctx IIntliteralContext) {
	this := p
	_ = this

	localctx = NewIntliteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, C0ParserRULE_intliteral)
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
		p.SetState(290)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<C0ParserDECIMAL)|(1<<C0ParserOCTAL)|(1<<C0ParserHEXADECIMAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *C0Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *C0Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
