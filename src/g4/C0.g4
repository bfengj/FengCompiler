grammar C0;

import Lex;

prog: stmtList;

stmtList: stmt*;

//语句
stmt
    : varDeclStmt
    | constDeclStmt
    | assignStmt ';'
    | funcDeclStmt
    | funcCallStmt ';'
    | block
    | ifStmt
    | whileStmt
    //| PUTINT '(' expression ')' ';'
    | RETURN expression? ';'
    | BREAK  ';'
    | CONTINUE ';'
    | ';'
    ;

block: '{' stmt* '}';

//函数声明
funcDeclStmt: typeTypeOrVoid IDENTIFIER formalParameters funcBody;

typeTypeOrVoid: typeType | VOID;
typeType: baseType ('[' ']')*;
baseType: INT;

formalParameters:  '(' paraList? ')';    //合法的参数声明
paraList: parameter (',' parameter)*;    //参数列表
parameter: baseType varDelcId;           //单个参数声明
varDelcId
    : IDENTIFIER            //参数ID
    | IDENTIFIER '[' ']' ('[' expression ']')*
    ;

funcBody: block | ';';

//变量声明
//varDeclStmt
//    : baseType IDENTIFIER (',' IDENTIFIER)* ';'
//    | baseType assignStmt ( ',' assignStmt)* ';'
//    | baseType IDENTIFIER ('[' expression ']')+ ';' //数组
//    ;
varDeclStmt
    : baseType varDef (',' varDef)* ';'
    ;
varDef
    : IDENTIFIER ('[' expression ']')*
    | assignStmt
    ;
constDeclStmt : CONST baseType assignStmt (',' assignStmt)* ';' ;
//赋值语句
assignStmt: IDENTIFIER  ('[' expression ']')*'=' expression ;

//if语句
ifStmt: IF '(' expression ')' stmt (ELSE stmt)?;

//for语句
whileStmt: WHILE '(' expression ')' stmt;

//函数调用语句
funcCallStmt: IDENTIFIER '(' expressionList? ')' ;

expressionList: expression (',' expression)*;

expression
    : '(' expression ')'|primary  | expression ('[' expression ']')+
    | '{' (expression (',' expression)*)? '}'
    //按运算符优先级由高到低
    | expression postfix=('++' | '--')
    | prefix=('+' | '-' | '++' | '--') expression
    | prefix='!' expression
    | expression bop=('*' | '/' | '%') expression
    | expression bop=('+' | '-') expression
    | expression bop=('<=' | '>=' | '>' | '<') expression
    | expression bop=('==' | '!=') expression
    | expression bop='&&' expression
    | expression bop='||' expression
    | funcCallStmt
    ;

primary
    : '(' expression ')'
    | IDENTIFIER
    | intliteral
    ;

intliteral
    : DECIMAL
    | OCTAL
    | HEXADECIMAL
    ;
//注释
COMMENT: (LINECOMMENT .*? NEWLINE | BEGINCOMMENT .*? ENDCOMMENT) -> skip;
//空格
WS: (' '|'\t')+ -> skip;
//空行
NEWLINE: '\r'? '\n' -> skip;