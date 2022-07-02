package visitor

/*
#include<stdio.h>
int getint(){
    int t; scanf("%d",&t);
    return t;
}
int getch(){
    char c; scanf("%c",&c);
    return (int)c;
}
int getarray(int a[]){
  int n;
  scanf("%d",&n);
  for(int i=0;i<n;i++)scanf("%d",&a[i]);
  return n;
}
void putint(int a){
    printf("%d",a);
}
void putch(int a){
    printf("%c",a);
}
void putarray(int n,int a[]){
  printf("%d:",n);
  for(int i=0;i<n;i++)printf(" %d",a[i]);
  printf("\n");
}
*/
import "C"
import (
	"compiler/llvmIr"
	parser "compiler/parser/g4"
	"compiler/symbolTable"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llvm-mirror/llvm/bindings/go/llvm"
	"strconv"
	"strings"
)

type C0Visitor struct {
	parser.BaseC0Visitor
	isDecl            bool
	isEnterFunc       bool
	isConst           bool
	isCallLibraryFunc bool
	hasReturnBlock    llvm.BasicBlock
	nowNumSign        []int //0为+，1为-，因为默认为+，所以就0。改成了栈来实现
	nowFuncStmtCount  int
	endWhileBBStack   []llvm.BasicBlock
	condWhileBBStack  []llvm.BasicBlock
	willInBlockStack  []llvm.BasicBlock //用于辅助逻辑运算
	willOutBlockStack []llvm.BasicBlock
	nowArraySize      []int
	nowArrayDeep      int
	ErrorInfo         string //记录语义分析过程中的错误信息
}

func NewC0Visitor() *C0Visitor {
	return &C0Visitor{}
}

func (v *C0Visitor) popEndBBStack() llvm.BasicBlock {
	if len(v.endWhileBBStack) < 1 {
		panic("endBBStack is empty unable to pop")
	}
	top := v.endWhileBBStack[len(v.endWhileBBStack)-1]
	v.endWhileBBStack = v.endWhileBBStack[0 : len(v.endWhileBBStack)-1]
	return top
}
func (v *C0Visitor) topEndBBStack() llvm.BasicBlock {
	if len(v.endWhileBBStack) < 1 {
		panic("endBBStack is empty unable to top")
	}
	top := v.endWhileBBStack[len(v.endWhileBBStack)-1]
	return top
}
func (v *C0Visitor) pushEndBBStack(bb llvm.BasicBlock) {
	v.endWhileBBStack = append(v.endWhileBBStack, bb)
}
func (v *C0Visitor) popCondWhileBBStack() llvm.BasicBlock {
	if len(v.condWhileBBStack) < 1 {
		panic("condWhileBBStack is empty unable to pop")
	}
	top := v.condWhileBBStack[len(v.condWhileBBStack)-1]
	v.condWhileBBStack = v.condWhileBBStack[0 : len(v.condWhileBBStack)-1]
	return top
}
func (v *C0Visitor) topCondWhileBBStack() llvm.BasicBlock {
	if len(v.condWhileBBStack) < 1 {
		panic("condWhileBBStack is empty unable to top")
	}
	top := v.condWhileBBStack[len(v.condWhileBBStack)-1]
	return top
}
func (v *C0Visitor) pushCondWhileBBStack(bb llvm.BasicBlock) {
	v.condWhileBBStack = append(v.condWhileBBStack, bb)
}

func (v *C0Visitor) popWillInBlockStack() llvm.BasicBlock {
	if len(v.willInBlockStack) < 1 {
		panic("willInBlockStack is empty unable to pop")
	}
	top := v.willInBlockStack[len(v.willInBlockStack)-1]
	v.willInBlockStack = v.willInBlockStack[0 : len(v.willInBlockStack)-1]
	return top
}
func (v *C0Visitor) topWillInBlockStack() llvm.BasicBlock {
	if len(v.willInBlockStack) < 1 {
		panic("willInBlockStack is empty unable to top")
	}
	top := v.willInBlockStack[len(v.willInBlockStack)-1]
	return top
}
func (v *C0Visitor) pushWillInBlockStack(bb llvm.BasicBlock) {
	v.willInBlockStack = append(v.willInBlockStack, bb)
}
func (v *C0Visitor) popWillOutBlockStack() llvm.BasicBlock {
	if len(v.willOutBlockStack) < 1 {
		panic("willOutBlockStack is empty unable to pop")
	}
	top := v.willOutBlockStack[len(v.willOutBlockStack)-1]
	v.willOutBlockStack = v.willOutBlockStack[0 : len(v.willOutBlockStack)-1]
	return top
}
func (v *C0Visitor) topWillOutBlockStack() llvm.BasicBlock {
	if len(v.willOutBlockStack) < 1 {
		panic("willOutBlockStack is empty unable to top")
	}
	top := v.willOutBlockStack[len(v.willOutBlockStack)-1]
	return top
}
func (v *C0Visitor) pushWillOutBlockStack(bb llvm.BasicBlock) {
	v.willOutBlockStack = append(v.willOutBlockStack, bb)
}

func (v *C0Visitor) visitRule(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *C0Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *C0Visitor) VisitProg(ctx *parser.ProgContext) interface{} {

	v.ErrorInfo = "" //========
	llvmIr.PrintfFunc = llvm.AddFunction(llvmIr.GlobalMod, "printf", llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, true))
	llvmIr.ScanfFunc = llvm.AddFunction(llvmIr.GlobalMod, "scanf", llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, true))
	//
	//全局frame
	frame := symbolTable.NewFrame(0, nil, symbolTable.NextNamespaceId)
	symbolTable.GlobalSymbolTable.PushFrame(frame)
	v.visitRule(ctx.StmtList())

	//debug.PrintFrames()
	symbolTable.GlobalSymbolTable.PopFrame()

	return nil
	//return nil
}

func (v *C0Visitor) VisitStmtList(ctx *parser.StmtListContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.visitRule(stmt)
	}
	return nil
}

func (v *C0Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if symbolTable.NowBasicBlock[0] == v.hasReturnBlock && !symbolTable.NowBasicBlock[0].IsNil() {
		return nil
	}
	v.nowFuncStmtCount++
	//antlr有没有一种比较简洁的写法。。。。。
	if ctx.VarDeclStmt() != nil {
		v.visitRule(ctx.VarDeclStmt())
	} else if ctx.ConstDeclStmt() != nil {
		v.visitRule(ctx.ConstDeclStmt())
	} else if ctx.AssignStmt() != nil {
		v.visitRule(ctx.AssignStmt())
	} else if ctx.FuncDeclStmt() != nil {
		v.visitRule(ctx.FuncDeclStmt())
	} else if ctx.FuncCallStmt() != nil {
		v.visitRule(ctx.FuncCallStmt())
	} else if ctx.Block() != nil {
		v.visitRule(ctx.Block())
	} else if ctx.IfStmt() != nil {
		v.visitRule(ctx.IfStmt())
	} else if ctx.WhileStmt() != nil {
		funcValue := llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name)
		funcValue.AddTargetDependentFunctionAttr("isInliner", "0")
		v.visitRule(ctx.WhileStmt())
	} else if strings.Contains(ctx.GetText(), "return") {
		//return
		var result llvm.Value
		if ctx.Expression() != nil {
			//类型方面会出问题吗？？？
			result = v.visitRule(ctx.Expression()).(llvm.Value) //类型转换
		} else {
			result = llvm.Value{}
		}
		//不需要在这里进行截断，卡了好久才发现原来编译器的返回值并不是在这里产生的而是在main.go里面。
		llvmIr.GlobalBuilder.CreateRet(result)
		v.hasReturnBlock = symbolTable.NowBasicBlock[0]

	} else if strings.Contains(ctx.GetText(), "break") {
		//break

		if len(v.endWhileBBStack) < 1 { //break只能出现在while语句里
			info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " break only used in While Block\n"
			v.ErrorInfo += info
		} else {
			endBB := v.topEndBBStack()
			llvmIr.GlobalBuilder.CreateBr(endBB)
		}
	} else if strings.Contains(ctx.GetText(), "continue") {

		if len(v.endWhileBBStack) < 1 { //continue只能出现在while语句里
			info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " break only used in While Block\n"
			v.ErrorInfo += info
		} else {
			condBB := v.topCondWhileBBStack()
			llvmIr.GlobalBuilder.CreateBr(condBB)
		}
		//continue
	} /*else if strings.Contains(ctx.GetText(), "putint") {
		//putint
		//已经删掉了。
	}*/
	return nil
}

