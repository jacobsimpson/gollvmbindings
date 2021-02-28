# generateir

Small program that uses the LLVM library and Golang bindings to generate some
simple IR.

The documentation (this document) contains the additional steps to translate
LLVM IR code into an executable.

## How to generate an executable from LLVM IR code

It can be useful to see how a program would compile to IR, as you build your
own IR representation. `clang` is a C++ compiler based on LLVM, and with a
commandline parameter will emit IR in an `.ll` file.

```
clang -S -emit-llvm -o data/sample.ll data/sample.c
```

Using the LLVM command line utilities it is possible to compile an IR file to
an executable.

```
llc -filetype=obj -o out.o out.ll
clang out.o
```

Or:

```
llc -filetype=asm -x86-asm-syntax=intel -o out.asm out.ll
as example.asm -o example.o
clang example.o
```

There is a quick sample program and `Makefile` that provides an example of how
a C program can be used as a comparison to the output generated by the Golang
LLVM program.

## Development

To compile and use the program in this directory:

```
make
./a.out
echo $?
```

There is also an LLVM IR interpreter available to directly execute the LLVM IR:

```
lli out.ll
```
