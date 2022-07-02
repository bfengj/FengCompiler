# 编译器报告

## 整体实现

整体使用Golang语言设计，前端使用antlr生成AST，然后只遍历一遍AST，在遍历的过程中一边进行语义分析，一遍生成LLVM IR。AST遍历结束后利用LLVM PASS对生成的LLVM IR进行优化，利用LLVM执行LLVM IR生成结果。

已实现：

- 变量定义（int，array）
- 变量赋值（int，array）
- 全局变量
- const
- if
- while
- 函数定义
- 函数调用
- SySY库函数调用（输入输出函数）





编译器比赛2021年的144个测试用例全部正确通过（functional用例和h_functional用例）。

## 前端

### 符号表设计

```go
type Symbol struct {
	// 变量类型因为只有int和int[],就不设置了。
	//如果不是数组，Size设-1。
	//size不是-1就认为是数组。
	Name        string      // 变量名
	NamespaceId int         // 作用域，为0即为全局
	Size        []int       // 数组size 
	IsConst     bool        //是否是常量const
	Value       interface{} //
}

type FuncSymbol struct {
	Name       string      // 函数名
	ReturnType int         // 函数返回类型 int  or void,没有数组。0为void,1为int
	ArgsType   []FuncParam // 函数参数 string为参数名
}
type FuncParam struct {
	ParamName string
	ParamType llvm.Type
}
```

```go
type Frame struct {
	Length      int
	Symbols     []*Symbol //
	NamespaceId int
}
type ProgStack struct {
	Length int
	//BasicBlocks []llvm.BasicBlock
	Funcs  []*FuncSymbol //并不是栈
	Frames []*Frame      //每进入一个block，在栈顶新加入一层，退出block的时候弹出这一层；
}
```

有变量Symbol和函数Symbol，变量Symbol记录了变量的变量名、作用域、数组的size（如果是数组）、是否是常量、以及在代码中的Value。函数Symbol记录函数名、返回类型还有形参的类型列表。

对于作用域的处理使用栈的形式，每进入了一个block就创建一个作用域的栈帧并将其入栈。每离开一个block的时候就弹出，每次寻找变量的时候自顶向下遍历栈中作用域的帧即可。

### AST

调用antlr工具生成AST。

具体文法：

```c0
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

```

```g4
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
    |IDENTIFIER '[' ']' ('[' expression ']')*
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
```

遍历AST则是继承工具生成的`BaseC0Visitor`，Visitor相比listenner虽然没那么方便，但是更加自由和高可控度，对于AST的遍历流程和代码都可自己操控。然后依次重写`VisitProg`、`VisitStmtList`等一系列方法，实现对AST的遍历。

### 语义分析

1.  同一作用域下变量重复定义             
2.  使用了未定义变量                     
3.  使用了未定义函数                                             
4.  数组定义时中括号里的值不能为负       
5.  函数重复定义                        
6.  函数形参重复定义                    
7.  break和continue只能用在while里      
8.  const常量不能被修改                 
9.  函数形参中数组的长度检查：（当 FuncParam 为数组定义时，其第一维的长度省去（用方括号[]表示），而后面的各维则需要用表达式指明长度，长度是常量。）

### LLVM IR生成

#### 变量定义

对于int，如果不是全局变量，直接alloca即可。如果是全局变量，需要`AddGlobal`并设置初始值0即可。

对于array，如果不是全局变量，从最低维开始依次遍历维度数，依据此来产生具体的ArrayType，遍历完之后alloca。如果是全局变量，再遍历第二遍对每个元素赋0值即可。

如果是在定义的时候赋值，则不做处理在赋值中处理。

#### 变量赋值

对于int，如果是在定义的时候赋值例如`int a = 1;`，则类似int的定义，只是再visit一下后面的expression获得赋的值然后load即可。如果是在定义之后赋的值，需要先根据变量名查找符号表，然后对Symbol的Value进行Store。

对于数组，如果是定义的时候给出类似`int a[2][3] = {1,2,3};`的赋值，则从最低维遍历，即生成类型，也是记录这个数组的每一个维的维度数。然后处理Expression，`{1,2,3}`的是根据数组的维度数进行填充0，最后得到一个`[]llvm.Value`，然后就是对数组的每个元素进行赋值，依次根据索引取指针拿到目标元素的指针，然后根据索引取的他在`[]llvm.Value`中相应的那个元素，`Store`即可。

如果是已经定义过了，类似`a[0][0] = 5;`这样的赋值，原理类似，仍然是依次取到对应索引的指针然后赋值。

但是具体的实现细节比较复杂，关于各种指针，以及`{1,2,3,{2]}}`这样的转换逻辑。

#### const

例如：

```c
const int a = 5;
return a;
```

这样的，实际产生的IR中并没有对变量5赋值的逻辑，只有一个`return 5`，因此对于const变量每次运算的时候没有`load`操作，直接拿`Symbol.Value`，调用其`ZExtValue`获得具体的Value然后操作。

如果是const的数组，发现实际的LLVM在真正生成LLVM IR的时候也并没有这样的处理，因此就把const数组当成了正常的数组来处理。

#### if