func (v *C0Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	if !v.isEnterFunc {
		frame := symbolTable.NewFrame(0, nil, symbolTable.NextNamespaceId)
		symbolTable.GlobalSymbolTable.PushFrame(frame)
		//v.isEnterFunc = false
	} else {
		v.isEnterFunc = false
	}
	stmts := ctx.AllStmt()
	for _, stmt := range stmts {
		v.visitRule(stmt)
	}
	symbolTable.GlobalSymbolTable.PopFrame()
	/*	if !symbolTable.IsDelayPopFrame {
		debug.PrintFrames()
		symbolTable.GlobalSymbolTable.PopFrame()
	}*/
	return nil
}

func (v *C0Visitor) VisitFuncDeclStmt(ctx *parser.FuncDeclStmtContext) interface{} {
	//生成符号表及语义分析

	//=======检查函数是否重复定义=======
	funcName := ctx.IDENTIFIER().GetText()
	res := symbolTable.GlobalSymbolTable.GetFuncSymbolByName(funcName)
	if res != nil {
		info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + funcName + " has been defined\n"
		v.ErrorInfo += info
	}
	//==============================

	//===检查函数是否定义在某个函数里====
	//curFrame := symbolTable.GlobalSymbolTable.TopFrame()
	//==============================

	funcSym := symbolTable.NewFuncSymbol(ctx.IDENTIFIER().GetText(), 1, []symbolTable.FuncParam{})
	symbolTable.GlobalSymbolTable.AddFunc(funcSym)

	v.isEnterFunc = true
	frame := symbolTable.NewFrame(0, nil, symbolTable.NextNamespaceId)
	symbolTable.GlobalSymbolTable.PushFrame(frame)
	v.nowFuncStmtCount = 0
	v.visitRule(ctx.TypeTypeOrVoid())
	v.visitRule(ctx.FormalParameters())
	v.CreateIrFuncDeclStmt()
	v.AddFormalParamToSymbolTable()
	v.visitRule(ctx.FuncBody())

	if funcSym.ReturnType == 0 {
		llvmIr.GlobalBuilder.CreateRetVoid()
	}
	//大于10条语句的函数就不内联了。
	if v.nowFuncStmtCount > 10 {
		funcValue := llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name)
		funcValue.AddTargetDependentFunctionAttr("isInliner", "0")
	}
	v.nowFuncStmtCount = 0
	symbolTable.NowBasicBlock = make([]llvm.BasicBlock, 1)
	return nil
}
func (v *C0Visitor) AddFormalParamToSymbolTable() interface{} {
	//params := symbolTable.GlobalSymbolTable.GetFuncSymbol().ArgsType
	funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbol()
	params := llvmIr.GlobalMod.NamedFunction(funcSymbol.Name).Params()
	var count int = 0
	for _, param := range funcSymbol.ArgsType {
		allocaTemp := llvmIr.GlobalBuilder.CreateAlloca(param.ParamType, "")
		//顺序有问题，因为ArgsType 即Go的Map是无序的。
		llvmIr.GlobalBuilder.CreateStore(params[count], allocaTemp)

		symbol := symbolTable.NewSymbol(param.ParamName, symbolTable.NowNamespaceId, nil, false, allocaTemp)
		//symbol := symbolTable.NewVarSymbol(name, symbolTable.NowNamespaceId, -1, false, params[count])
		symbolTable.GlobalSymbolTable.AddSymbol(symbol)
		count++
	}
	return nil
}
func (v *C0Visitor) CreateIrFuncDeclStmt() interface{} {
	funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbol()
	var funcName string
	var returnType llvm.Type
	var paramTypes []llvm.Type
	funcName = funcSymbol.Name
	if funcSymbol.ReturnType == 1 {
		returnType = llvm.Int32Type()
	} else {
		returnType = llvm.VoidType()
	}
	for _, param := range funcSymbol.ArgsType {

		paramTypes = append(paramTypes, param.ParamType)
	}
	f := llvm.FunctionType(returnType, paramTypes, false)
	llvm.AddFunction(llvmIr.GlobalMod, funcName, f)

	block := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(funcName), "entry") //入口block
	llvmIr.GlobalBuilder.SetInsertPoint(block, block.FirstInstruction())
	symbolTable.NowBasicBlock[0] = block
	return nil
}

func (v *C0Visitor) VisitTypeTypeOrVoid(ctx *parser.TypeTypeOrVoidContext) interface{} {
	if ctx.VOID() != nil {
		//void
		symbolTable.GlobalSymbolTable.SetFuncReturnType()
	}
	return nil
}

func (v *C0Visitor) VisitTypeType(ctx *parser.TypeTypeContext) interface{} {
	return nil
}

func (v *C0Visitor) VisitBaseType(ctx *parser.BaseTypeContext) interface{} {
	return nil
}

func (v *C0Visitor) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	//2个任务，FuncSymbol的ArgsType和当前frame的Symbol
	//symbolTable.NowFuncArgsType = []string{}
	if ctx.ParaList() != nil {
		v.visitRule(ctx.ParaList())
	}
	return nil
}

func (v *C0Visitor) VisitParaList(ctx *parser.ParaListContext) interface{} {
	parameters := ctx.AllParameter()
	for _, parameter := range parameters {
		v.visitRule(parameter)
	}
	return nil
}

func (v *C0Visitor) VisitParameter(ctx *parser.ParameterContext) interface{} {
	v.visitRule(ctx.VarDelcId())
	return nil
}

func (v *C0Visitor) VisitVarDelcId(ctx *parser.VarDelcIdContext) interface{} {
	identifier := ctx.IDENTIFIER()

	ttype := llvm.Int32Type()
	if strings.Contains(ctx.GetText(), "[") {
		expressions := ctx.AllExpression()
		//数组
		for i := len(expressions) - 1; i >= 0; i-- {

			//======函数形参数组的长度检查=====
			//一维数组的大小可以指定也可以不指定; 高维数组只有第一维可以不指定大小，其余必须指定且为常量
			for _, s := range expressions[i].GetText() {
				//fmt.Println(s)
				if s < 48 || s > 57 { //0~9
					info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + identifier.GetText() + "'s " + strconv.Itoa(i+1) + "_dimension_size must be a int literal\n"
					v.ErrorInfo += info
					break
				}
			}
			//=============================

			length := int(v.visitRule(expressions[i]).(llvm.Value).ZExtValue())
			ttype = llvm.ArrayType(ttype, length)
		}
		ttype = llvm.PointerType(ttype, 0)
	}

	//======函数形参名不能重复========
	funcSym := symbolTable.GlobalSymbolTable.GetFuncSymbol()
	if len(funcSym.ArgsType) > 0 {
		for _, param := range funcSym.ArgsType {
			if param.ParamName == identifier.GetText() {
				v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " duplicate function param \"" + param.ParamName + "\"\n"
			}
		}
	}
	//=============================

	symbolTable.GlobalSymbolTable.AddFuncArgsType(identifier.GetText(), ttype)
	//symbol := symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, -1, false,)
	//symbolTable.GlobalSymbolTable.AddSymbol(symbol)
	//fmt.Println(identifier.GetText(), "==========")
	//真的有用吗？
	/*	params := llvmIr.GlobalMod.NamedFunction("add").Params()
		for _, i := range params {
			if identifier.GetText() == i.Name() {

				break
			}
		}*/
	return nil
}

