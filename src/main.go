package main

import (
	"bufio"
	"compiler/llvmIr"
	parser "compiler/parser/g4"
	"compiler/visitor"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llvm-mirror/llvm/bindings/go/llvm"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func evaluation() {
	//写的不对，之后再写了。
	var files []string

	root := "../test/functional/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	/*	for _, file := range files {
		fmt.Println(file)
	}*/
	for i := 0; i < len(files); i = i + 2 {
		outfile := files[i]
		syfile := files[i+1]
		fmt.Println(syfile, ":")
		inputFile, inputError := os.Open(syfile)
		if inputError != nil {
			fmt.Println("file error!")
			return
		}
		defer inputFile.Close()

		inputReader := bufio.NewReader(inputFile)
		var inputString string
		for {
			temp, readerError := inputReader.ReadString('\n')
			inputString += temp
			if readerError == io.EOF {
				break
			}
		}
		is := antlr.NewInputStream(inputString)

		// Create the Lexer
		lexer := parser.NewC0Lexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		// Create the Parser
		p := parser.NewC0Parser(stream)
		v1 := visitor.NewC0Visitor()
		p.Prog().Accept(v1)

		if ok := llvm.VerifyModule(llvmIr.GlobalMod, llvm.ReturnStatusAction); ok != nil {
			fmt.Println(ok.Error())
		}
		engine, err := llvm.NewExecutionEngine(llvmIr.GlobalMod)

		if err != nil {
			fmt.Println("[+]打印错误信息:")
			fmt.Println(err.Error())
		}

		funcResult := engine.RunFunction(llvmIr.GlobalMod.NamedFunction("main"), []llvm.GenericValue{})
		fmt.Println()
		fmt.Printf("%d\n", uint8(funcResult.Int(false)))
		fmt.Println("---------------------")
		inputFile, _ = os.Open(outfile)
		inputReader = bufio.NewReader(inputFile)
		inputString = ""
		for {
			temp, readerError := inputReader.ReadString('\n')
			inputString += temp
			if readerError == io.EOF {
				break
			}
		}
		fmt.Println(inputString)
	}
}

func main() {
	//evaluation()
	//fmt.Println(reflect.TypeOf(C.putint))
	fmt.Println("\n\n[+]start working!")
	// Setup the input
	inputFile, inputError := os.Open("test.sy")
	if inputError != nil {
		fmt.Println("file error!")
		return
	}
	//fmt.Println(reflect.TypeOf(inputFile))
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	var inputString string
	for {
		temp, readerError := inputReader.ReadString('\n')
		inputString += temp
		//fmt.Println(inputString)
		if readerError == io.EOF {
			break
		}
	}
	is := antlr.NewInputStream(inputString)

	// Create the Lexer
	lexer := parser.NewC0Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewC0Parser(stream)
	v1 := visitor.NewC0Visitor()
	//v2 := visitor.NewC0LlvmVisitor()
	//遍历AST，生成符号表并进行语义分析。
	p.Prog().Accept(v1)

	if v1.ErrorInfo != "" {
		fmt.Println(v1.ErrorInfo)
		return
	}

	fmt.Println()
	// verify it's all good
	if ok := llvm.VerifyModule(llvmIr.GlobalMod, llvm.ReturnStatusAction); ok != nil {
		fmt.Println(ok.Error())
	}
	fmt.Println("[+]开始将优化前的LLVM IR写入文件")
	beforeOptIr := llvmIr.GlobalMod.String()
	err := ioutil.WriteFile("./resultIr/beforeOpt.ll", []byte(beforeOptIr), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+]写入完毕")

	fmt.Println("[+]打印优化前的LLVM IR：")
	llvmIr.GlobalMod.Dump()
	fmt.Println()
	//pm := llvm.NewPassManager()
	/*		pmb := llvm.NewPassManagerBuilder()
			fpm := llvm.NewFunctionPassManagerForModule(llvmIr.GlobalMod)

			pmb.PopulateLTOPassManager(pm, false, false)
			pmb.PopulateFunc(fpm)

			fpm.Run(llvmIr.GlobalMod)*/
	//pm.AddPromoteMemoryToRegisterPass()
	//pm.AddDemoteMemoryToRegisterPass()
	/*	pm.AddInstructionCombiningPass()
		pm.AddReassociatePass()
		pm.AddGlobalDCEPass()*/

	//pm.Run(llvmIr.GlobalMod)
	fmt.Println()
	fmt.Println("[+]开始将优化后的LLVM IR写入文件")
	afterOptIr := llvmIr.GlobalMod.String()
	err = ioutil.WriteFile("./resultIr/afterOpt.ll", []byte(afterOptIr), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+]写入完毕")
	fmt.Println()

	fmt.Println("[+]打印优化后的LLVM IR：")
	llvmIr.GlobalMod.Dump()
	//fmt.Println(llvmIr.GlobalMod.String())

	// create our exe engine
	engine, err := llvm.NewExecutionEngine(llvmIr.GlobalMod)

	if err != nil {
		fmt.Println("[+]打印错误信息:")
		fmt.Println(err.Error())
	}

	// run the function!
	fmt.Println("[+]开始运行程序")
	funcResult := engine.RunFunction(llvmIr.GlobalMod.NamedFunction("main"), []llvm.GenericValue{})
	//fmt.Println()
	//fmt.Println("[+]打印程序的运行结果：")
	//截断结果
	fmt.Printf("%d\n", uint8(funcResult.Int(false)))

}

//pm := llvm.NewPassManager()
/*	pmb := llvm.NewPassManagerBuilder()
	fpm := llvm.NewFunctionPassManagerForModule(llvmIr.GlobalMod)

	pmb.PopulateLTOPassManager(pm, false, false)
	pmb.PopulateFunc(fpm)

	fpm.Run(llvmIr.GlobalMod)*/
/*	pm.AddPromoteMemoryToRegisterPass()
	pm.AddDemoteMemoryToRegisterPass()
	pm.AddInstructionCombiningPass()
	pm.AddReassociatePass()
	pm.AddGlobalDCEPass()

	pm.Run(llvmIr.GlobalMod)*/
