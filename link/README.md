# link

The purpose of this project is to do the simplest program that links to the
LLVM library.

The build of the LLVM library and bindings are outside of this project.

Key points in this project are:

1.   The LLVM package import in `main.go`.
2.   I manually added these lines to `go.mod` to specify the locally built
     library and Golang bindings.

    ```
    require llvm.org/llvm/bindings/go/llvm v0.0.0
    replace llvm.org/llvm/bindings/go/llvm v0.0.0 => /Users/jsimpson/src/llvm/src/llvm.org/llvm/bindings/go/llvm
    ```

## Development

Should be simple. If LLVM is built, and the changes to `go.mod` are correct,
this should be sufficient.

```
go build .
```

This build runs extra long compared to a regular go build.