func (v *C0Visitor) VisitConstDeclStmt(ctx *parser.ConstDeclStmtContext) interface{} {
	v.isConst = true
	v.isDecl = true
	assignStmts := ctx.AllAssignStmt()
	for _, assignStmt := range assignStmts {
		v.visitRule(assignStmt)
	}
	v.isDecl = false
	v.isConst = false
	return nil
}

func (v *C0Visitor) VisitFuncBody(ctx *parser.FuncBodyContext) interface{} {

	//暂时不考虑if、while这些复杂结构，后续可能把block的IR生成代码扔到Block中处理。
	v.visitRule(ctx.Block())
	return nil
}

func (v *C0Visitor) VisitVarDef(ctx *parser.VarDefContext) interface{} {
	if ctx.AssignStmt() != nil {
		v.isDecl = true
		v.visitRule(ctx.AssignStmt())
		v.isDecl = false
	} else if len(ctx.AllExpression()) > 0 {
		//数组，一行只会有一个
		//int a[4][5];这样的
		//const之后再考虑。
		identifier := ctx.IDENTIFIER()

		//=======同一作用域下变量不能重复定义=======
		varName := identifier.GetText()
		frame := symbolTable.GlobalSymbolTable.TopFrame()
		for i := 0; i < len(frame.Symbols); i++ {
			if varName == frame.Symbols[i].Name {
				info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + varName + " has been defined in this block\n"
				v.ErrorInfo += info
			}
		}
		//=====================================

		expressions := ctx.AllExpression()
		var size []int
		var arrayValue llvm.Value
		var arrayType llvm.Type = llvm.Int32Type()
		var count int = 1
		var typeSlice []llvm.Type
		for i := len(expressions) - 1; i >= 0; i-- {
			//从最低维开始遍历
			expression := expressions[i]
			value := v.visitRule(expression).(llvm.Value)
			dimenValue := int(value.SExtValue())

			//==数组定义时中括号里的值不能为负=
			//fmt.Println(dimenValue)
			if dimenValue <= 0 {
				info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + varName + "'s dimensions must be positive num\n"
				v.ErrorInfo += info
			}
			//===========================

			count *= dimenValue
			size = append(size, dimenValue)
			typeSlice = append(typeSlice, arrayType)
			arrayType = llvm.ArrayType(arrayType, dimenValue)
		}
		if symbolTable.NowNamespaceId == 0 {
			//全局变量
			var constVals []llvm.Value
			for i := 0; i < count; i++ {
				constVals = append(constVals, llvm.ConstInt(llvm.Int32Type(), uint64(0), false))
			}
			//再遍历一遍
			for i := 0; i <= len(expressions)-1; i++ {
				//从最低维开始遍历
				dimenValue := size[i]
				beforeType := typeSlice[i]
				var tempVals []llvm.Value
				for j := 0; j < len(constVals); j = j + dimenValue {
					tempVals = append(tempVals, llvm.ConstArray(beforeType, constVals[j:j+dimenValue]))
				}
				constVals = tempVals
			}

			arrayValue = llvm.AddGlobal(llvmIr.GlobalMod, arrayType, identifier.GetText())
			arrayValue.SetAlignment(4)
			arrayValue.SetInitializer(constVals[0])
			//globalArrayValue.SetInitializer()
		} else {
			arrayValue = llvmIr.GlobalBuilder.CreateAlloca(arrayType, "")
			arrayValue.SetAlignment(4)
		}
		symbol := symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, size, false, arrayValue)
		symbolTable.GlobalSymbolTable.AddSymbol(symbol)

	} else {
		identifier := ctx.IDENTIFIER()

		//=======同一作用域下变量不能重复定义=======
		varName := identifier.GetText()
		frame := symbolTable.GlobalSymbolTable.TopFrame()
		for i := 0; i < len(frame.Symbols); i++ {
			if varName == frame.Symbols[i].Name {
				info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + varName + " has been defined in this block\n"
				v.ErrorInfo += info
			}
		}
		//=====================================

		var value llvm.Value
		if symbolTable.NowNamespaceId == 0 {
			//全局变量，默认为0
			value = llvm.AddGlobal(llvmIr.GlobalMod, llvm.Int32Type(), identifier.GetText())
			value.SetAlignment(4)
			value.SetInitializer(llvm.ConstInt(llvm.Int32Type(), 0, false))
		} else {
			//局部变量
			//alloca
			value = llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), identifier.GetText())
			value.SetAlignment(4)
		}
		//
		symbol := symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, nil, false, value)
		symbolTable.GlobalSymbolTable.AddSymbol(symbol)
	}
	return nil
}

func (v *C0Visitor) VisitVarDeclStmt(ctx *parser.VarDeclStmtContext) interface{} {
	varDefs := ctx.AllVarDef()
	for _, varDef := range varDefs {
		v.visitRule(varDef)
	}
	return nil
}

