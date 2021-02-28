
default::
	clang -S -emit-llvm -o sample.ll sample.c
	llc -filetype=asm -x86-asm-syntax=intel -o sample.asm sample.ll
	as sample.asm -o sample.o
	clang -o sample sample.o

clean::
	rm -f sample.ll sample.asm sample.o sample

