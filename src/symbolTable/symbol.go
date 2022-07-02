package symbolTable

import "github.com/llvm-mirror/llvm/bindings/go/llvm"

//是否需要更多的信息？
//需要更多的信息之后再添加？
type Symbol struct {
	// 变量类型因为只有int和int[],就不设置了。
	//如果不是数组，Size设-1。
	//size不是-1就认为是数组。
	Name        string      // 变量名
	NamespaceId int         // 作用域，为0即为全局
	Size        []int       // 数组size 从最低维开始放入的
	IsConst     bool        //是否是常量const
	Value       interface{} //为了适配const才改成了interface{}，实际上这里也可以直接llvm.Value，懒得改回去了，毕竟程序现在跑得起来。
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

func NewSymbol(name string, namespaceId int, size []int, isConst bool, value interface{}) *Symbol {
	return &Symbol{name, namespaceId, size, isConst, value}
}

func NewFuncSymbol(name string, returnType int, argsType []FuncParam) *FuncSymbol {
	return &FuncSymbol{name, returnType, argsType}
}