func (v *C0Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	identifier := ctx.IDENTIFIER()
	expressions := ctx.AllExpression()
	//可能是声明中赋值，也可能是声明后赋值
	if v.isDecl == true {
		if len(expressions) > 1 {
			//数组，一行只会有一个。
			//暂时实现一维数组
			//
			var size []int
			var arrayValue llvm.Value
			var arrayType llvm.Type = llvm.Int32Type()
			var beforeType llvm.Type
			var typeSlice []llvm.Type
			//最后一个expression是 = 之后的
			for i := len(expressions) - 2; i >= 0; i-- {
				//从最低维开始遍历
				expression := expressions[i]
				value := v.visitRule(expression).(llvm.Value)
				dimenValue := int(value.ZExtValue())
				size = append(size, dimenValue)
				typeSlice = append(typeSlice, arrayType)
				arrayType = llvm.ArrayType(arrayType, dimenValue)
			}
			if symbolTable.NowNamespaceId == 0 {
				//全局变量，一维的已经实现成功了，正在弄多维的逻辑（llvm 的傻逼类型我cnm）
				v.nowArraySize = size
				////values := v.visitRule(expressions[len(expressions)-1]).([]int)
				constVals := v.visitRule(expressions[len(expressions)-1]).([]llvm.Value)
				/*				var constVals []llvm.Value
								for _, value := range values {
									constVals = append(constVals, llvm.ConstInt(llvm.Int32Type(), uint64(value), false))
								}*/
				//再遍历一遍
				for i := 0; i <= len(expressions)-2; i++ {
					//从最低维开始遍历
					dimenValue := size[i]
					beforeType = typeSlice[i]
					var tempVals []llvm.Value
					for j := 0; j < len(constVals); j = j + dimenValue {
						tempVals = append(tempVals, llvm.ConstArray(beforeType, constVals[j:j+dimenValue]))
					}
					constVals = tempVals
				}

				globalArrayValue := llvm.AddGlobal(llvmIr.GlobalMod, arrayType, identifier.GetText())
				//
				arrayType.String()
				globalArrayValue.Type().String()
				//
				globalArrayValue.SetAlignment(4)
				//v.nowArraySize = size
				//t1 := llvm.ConstInt(llvm.Int32Type(), 5, false)
				//constArray := llvm.ConstArray(beforeType, []llvm.Value{t1})
				//fmt.Println(constArray.Type().String())
				globalArrayValue.SetInitializer(constVals[0])

				symbol := symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, size, v.isConst, globalArrayValue)
				symbolTable.GlobalSymbolTable.AddSymbol(symbol)
			} else {
				arrayValue = llvmIr.GlobalBuilder.CreateAlloca(arrayType, "")
				arrayValue.SetAlignment(4)
				symbol := symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, size, v.isConst, arrayValue)
				symbolTable.GlobalSymbolTable.AddSymbol(symbol)
				//赋值
				//ptrToIntValue := llvmIr.GlobalBuilder.CreatePtrToInt(arrayValue, llvm.Int32Type(), "")
				//%1 = getelementptr inbounds [1 x i32], [1 x i32]* %0, i32 0, i32 0
				//复杂的{}的处理放在expression那边吧。
				//expression处理的时候要用到这个size。
				v.nowArraySize = size
				values := v.visitRule(expressions[len(expressions)-1]).([]llvm.Value)
				//valuesLen := len(values)
				//这里还是比较麻烦的。并不能顺序赋值，它这指针还是指向指针的，需要一直解引用。
				for i := 0; i < len(values); i++ {
					location := v.getIndexLocation(i)
					structGepValue := arrayValue
					for _, j := range location {
						//structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), uint64(j), false)}, "")
						//xxxx
						structGepValue = llvmIr.GlobalBuilder.CreateStructGEP(structGepValue, j, "")
					}
					//structGepValue := llvmIr.GlobalBuilder.CreateStructGEP(arrayValue, i, "")
					//补0 在expression的处理中弄了
					llvmIr.GlobalBuilder.CreateStore(values[i], structGepValue)
					//llvmIr.GlobalBuilder.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(values[i]), false), structGepValue)
				}
			}

		} else {
			var value llvm.Value
			/*			if expressions[0].GetText() == "getint()" {
						if symbolTable.NowNamespaceId == 0 {
							//全局变量
							//不过测试样例中并没有这样的
							value = llvm.AddGlobal(llvmIr.GlobalMod, llvm.Int32Type(), identifier.GetText())
							value.SetAlignment(4)
							v.nowScanfParam = value
						} else {
							value = llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), identifier.GetText())
							value.SetAlignment(4)
							v.nowScanfParam = value
							v.visitRule(expressions[0])
							//llvmIr.GlobalBuilder.CreateStore(num, value).SetAlignment(4)
						}
						var symbol *symbolTable.Symbol
						symbol = symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, nil, v.isConst, value)
						symbolTable.GlobalSymbolTable.AddSymbol(symbol)
					} else {*/
			num := v.visitRule(expressions[0]).(llvm.Value)
			if symbolTable.NowNamespaceId == 0 {
				//全局变量
				value = llvm.AddGlobal(llvmIr.GlobalMod, llvm.Int32Type(), identifier.GetText())
				value.SetAlignment(4)
				value.SetInitializer(num)
			} else {
				value = llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), identifier.GetText())
				value.SetAlignment(4)
				llvmIr.GlobalBuilder.CreateStore(num, value).SetAlignment(4)
			}
			var symbol *symbolTable.Symbol
			if v.isConst {
				symbol = symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, nil, v.isConst, num.ZExtValue())
			} else {
				symbol = symbolTable.NewSymbol(identifier.GetText(), symbolTable.NowNamespaceId, nil, v.isConst, value)
			}
			symbolTable.GlobalSymbolTable.AddSymbol(symbol)
			//}
		}
	} else { //声明后赋值
		//查找符号表
		if len(expressions) > 1 {
			//数组
			symbol := symbolTable.GlobalSymbolTable.FindVarSymbol(identifier.GetText())
			//==========================================
			if symbol == nil { //使用了未定义变量
				v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " \"" + identifier.GetText() + "\" not defined before\n"
			} else if symbol.IsConst == true {
				v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " \"" + identifier.GetText() + "\" can't be assigned (CONST)\n"
				//==========================================
			} else {
				//a[1] = 5
				expresLen := len(expressions)
				num := v.visitRule(expressions[expresLen-1]).(llvm.Value)
				structGepValue := symbol.Value.(llvm.Value)
				funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbol()
				var isFuncParam bool
				for _, param := range funcSymbol.ArgsType {
					if param.ParamName == symbol.Name {
						structGepValue = llvmIr.GlobalBuilder.CreateLoad(structGepValue, "")
						isFuncParam = true
					}
				}
				for i := 0; i < expresLen-1; i++ {
					location := v.visitRule(expressions[i]).(llvm.Value)
					if i == 0 && isFuncParam == true {
						structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{location}, "")
					} else {
						structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), location}, "")
					}
					//location := int(v.visitRule(expressions[i]).(llvm.Value).ZExtValue())
					//structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), location}, "")
					//xxx
					//structGepValue = llvmIr.GlobalBuilder.CreateStructGEP(structGepValue, location, "")
				}
				llvmIr.GlobalBuilder.CreateStore(num, structGepValue)
				//llvmIr.GlobalBuilder.CreateStore(llvm.ConstInt(llvm.Int32Type(), num.ZExtValue(), false), structGepValue)

				//location := v.getIndexLocation(i)
				//structGepValue := arrayValue
				//for _, j := range location {
				//	structGepValue = llvmIr.GlobalBuilder.CreateStructGEP(structGepValue, j, "")
				//}
				//structGepValue := llvmIr.GlobalBuilder.CreateStructGEP(arrayValue, i, "")
				//补0 在expression的处理中弄了
				//llvmIr.GlobalBuilder.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(values[i]), false), structGepValue)

				//llvmIr.GlobalBuilder.CreateStore(num, symbol.Value.(llvm.Value)).SetAlignment(4)
			}
		} else {
			symbol := symbolTable.GlobalSymbolTable.FindVarSymbol(identifier.GetText())
			//==========================================
			if symbol == nil { //使用了未定义变量
				v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " \"" + identifier.GetText() + "\" not defined before\n"
			} else if _, ok := symbol.Value.(uint64); ok == true { //判断变量类型是否是const(const: Value是int型; 其他为llvm.Value)
				v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " \"" + identifier.GetText() + "\" can't be assigned (CONST)\n"
				//==========================================
			} else {
				/*				if expressions[0].GetText() == "getint()" {
								value := symbol.Value.(llvm.Value)
								//value.SetAlignment(4)
								v.nowScanfParam = value
								v.visitRule(expressions[0])
							} else {*/
				num := v.visitRule(expressions[0]).(llvm.Value)
				//fmt.Println("====:", symbol.Value.(llvm.Value).Type())
				llvmIr.GlobalBuilder.CreateStore(num, symbol.Value.(llvm.Value)).SetAlignment(4)
				//}
			}
		}

	}
	//v.visitRule(ctx.Expression())

	return nil
}
func (v *C0Visitor) getIndexLocation(index int) []int {
	//a[2][3][4]
	//index 13       a[1][0][1]
	var res []int
	for i := 0; i < len(v.nowArraySize); i++ {
		count := 1
		for j := len(v.nowArraySize) - i - 2; j >= 0; j-- {
			count *= v.nowArraySize[j]
		}
		res = append(res, index/count)
		index %= count
	}
	return res
}
func (v *C0Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	//先简单的实现一下生成IR的逻辑，至于if中关于符号表、作用域这些之后再考虑。
	//if的符号表和作用域似乎也不需要额外考虑？以后遇到特例了再修改即可。
	//ifStmt: IF '(' expression ')' stmt (ELSE stmt)?;
	blockThen := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
	var blockElse llvm.BasicBlock
	blockMerge := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
	//blockMerge := llvm.PrevBasicBlock(blockThen)
	//必须得是IntegerTy(1)

	if ctx.ELSE() != nil {
		//有else
		blockElse = llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
		v.pushWillInBlockStack(blockThen)
		v.pushWillOutBlockStack(blockElse)
	} else {
		//无else
		v.pushWillInBlockStack(blockThen)
		v.pushWillOutBlockStack(blockMerge)
	}
	var condValue llvm.Value
	result := v.visitRule(ctx.Expression())
	if result == nil {
		//带&&和||的
		if ctx.ELSE() != nil {
			v.CreateThenIr(ctx, blockThen, blockMerge)
			v.CreateElseIr(ctx, blockElse, blockMerge)
		} else {
			v.CreateThenIr(ctx, blockThen, blockMerge)
		}
	} else {
		condValue = result.(llvm.Value)
		if condValue.Type() != llvm.Int1Type() {
			condValue = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, condValue, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
		}
		//创建条件跳转
		//这部分之后考虑用phi优化？
		if ctx.ELSE() != nil {
			llvmIr.GlobalBuilder.CreateCondBr(condValue, blockThen, blockElse)
			v.CreateThenIr(ctx, blockThen, blockMerge)
			v.CreateElseIr(ctx, blockElse, blockMerge)
		} else {
			llvmIr.GlobalBuilder.CreateCondBr(condValue, blockThen, blockMerge)
			v.CreateThenIr(ctx, blockThen, blockMerge)
		}
	}
	v.popWillInBlockStack()
	v.popWillOutBlockStack()
	llvmIr.GlobalBuilder.SetInsertPoint(blockMerge, blockMerge.FirstInstruction())
	symbolTable.NowBasicBlock[0] = blockMerge
	return nil
}
func (v *C0Visitor) CreateThenIr(ctx *parser.IfStmtContext, thenBB llvm.BasicBlock, brBB llvm.BasicBlock) interface{} {
	llvmIr.GlobalBuilder.SetInsertPoint(thenBB, thenBB.FirstInstruction())
	symbolTable.NowBasicBlock[0] = thenBB
	v.visitRule(ctx.Stmt(0))
	if v.hasReturnBlock != symbolTable.NowBasicBlock[0] {
		llvmIr.GlobalBuilder.CreateBr(brBB)
	}
	return nil
}
func (v *C0Visitor) CreateElseIr(ctx *parser.IfStmtContext, elseBB llvm.BasicBlock, brBB llvm.BasicBlock) interface{} {
	llvmIr.GlobalBuilder.SetInsertPoint(elseBB, elseBB.FirstInstruction())
	symbolTable.NowBasicBlock[0] = elseBB
	v.visitRule(ctx.Stmt(1))
	if v.hasReturnBlock != symbolTable.NowBasicBlock[0] {
		llvmIr.GlobalBuilder.CreateBr(brBB)
	}
	return nil
}

