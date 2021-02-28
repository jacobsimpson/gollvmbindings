# Building the LLVM Library and Golang Bindings

I found it difficult to figure out how to build LLVM and link it to my Golang
program. It could be I didn't find the right documentation. In any case, I
wrote down the steps I took to get LLVM working as a Golang library.

Get the LLVM source. It's (relatively) recently been moved to Github, so some
documentation that is out there points to the old SVN repository.

```
git clone https://github.com/llvm/llvm-project.git
```

LLVM requires you build the build system, before you can build the project.
Multiple build systems are available. This mechanism generates the `ninja`
build system.

```
cd llvm-project
mkdir build
cd build
cmake -G Ninja ../llvm
```

In order to build the Golang bindings, something needs to be done about
modules. Golang modules were first introduced in Go 1.11. As far as I can tell,
in order to build the Golang bindings, either modules have to be turned off for
the Go compiler, or the bindings code needs to be enabled as a module.

`GO111MODULE` environment variable will disable modules for the Go compiler.
One of two mechanisms can be used to set it for the build.

Based on a code review, it looks like this will do the trick, no need to modify
the code. I haven't tried it yet so I'm not 100% sure. If it doesn't work,
proceed to the next mechanism, which is what I actually used.

```
export GO111MODULE=off
```

Add a line to `llvm_config.go` the turned Golang modules off. This option will
probably be removed from the Go compiler in the future.

```
../llvm/bindings/go/llvm/llvm_config.go
+       newenv = append(newenv, "GO111MODULE=off")
before os.StartProcess(...)
```

Or, instead of disabling modules for the Go compiler, configure the bindings
directory as a module:

```
cd ../llvm/bindings/go/llvm
go mod init llvm.org/llvm/bindings/go/llvm
```

In order to succeed in building the bindings, I needed to build the library
first. The `check-llvm-bindings-go` target will generate the necessary
libraries. There might be a better target, I wasn't able to find anything that
looked relevant.

```
cd llvm/bindings/go/llvm/workdir/llvm_build
ninja check-llvm-bindings-go
```

Once the libraries and binaries are generated, add them to the LLVM binaries to
the shell PATH. One of the commands in the Golang bindings build script will
try to execute one of the LLVM binaries, so the generated binaries need to be
available:

```
# llvm-go.go will execute llvm-config. Make sure it is on the PATH.
export PATH=$(pwd)/bin/:$PATH
```

From what I can determine, the `llvm-go` command is supposed to generate the
installation specific `../llvm/bindings/go/llvm/llvm_config.go` file. It
contains the paths of all the libraries used by the Golang bindings, and is
used by the Go compiler to link the LLVM libraries. As far as I can tell the
build script attempts to provide `llvm_config.go` as a command line argument to
indicate the destination to save the generated code to. However, the
`llvm_config.go` program doesn't take arguments, and in order to generate the
file, you have to run the command manually with a redirection.

```
./llvm-go print-config > ../llvm/bindings/go/llvm/llvm_config.go
```

Use the LLVM script to build the Golang bindings:

```
cd ../..
./build.sh
```

See the [link](link) project for information on how to link a Go program to the
LLVM bindings.

## References

1.   https://dannypsnl.github.io/blog/2017/12/04/cs/llvm-go-bindings/
1.   https://dannypsnl.github.io/blog/2018/10/06/cs/test-llvm-go-binding-in-travis/
1.   https://dev.to/dannypsnl/test-llvm-go-binding-in-travis-1292
1.   https://felixangell.com/blogs/an-introduction-to-llvm-in-go
1.   https://pkg.go.dev/github.com/llir/llvm/ir
1.   https://github.com/go-llvm/llvm
1.   https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/
     - Native golang that generates IR code.
1.   https://llvm.org/docs/GettingStarted.html#getting-the-source-code-and-building-llvm
     - follow these instructions to build.
