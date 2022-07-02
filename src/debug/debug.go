package debug

import (
	"compiler/symbolTable"
	"fmt"
)

func PrintFrames() {
	fmt.Println("[+]开始打印当前符号表中的所有frame信息")
	for _, i := range symbolTable.GlobalSymbolTable.Frames {
		fmt.Println("frame length", i.Length)
		fmt.Println("frame namespaceId", i.NamespaceId)
		for _, j := range i.Symbols {
			fmt.Println(j.Name, j.NamespaceId, j.Size, j.IsConst)
		}
		fmt.Println()
	}
	fmt.Println("[+]打印完成\n")
}
func PrintFuncs() {
	fmt.Println("[+]开始打印当前符号表中的所有函数信息")
	for _, i := range symbolTable.GlobalSymbolTable.Funcs {
		fmt.Println("function name", i.Name)
		fmt.Println("function ReturnType", i.ReturnType)
		fmt.Print("function ArgsTypes：")
		for x1, x2 := range i.ArgsType {
			fmt.Print("name:", x1, " , type:  ", x2)
		}
		fmt.Println("\n")
	}
	fmt.Println("[+]打印完成\n")
}