func (v *C0Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	//while需要创建3个block，condBB，bodyBB,endBB
	//符号表和作用域？
	condBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
	bodyBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
	endBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "end")

	v.CreateWhileCondIr(ctx, condBB, bodyBB, endBB)
	v.CreateWhileBodyIr(ctx, condBB, bodyBB, endBB)
	v.CreateWhileEndIr(endBB)
	return nil
}
func (v *C0Visitor) CreateWhileCondIr(ctx *parser.WhileStmtContext, condBB llvm.BasicBlock, bodyBB llvm.BasicBlock, endBB llvm.BasicBlock) interface{} {
	v.pushCondWhileBBStack(condBB)
	v.pushWillOutBlockStack(endBB)
	v.pushWillInBlockStack(bodyBB)
	llvmIr.GlobalBuilder.CreateBr(condBB)
	llvmIr.GlobalBuilder.SetInsertPoint(condBB, condBB.FirstInstruction())
	symbolTable.NowBasicBlock[0] = condBB
	result := v.visitRule(ctx.Expression())
	if result != nil {
		condValue := result.(llvm.Value)
		if condValue.Type() != llvm.Int1Type() {
			condValue = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, condValue, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
		}
		//由CondBB来决定是否跳转到endBB
		llvmIr.GlobalBuilder.CreateCondBr(condValue, bodyBB, endBB)
	}
	v.popWillOutBlockStack()
	v.popWillInBlockStack()
	return nil
}
func (v *C0Visitor) CreateWhileBodyIr(ctx *parser.WhileStmtContext, condBB llvm.BasicBlock, bodyBB llvm.BasicBlock, endBB llvm.BasicBlock) interface{} {
	llvmIr.GlobalBuilder.SetInsertPoint(bodyBB, bodyBB.FirstInstruction())
	symbolTable.NowBasicBlock[0] = bodyBB
	//需要考虑break
	v.pushEndBBStack(endBB)
	v.visitRule(ctx.Stmt())
	llvmIr.GlobalBuilder.CreateBr(condBB)
	return nil
}
func (v *C0Visitor) CreateWhileEndIr(endBB llvm.BasicBlock) interface{} {
	llvmIr.GlobalBuilder.SetInsertPoint(endBB, endBB.FirstInstruction())
	symbolTable.NowBasicBlock[0] = endBB
	v.popEndBBStack()
	v.popCondWhileBBStack()
	return nil
}

