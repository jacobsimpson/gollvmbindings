package main

import (
	"fmt"
	"os"
	"path/filepath"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s <destination-file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	dst := os.Args[1]

	// Return type of the main function.
	returnType := llvm.GlobalContext().Int32Type()

	// LLVM organizes objects hierarchically. A module is the top of the object
	// hierarchy. It corresponds roughly to a file.
	m := llvm.NewModule("Module Name")

	// Make a `main` function to contain the instructions to execute.
	functionType := llvm.FunctionType(returnType, []llvm.Type{}, false)
	function := llvm.AddFunction(m, "main", functionType)

	// Create a builder, in order to add instructions.
	b := llvm.NewBuilder()
	// Use the builder to attach a basic block to the `main` function.
	basicBlock := llvm.AddBasicBlock(function, "initial")
	// Move to the end of the basic block.
	b.SetInsertPointAtEnd(basicBlock)
	// Add the return instruction at the end of the basic block.
	constFive := llvm.ConstInt(returnType, 5, true)
	b.CreateRet(constFive)

	b.Dispose()

	fmt.Println("=====================================================================")
	m.Dump()

	f, err := os.Create(dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file %q:%+v\n", dst, err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(m.String())

	m.Dispose()
}
