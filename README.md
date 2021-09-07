# LittleDuck2020+1

Simple lexer + parser combo for LittleDuck2020+1 programming language, using [antlr4](https://github.com/antlr/antlr4) and [gocc](https://github.com/goccmack/gocc). The code generation target is [`Go`](https://golang.org/) for both tools.

![My Little Duck](./img/MyLittleDuck.png?raw=true "My Little Duck")
(**Note**: This is my duck)

## Requirements

### ANTLR

- [antlr4](https://github.com/antlr/antlr4)

### Gocc

- [gocc](https://github.com/goccmack/gocc)

## Running

Make sure you've already fulfilled the requirements. Simply execute the makefile.

```
$ make
```

Both tools are tested using the [`testing`](https://pkg.go.dev/testing) package.
(**Note**: The second input was purposely set as _valid_ for the test, to see the output.)

### ANTLR

While being in `LittleDuck/ANTLR`:

```
$ go test -v .
```

### Gocc

While being in `LittleDuck/Gocc`:

```
$ go test -v .
```