func (v *C0Visitor) VisitFuncCallStmt(ctx *parser.FuncCallStmtContext) interface{} {

	funcName := ctx.IDENTIFIER().GetText()
	switch funcName {
	case "getint":
		//num := uint64(C.getint())
		//return llvm.ConstInt(llvm.Int32Type(), num, false)
		if llvmIr.IntFormat == nil {
			llvmIr.IntFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%d", "intformat")
		}
		temp := llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), "")
		temp.SetAlignment(4)
		llvmIr.GlobalBuilder.CreateCall(llvmIr.ScanfFunc, []llvm.Value{llvmIr.IntFormat.(llvm.Value), temp}, "")
		return llvmIr.GlobalBuilder.CreateLoad(temp, "")
	case "getch":
		//int getch(){char c; scanf("%c",&c); return (int)c; }
		if llvmIr.ChFormat == nil {
			llvmIr.ChFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%c", "chformat")
		}
		temp := llvmIr.GlobalBuilder.CreateAlloca(llvm.Int8Type(), "")
		temp.SetAlignment(1)
		llvmIr.GlobalBuilder.CreateCall(llvmIr.ScanfFunc, []llvm.Value{llvmIr.ChFormat.(llvm.Value), temp}, "")
		ch := llvmIr.GlobalBuilder.CreateLoad(temp, "")
		ch.SetAlignment(1)
		return llvmIr.GlobalBuilder.CreateSExt(ch, llvm.Int32Type(), "")
	case "getarray":
		/*		int getarray(int a[]){
				int n;
				scanf("%d",&n);
				for(int i=0;i<n;i++)scanf("%d",&a[i]);
				return n;
			    }*/
		if llvmIr.IntFormat == nil {
			llvmIr.IntFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%d", "intformat")
		}
		//array := v.visitRule(ctx.ExpressionList()).([]llvm.Value)[0]
		temp := llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), "")
		temp.SetAlignment(4)
		llvmIr.GlobalBuilder.CreateCall(llvmIr.ScanfFunc, []llvm.Value{llvmIr.IntFormat.(llvm.Value), temp}, "")
		n := llvmIr.GlobalBuilder.CreateLoad(temp, "")

		condBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
		bodyBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
		endBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "end")

		//CondIR
		//iii := llvm.ConstInt(llvm.Int32Type(), count, false)
		iii := llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), "")
		llvmIr.GlobalBuilder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 0, false), iii)
		llvmIr.GlobalBuilder.CreateBr(condBB)
		llvmIr.GlobalBuilder.SetInsertPoint(condBB, condBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = condBB
		iiiValue := llvmIr.GlobalBuilder.CreateLoad(iii, "")
		condValue := llvmIr.GlobalBuilder.CreateICmp(llvm.IntSLT, iiiValue, n, "")
		//由CondBB来决定是否跳转到endBB
		llvmIr.GlobalBuilder.CreateCondBr(condValue, bodyBB, endBB)

		//BodyIr
		llvmIr.GlobalBuilder.SetInsertPoint(bodyBB, bodyBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = bodyBB
		v.isCallLibraryFunc = true
		trueParam := v.visitRule(ctx.ExpressionList()).([]llvm.Value)[0]
		v.isCallLibraryFunc = false
		//scanfParam := llvmIr.GlobalBuilder.CreateStructGEP(trueParam, int(iii.ZExtValue()), "")
		scanfParam := llvmIr.GlobalBuilder.CreateInBoundsGEP(trueParam, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), iiiValue}, "")
		//scanfParam := llvmIr.GlobalBuilder.CreateInBoundsGEP(trueParam, []llvm.Value{iiiValue}, "")
		llvmIr.GlobalBuilder.CreateCall(llvmIr.ScanfFunc, []llvm.Value{llvmIr.IntFormat.(llvm.Value), scanfParam}, "")
		lvalue := llvmIr.GlobalBuilder.CreateLoad(iii, "")
		addResult := llvmIr.GlobalBuilder.CreateNSWAdd(lvalue, llvm.ConstInt(llvm.Int32Type(), 1, false), "")
		llvmIr.GlobalBuilder.CreateStore(addResult, iii)
		llvmIr.GlobalBuilder.CreateBr(condBB)

		//EndIr
		llvmIr.GlobalBuilder.SetInsertPoint(endBB, endBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = endBB
		return n
	case "putint":
		num := v.visitRule(ctx.ExpressionList()).([]llvm.Value)[0]
		if llvmIr.IntFormat == nil {
			llvmIr.IntFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%d", "intformat")
		}
		llvmIr.GlobalBuilder.CreateCall(llvmIr.PrintfFunc, []llvm.Value{llvmIr.IntFormat.(llvm.Value), num}, "")
		return nil
	case "putch":
		num := v.visitRule(ctx.ExpressionList()).([]llvm.Value)[0]
		if llvmIr.ChFormat == nil {
			llvmIr.ChFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%c", "chformat")
		}
		llvmIr.GlobalBuilder.CreateCall(llvmIr.PrintfFunc, []llvm.Value{llvmIr.ChFormat.(llvm.Value), num}, "")
		return nil
	case "putarray":
		/*		int getarray(int a[]){
				int n;
				scanf("%d",&n);
				for(int i=0;i<n;i++)scanf("%d",&a[i]);
				return n;
		    }*/
		if llvmIr.SpaceIntFormat == nil {
			llvmIr.SpaceIntFormat = llvmIr.GlobalBuilder.CreateGlobalStringPtr(" %d", "spaceintformat")
		}
		if llvmIr.OutputN == nil {
			llvmIr.OutputN = llvmIr.GlobalBuilder.CreateGlobalStringPtr("%d:", "outputn")
		}
		if llvmIr.NewLine == nil {
			llvmIr.NewLine = llvmIr.GlobalBuilder.CreateGlobalStringPtr("\n", "newline")
		}
		v.isCallLibraryFunc = true
		expressionListResult := v.visitRule(ctx.ExpressionList()).([]llvm.Value)
		v.isCallLibraryFunc = false
		n := expressionListResult[0]
		llvmIr.GlobalBuilder.CreateCall(llvmIr.PrintfFunc, []llvm.Value{llvmIr.OutputN.(llvm.Value), n}, "")

		condBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
		bodyBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
		endBB := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "end")

		//CondIR
		//iii := llvm.ConstInt(llvm.Int32Type(), count, false)
		iii := llvmIr.GlobalBuilder.CreateAlloca(llvm.Int32Type(), "")
		llvmIr.GlobalBuilder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 0, false), iii)
		llvmIr.GlobalBuilder.CreateBr(condBB)
		llvmIr.GlobalBuilder.SetInsertPoint(condBB, condBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = condBB
		iiiValue := llvmIr.GlobalBuilder.CreateLoad(iii, "")
		condValue := llvmIr.GlobalBuilder.CreateICmp(llvm.IntSLT, iiiValue, n, "")
		//由CondBB来决定是否跳转到endBB
		llvmIr.GlobalBuilder.CreateCondBr(condValue, bodyBB, endBB)

		//BodyIr
		llvmIr.GlobalBuilder.SetInsertPoint(bodyBB, bodyBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = bodyBB
		trueParam := expressionListResult[1]

		//scanfParam := llvmIr.GlobalBuilder.CreateStructGEP(trueParam, int(iii.ZExtValue()), "")
		printfParam := llvmIr.GlobalBuilder.CreateInBoundsGEP(trueParam, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), iiiValue}, "")
		printfParam = llvmIr.GlobalBuilder.CreateLoad(printfParam, "")
		//scanfParam := llvmIr.GlobalBuilder.CreateInBoundsGEP(trueParam, []llvm.Value{iiiValue}, "")
		llvmIr.GlobalBuilder.CreateCall(llvmIr.PrintfFunc, []llvm.Value{llvmIr.SpaceIntFormat.(llvm.Value), printfParam}, "")
		lvalue := llvmIr.GlobalBuilder.CreateLoad(iii, "")
		addResult := llvmIr.GlobalBuilder.CreateNSWAdd(lvalue, llvm.ConstInt(llvm.Int32Type(), 1, false), "")
		llvmIr.GlobalBuilder.CreateStore(addResult, iii)
		llvmIr.GlobalBuilder.CreateBr(condBB)

		//EndIr
		llvmIr.GlobalBuilder.SetInsertPoint(endBB, endBB.FirstInstruction())
		symbolTable.NowBasicBlock[0] = endBB
		llvmIr.GlobalBuilder.CreateCall(llvmIr.PrintfFunc, []llvm.Value{llvmIr.NewLine.(llvm.Value)}, "")
		return n

	}
	//
	//=============使用了未定义函数============
	funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbolByName(funcName)
	if funcSymbol == nil {
		v.ErrorInfo += "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " \"" + funcName + "\"" + " not defined before\n"
		return nil
		//======================================
	} else {
		var params []llvm.Value
		if ctx.ExpressionList() != nil {
			params = v.visitRule(ctx.ExpressionList()).([]llvm.Value)
		}
		return llvmIr.GlobalBuilder.CreateCall(llvmIr.GlobalMod.NamedFunction(funcName), params, "")
	}
	//return nil
}

