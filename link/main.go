package main

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	fmt.Printf("LLVM Library: %+v\n", llvm.Version)
}
