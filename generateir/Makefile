
default::
	go build .
	./generateir out.ll
	llc -filetype=obj -o out.o out.ll
	clang out.o

clean::
	rm out.ll out.o a.out

