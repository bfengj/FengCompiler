lexer grammar Lex;

//关键字
INT: 'int';
CHAR: 'char';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
TRUE: 'true';
FALSE: 'false';
VOID: 'void';
RETURN: 'return';
CONST: 'const';
//PUTINT: 'putint'; //只支持输出int型数据
BREAK: 'break';
CONTINUE: 'continue';

//字面量和标识符
DECIMAL : [1-9][0-9]* ;//10进制
OCTAL
    : '0'
    | '0'[0-7]*//8进制
    ;
HEXADECIMAL : HEXADECIMALPREFIX [0-9a-fA-F]* ;//16进制
HEXADECIMALPREFIX : '0x' | '0X' ;

IDENTIFIER: ('a'..'z'|'A'..'Z'|'_') ('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;


//符号
COMMA: ',';
SEMICOLON: ';';
COLON: ':';
LBRACE: '{';
RBRACE: '}';
LPARENTHESIS: '(';
RPARENTHESIS: ')';
LSQUAREBRACKET: '[';
RSQUAREBRACKET: ']';

ASSIGN: '=';
PLUS: '+';
MINUS: '-';
MUL: '*';
DIV: '/';
MOD: '%';
MULPLUS: '++';
MULMINUS: '--';

EQ: '==';
NE: '!=';
LT: '<';
GT: '>';
LE: '<=';
GE: '>=';
NOT: '!';
AND: '&&';
OR: '||';




//注释
LINECOMMENT: '//';
BEGINCOMMENT: '/*';
ENDCOMMENT: '*/';