func (v *C0Visitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	//函数调用中的 funcCallStmt: IDENTIFIER '(' expressionList? ')' ;
	//expressionList: expression (',' expression)*;
	var params []llvm.Value
	expressions := ctx.AllExpression()
	for _, expression := range expressions {
		param := v.visitRule(expression).(llvm.Value)
		//fmt.Println(param.Type())
		if strings.Contains(param.Type().String(), "Array") && !v.isCallLibraryFunc {
			//llvmIr.GlobalBuilder.CreateInBoundsGEP(param, nil, "")
			//xxx
			param = llvmIr.GlobalBuilder.CreateStructGEP(param, 0, "")
		}
		params = append(params, param)
	}
	return params
}

func (v *C0Visitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.Primary() != nil {

		return v.visitRule(ctx.Primary())
	} else if ctx.GetBop() != nil {
		switch ctx.GetBop().GetTokenType() {
		case parser.C0LexerPLUS:
			//+
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateNSWAdd(value1, value2, "")
		case parser.C0LexerMINUS:
			//-
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateNSWSub(value1, value2, "")
		case parser.C0ParserMUL:
			//*
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateNSWMul(value1, value2, "")
		case parser.C0LexerDIV:
			///
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateSDiv(value1, value2, "")
		case parser.C0LexerMOD:
			//%
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateSRem(value1, value2, "")
		case parser.C0ParserLT:
			//<
			//llvmIr.GlobalMod.Dump()
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntSLT, value1, value2, "")
		case parser.C0ParserGT:
			//>
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntSGT, value1, value2, "")
		case parser.C0ParserGE:
			//>=
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntSGE, value1, value2, "")
		case parser.C0ParserLE:
			//<=
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntSLE, value1, value2, "")
		case parser.C0ParserEQ:
			//==
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntEQ, value1, value2, "")
		case parser.C0ParserNE:
			//!=
			value1 := v.visitRule(ctx.Expression(0)).(llvm.Value)
			if value1.Type() != llvm.Int32Type() {
				value1 = llvmIr.GlobalBuilder.CreateZExt(value1, llvm.Int32Type(), "")
			}
			value2 := v.visitRule(ctx.Expression(1)).(llvm.Value)
			if value2.Type() != llvm.Int32Type() {
				value2 = llvmIr.GlobalBuilder.CreateZExt(value2, llvm.Int32Type(), "")
			}
			return llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value1, value2, "")
		case parser.C0ParserAND:
			//&&
			rightBlock := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
			v.pushWillInBlockStack(rightBlock)
			value1 := v.visitRule(ctx.Expression(0))
			v.popWillInBlockStack()
			if value1 != nil {
				value1a := value1.(llvm.Value)
				if value1a.Type() != llvm.Int1Type() {
					value1a = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value1a, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
				}
				llvmIr.GlobalBuilder.CreateCondBr(value1a, rightBlock, v.topWillOutBlockStack())
			}
			//llvmIr.GlobalBuilder.CreateCondBr(value1, rightBlock, v.topWillOutBlockStack())
			llvmIr.GlobalBuilder.SetInsertPoint(rightBlock, rightBlock.FirstInstruction())
			symbolTable.NowBasicBlock[0] = rightBlock
			value2 := v.visitRule(ctx.Expression(1))
			if value2 != nil {
				value2a := value2.(llvm.Value)
				if value2a.Type() != llvm.Int1Type() {
					value2a = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value2a, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
				}
				llvmIr.GlobalBuilder.CreateCondBr(value2a, v.topWillInBlockStack(), v.topWillOutBlockStack())
			}
			return nil
		case parser.C0ParserOR:
			//||
			rightBlock := llvm.AddBasicBlock(llvmIr.GlobalMod.NamedFunction(symbolTable.GlobalSymbolTable.GetFuncSymbol().Name), "")
			v.pushWillOutBlockStack(rightBlock)
			value1 := v.visitRule(ctx.Expression(0))
			v.popWillOutBlockStack()
			if value1 != nil {
				value1a := value1.(llvm.Value)
				if value1a.Type() != llvm.Int1Type() {
					value1a = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value1a, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
				}
				llvmIr.GlobalBuilder.CreateCondBr(value1a, v.topWillInBlockStack(), rightBlock)
			}
			llvmIr.GlobalBuilder.SetInsertPoint(rightBlock, rightBlock.FirstInstruction())
			symbolTable.NowBasicBlock[0] = rightBlock
			value2 := v.visitRule(ctx.Expression(1))
			if value2 != nil {
				value2a := value2.(llvm.Value)
				if value2a.Type() != llvm.Int1Type() {
					value2a = llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value2a, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
				}
				llvmIr.GlobalBuilder.CreateCondBr(value2a, v.topWillInBlockStack(), v.topWillOutBlockStack())
			}
			return nil
		}
		//u开头的是无符号比较。
		//s开头的是有符号比较n
		/*
			const (
				IntEQ  IntPredicate = C.LLVMIntEQ
				IntNE  IntPredicate = C.LLVMIntNE
				IntUGT IntPredicate = C.LLVMIntUGT
				IntUGE IntPredicate = C.LLVMIntUGE
				IntULT IntPredicate = C.LLVMIntULT
				IntULE IntPredicate = C.LLVMIntULE
				IntSGT IntPredicate = C.LLVMIntSGT
				IntSGE IntPredicate = C.LLVMIntSGE
				IntSLT IntPredicate = C.LLVMIntSLT
				IntSLE IntPredicate = C.LLVMIntSLE
			)
		*/
	} else if len(ctx.AllLSQUAREBRACKET()) > 0 {
		//数组 a[1]这样的取值
		expressions := ctx.AllExpression()
		exprsLen := len(expressions)
		arrayValue := v.visitRule(expressions[0]).(llvm.Value)
		var isFuncParam bool
		funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbol()
		for _, param := range funcSymbol.ArgsType {
			if param.ParamName == expressions[0].GetText() {
				//structGepValue = llvmIr.GlobalBuilder.CreateLoad(structGepValue, "")
				isFuncParam = true
			}
		}
		structGepValue := arrayValue
		for i := 1; i < exprsLen; i++ {
			index := v.visitRule(expressions[i]).(llvm.Value)
			//怎么办呢？？？
			//llvmIr.GlobalBuilder.CreateIntToPtr(index,)
			//temp := llvm.ConstIntCast(index, llvm.Int32Type(), false)
			//temp.ZExtValue()
			//structGepValue = llvmIr.GlobalBuilder.CreateStructGEP(structGepValue, int(index.ZExtValue()), "")
			//structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), index}, "")
			//structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{index}, "")
			if i == 1 && isFuncParam {
				structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{index}, "")
			} else {
				structGepValue = llvmIr.GlobalBuilder.CreateInBoundsGEP(structGepValue, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false), index}, "")
			}
		}
		symbol := symbolTable.GlobalSymbolTable.FindVarSymbol(expressions[0].GetText())
		if len(symbol.Size) > exprsLen-1 {
			return structGepValue
		} else {
			result := llvmIr.GlobalBuilder.CreateLoad(structGepValue, "")
			result.SetAlignment(4)
			return result
		}
		//因为需要用到数组最低维的信息，所以还是得再查一次符号表。
		//这样做肯定是麻烦了
		//varName := ctx.Expression(0).GetText()
		//symbol := symbolTable.GlobalSymbolTable.FindVarSymbol(varName)
		//size := symbol.Size

		//var index int = int(v.visitRule(expressions[exprsLen-1]).(llvm.Value).ZExtValue())
		//a[1][2][3]   6个元素      那么 a[0][1][2]  0*6+1*3+2 = 5
		/*		for i := exprsLen - 2; i > 0; i++ {
					temp := int(v.visitRule(expressions[i]).(llvm.Value).ZExtValue())
					for j := 0; j < exprsLen-i-1; j++ {
						temp *= size[j]
					}
					index += temp
				}
				structGepValue := llvmIr.GlobalBuilder.CreateStructGEP(arrayValue, index, "")
				result := llvmIr.GlobalBuilder.CreateLoad(structGepValue, "")
				result.SetAlignment(4)
				return result*/
		//var
		//先简单测试取值
		/*		arrayValue := v.visitRule(expressions[0]).(llvm.Value)
				indexValue := v.visitRule(expressions[1]).(llvm.Value)
				//总感觉 indexValue.ZExtValue() 之后会出问题
				fmt.Println(indexValue.Type())
				structGepValue := llvmIr.GlobalBuilder.CreateStructGEP(arrayValue, int(indexValue.ZExtValue()), "")
				result := llvmIr.GlobalBuilder.CreateLoad(structGepValue, "")
				result.SetAlignment(4)
				return result*/

	} else if ctx.LBRACE() != nil {
		//进入了一层{}
		v.nowArrayDeep++
		//{1,2,3}这样的。
		//{1,{2}}怎么办呢
		//var values []int
		var values []llvm.Value
		expressions := ctx.AllExpression()
		for _, expression := range expressions {
			visitResult := v.visitRule(expression)
			value, ok := visitResult.(llvm.Value)
			if ok {
				//llvm.Value，直接返回的是值
				//values = append(values, int(value.ZExtValue()))
				values = append(values, value)
			} else {
				//目标expresion是个{x,x,x}这样的
				//要开始补0了。
				//[2][3][4]
				if len(values) != 0 {
					var count int = 1
					for i := v.nowArrayDeep - 1; i < len(v.nowArraySize)-v.nowArrayDeep; i++ {
						count *= v.nowArraySize[i]
					}
					count -= len(values)
					for i := 0; i < count; i++ {
						values = append(values, llvm.ConstInt(llvm.Int32Type(), 0, false))
					}
				}
				newValue := visitResult.([]llvm.Value)
				values = append(values, newValue...)
			}
		}
		//开始补0
		//[2][3][4]
		var needAddZeroCount int = 1
		for i := 0; i <= len(v.nowArraySize)-v.nowArrayDeep; i++ {
			needAddZeroCount *= v.nowArraySize[i]
		}
		needAddZeroCount -= len(values)
		for i := 0; i < needAddZeroCount; i++ {
			values = append(values, llvm.ConstInt(llvm.Int32Type(), 0, false))
		}

		//退出这层{}
		v.nowArrayDeep--
		//返回的是 []int
		return values
	} else if ctx.GetPrefix() != nil {
		//前缀
		switch ctx.GetPrefix().GetTokenType() {
		case parser.C0LexerNOT:
			/*			%3 = load i32, i32* %2, align 4
						%4 = icmp ne i32 %3, 0
						%5 = xor i1 %4, true
						%6 = zext i1 %5 to i32*/
			value := v.visitRule(ctx.Expression(0)).(llvm.Value)
			icmpResult := llvmIr.GlobalBuilder.CreateICmp(llvm.IntNE, value, llvm.ConstInt(llvm.Int32Type(), 0, false), "")
			xorResult := llvmIr.GlobalBuilder.CreateXor(icmpResult, llvm.ConstInt(llvm.Int1Type(), 1, false), "")
			return llvmIr.GlobalBuilder.CreateZExt(xorResult, llvm.Int32Type(), "")

			//return llvmIr.GlobalBuilder.CreateNot(value, "")
		case parser.C0LexerMULPLUS:
			/*
			  ++的IR
			  %3 = load i32, i32* %2, align 4
			  %4 = add nsw i32 %3, 1
			  store i32 %4, i32* %2, align 4
			*/
			//++和--该怎么办呢？？前缀中还有+和-，这两个是干什么的？
			//return llvmIr.GlobalBuilder.CreateADD(value, "")Ï
			//似乎没有自增和自减。
			return nil
		case parser.C0LexerMULMINUS:
			//--
			return nil
		case parser.C0LexerPLUS:
			//不知道到底前缀+和前缀减是什么意思，在go里面测的话，+a就是把a的值变正，-a就是把a的符号变化，无论正负。
			//但是应该可能单纯的是为了表示正负？？？所以暂时不考虑运算符的含义，只是单纯的把这个符号认为是正负的表示。
			//但是拿运算符来弄似乎也没问题。。。
			//前缀+
			//v.nowNumSign = append(v.nowNumSign, 0)
			num := v.visitRule(ctx.Expression(0)).(llvm.Value)
			return num
		case parser.C0LexerMINUS:
			//前缀-
			//v.nowNumSign = append(v.nowNumSign, 1)
			num := v.visitRule(ctx.Expression(0)).(llvm.Value)
			return llvmIr.GlobalBuilder.CreateNSWSub(llvm.ConstInt(llvm.Int32Type(), 0, false), num, "")

		}
	} else if ctx.GetPostfix() != nil {
		//后缀
	} else if ctx.FuncCallStmt() != nil {
		return v.visitRule(ctx.FuncCallStmt())
	} else if ctx.LPARENTHESIS() != nil {
		//(repression)
		return v.visitRule(ctx.Expression(0))
	} else {
		return nil
		//数组，暂时不考虑
	}
	return nil

}