生成3个block，一个是blockThen即根据cond为true将会进入的block，一个是blockElse即Else进入的block（如果有else），另外一个是blockMerge即if结束后要进入的block。

首先访问if中的condition，使用icmp来根据cond决定跳转到哪个block，然后依次进入相应的block，在对应的block里面产生LLVM IR。最后都Br到blockMerge里面。

#### while

产生3个block，condBB、bodyBB和endBB。具体的跟if类似。不过需要额外考虑break和continue。也都是借助栈这个结构来处理

break的话就是每进入一个while的body就把当前while的endBB入栈，每遇到一个break就将end的栈顶弹出来获取这个endBB然后跳转到它那。

continue的话是每进入一个while的cond就把当前while的condBB入栈，每遇到一个continue就取的cond栈的栈顶的block然后跳转到那个block。在end的时候将栈顶弹出即可。

#### 函数定义

先获取函数的函数名、返回值、形式参数这些信息来构造一个Function，同时把形势参数的Symbol放到函数body的作用域里面。然后访问函数体即可。

函数的形参中的数组也需要额外考虑，函数形参传递的实际上是数组的起始地址，在LLVM IR中则是需要在迭代ArrayType的基础上再加一层PointerType即可。

#### 函数调用

更多的是Expression方面的处理，对于`add(a,b)`这样的调用，首先处理`a,b`这个expressionList，最后返回一个`[]llvm.Value`，然后产生`Call`即可。

#### 库函数调用

库函数调用实际上就是对printf和scanf的调用，LLVM中想要调用C语言的库函数，只需要声明一个和C语言库函数一样的函数然后去调用即可。

比较麻烦的则是`putarray`和`getarray`，因为他们的实现：

```c
void putarray(int n,int a[]){
  printf("%d:",n);
  for(int i=0;i<n;i++)printf(" %d",a[i]);
  printf("\n");
}
int getarray(int a[]){
  int n;
  scanf("%d",&n);
  for(int i=0;i<n;i++)scanf("%d",&a[i]);
  return n;
}
```

不仅仅调用了scanf而且还需要for循环，因此类似之前while的实现，再实现一下c语言的for循环即可。

#### 其他

##### return之后的IR优化

例如这样的：

```c
if(a=1){
   return 1; 
   int b;
}
```

实际上在return之后就不应该再继续访问之后的stmt来生成IR，但因为antlr的Visit并不能直接实现这样的控制，因此需要做额外的处理。

在遇到return之后把当前的block作为`hasReturnBlock`，之后访问stmt的时候如果当前的block和`hasReturnBlock`是同一个则不再遍历此stmt直到跳出了这个block





##### 逻辑表达式

此外逻辑运算也是非常的复杂，即`&&`和`||`。对于`&&`来说需要全部的条件都为真才能进入blockThen，否则就要进入别的block，这实际上跟if和while的实现也有限类似，因此在if和while的cond里面，还需要做类似的工作。

这里以`x&&y`举例，需要先取的x然后让x与0做比较，如果不是0才能进入y，否则则跳出，而x和y则仍然可能是逻辑expression，依次嵌套，因此每进入一个`x&&y`都要把y这个block入栈，栈的最底部是这个cond将要跳入的block，如果x为真则弹出栈顶的block并跳入，以此嵌套，最后成功进入cond将要跳入或者跳出的block。





##### main函数return值的截断

因为一个程序的返回值实际上是一个0-255的信号值，不同的值代表了不同的运行结果（可能不同的值代表着不同的错误），因此需要截断后8位并转成uint8。

## 后端

### 优化

使用LLVM PASS对LLVM IR进行优化。因为Go中的LLVM没法自定义LLVM PASS，因此我们的优化有2种可选的方式，第一种是直接在Go里面调用LLVM库已经提供好了的各种PASS来进行优化，另外一种则是我们自己拿C++简单实现了死代码删除和函数内联的优化。



函数内联的逻辑是如果一个函数没有while而且语句数不超过10则将其设计为内联函数。删除的死代码则是删除非下面要求的指令：

1. 该指令是调试相关的指令。
2. 该指令是伪代码指令。
3. 该指令作为基本块的出口，比如：br、ret等指令。
4. 该指令是函数调用，并且该指令可能有副作用（即可能写入内存或者抛出异常）。
5. 以及与上述四种指令之间存在Use边的指令。



然后直接在Go里面让LLVM直接优化后的IR来产生结果或者调用lli命令来执行优化后的IR。







## 遇到的困难

因为正好在学Go就拿Go写了Compiler，结果发现go里面用antlr和llvm的几乎没有，这就导致了找不到相关的教程或者文章来供学习和参考，官方文档也都没有go的东西，因此全程只能一点一点摸索，学习LLVM以及它的IR和PASS。对于每个实现都先用LLVM的命令来生成一下IR，把IR给看懂之后理解LLVM生成IR的逻辑然后仿照着在Go里面不断摸索的来产生IR，虽然非常的费时费力但是这种探索也是特别有意思的，这也是第一次我有了一种”我做的东西是之前没有人做过的“的感觉。最后也是成功的解决了各种问题完成了编译器。









