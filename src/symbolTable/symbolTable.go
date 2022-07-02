package symbolTable

import (
	"fmt"
	"github.com/llvm-mirror/llvm/bindings/go/llvm"
)

//是否需要维护这么多全局变量？？？
var GlobalSymbolTable = NewFrogStack(0, nil, nil)
var NowNamespaceId = -1
var NextNamespaceId = 0
var NowBasicBlock []llvm.BasicBlock = make([]llvm.BasicBlock, 1)

type Frame struct {
	Length      int
	Symbols     []*Symbol //栈是否会不利于查找？。。。是否要考虑利用堆或者哈希表等来优化查找？
	NamespaceId int
}
type ProgStack struct {
	Length int
	//BasicBlocks []llvm.BasicBlock
	Funcs  []*FuncSymbol //并不是栈
	Frames []*Frame      //每进入一个block，在栈顶新加入一层，退出block的时候弹出这一层；
}

func NewFrame(length int, symbols []*Symbol, namespaceId int) *Frame {
	return &Frame{length, symbols, namespaceId}
}
func NewFrogStack(length int, funcs []*FuncSymbol, frames []*Frame) *ProgStack {
	return &ProgStack{length, funcs, frames}
}

/*func NewFrogStack(length int, basicBlocks []llvm.BasicBlock,funcs []*FuncSymbol, frames []*Frame) *ProgStack {
	return &ProgStack{length,basicBlocks, funcs, frames}
}*/
/*func (st *ProgStack) PushBasicBlock(bb llvm.BasicBlock){
	st.BasicBlocks = append(st.BasicBlocks,bb)
}
func (st *ProgStack) PopBasicBlock(){
	st.BasicBlocks = st.BasicBlocks[:len(st.BasicBlocks)-1]
}*/
func (st *ProgStack) PushFrame(frame *Frame) {
	st.Frames = append(st.Frames, frame)
	//fmt.Println(st.Frames)
	st.Length++
	NowNamespaceId = NextNamespaceId
	NextNamespaceId++

}
func (st *ProgStack) PopFrame() {
	if st.Length < 1 {
		panic("frames is empty unable to pop")
	}
	st.Frames = st.Frames[:st.Length-1]
	st.Length--
	if st.Length == 0 {
		NowNamespaceId = -1
	} else {
		NowNamespaceId = st.Frames[st.Length-1].NamespaceId
	}
}
func (st *ProgStack) TopFrame() *Frame{
	if st.Length < 1 {
		panic("frames is empty")
		return  nil
	} else {
		return  st.Frames[st.Length-1]
	}
}

func (st *ProgStack) AddSymbol(symbol *Symbol) {
	frame := st.Frames[st.Length-1]
	frame.Symbols = append(frame.Symbols, symbol)
	frame.Length++
}

/*func (st *ProgStack) UpdateSymbol(name string,value llvm.Value) {
	st.FindVarSymbol(name)
}*/
func (st *ProgStack) AddFunc(funct *FuncSymbol) {
	st.Funcs = append(st.Funcs, funct)
}

/*func (st *ProgStack) SetFuncName(name string) {
	st.Funcs[len(st.Funcs)-1].Name = name
}*/
func (st *ProgStack) SetFuncReturnType() {
	//因为默认是1:int
	st.Funcs[len(st.Funcs)-1].ReturnType = 0
}
func (st *ProgStack) AddFuncArgsType(name string, ttype llvm.Type) {
	funcParam := FuncParam{name, ttype}
	st.Funcs[len(st.Funcs)-1].ArgsType = append(st.Funcs[len(st.Funcs)-1].ArgsType, funcParam)
}

func (st *ProgStack) GetFuncSymbol() *FuncSymbol {
	return st.Funcs[len(st.Funcs)-1]
}
func (st *ProgStack) GetFuncSymbolByName(name string) *FuncSymbol {
	for _, i := range st.Funcs {
		if i.Name == name {
			return i
		}
	}
	return nil
	//panic("[+]error!!!使用了符号表中不存在的函数")
}
func (st *ProgStack) FindVarSymbol(varName string) *Symbol {
	//暂时简单逻辑
	//从栈顶从上到下查。如果查不到就报错
	for i := len(st.Frames) - 1; i >= 0; i-- {
		frame := st.Frames[i]
		for _, j := range frame.Symbols {
			if varName == j.Name {
				return j
			}
		}
	}
	//debug.PrintFrames()
	fmt.Println("[+]开始打印当前符号表中的所有frame信息")
	for _, i := range GlobalSymbolTable.Frames {
		fmt.Println("frame length", i.Length)
		fmt.Println("frame namespaceId", i.NamespaceId)
		for _, j := range i.Symbols {
			fmt.Println(j.Name, j.NamespaceId, j.Size, j.IsConst)
		}
		fmt.Println()
	}
	fmt.Println("[+]打印完成\n")
//	panic("[+]error!!!使用了符号表中不存在的变量:" + varName)
	return nil
}