func (v *C0Visitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	//Primary返回的都是值，除了数组
	if ctx.IDENTIFIER() != nil {
		varName := ctx.IDENTIFIER().GetText()
		symbol := symbolTable.GlobalSymbolTable.FindVarSymbol(varName)

		// 表达式中的符号如果在之前并没有定义
		if symbol == nil {
			info := "error: line " + strconv.Itoa(ctx.GetStart().GetLine()) + " " + varName + " not define before\n"
			v.ErrorInfo += info
			return nil
		}
		//===============================

		if symbolTable.NowNamespaceId != 0 {
			funcSymbol := symbolTable.GlobalSymbolTable.GetFuncSymbol()
			for _, param := range funcSymbol.ArgsType {
				if param.ParamName == symbol.Name {
					result := llvmIr.GlobalBuilder.CreateLoad(symbol.Value.(llvm.Value), "")
					result.SetAlignment(4)
					return result
				}
			}
		}
		//llvmIr.GlobalMod.NamedFunction("main").param
		var result llvm.Value
		if symbol.IsConst && symbol.Size == nil {
			result = llvm.ConstInt(llvm.Int32Type(), symbol.Value.(uint64), false)
		} else {
			if symbol.Size != nil {
				//数组，返回的是指针
				result = symbol.Value.(llvm.Value)
			} else {
				result = llvmIr.GlobalBuilder.CreateLoad(symbol.Value.(llvm.Value), "")
				result.SetAlignment(4)
			}
		}
		return result
	} else if ctx.Intliteral() != nil {
		return v.visitRule(ctx.Intliteral())
	} else {
		//其他情况暂时不考虑

	}
	return nil
}

func (v *C0Visitor) VisitIntliteral(ctx *parser.IntliteralContext) interface{} {
	//其实改完才发现可以直接base为0来根据字符串自己判断进制，但是g4懒得改回去了。
	var num int64
	num, _ = strconv.ParseInt(ctx.GetText(), 0, 64)
	value := llvm.ConstInt(llvm.Int32Type(), uint64(num), false)
	return value
}
