package llvmIr

import "github.com/llvm-mirror/llvm/bindings/go/llvm"

var GlobalBuilder = llvm.NewBuilder()
var GlobalMod = llvm.NewModule("feng_module")
var PrintfFunc llvm.Value
var ScanfFunc llvm.Value
var IntFormat interface{} //为了方便与nil比较
var ChFormat interface{}
var SpaceIntFormat interface{}
var OutputN interface{}
var NewLine interface{}
