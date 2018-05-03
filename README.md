# Slang (**S**-expression **Lang**uage)

## Installation

To install slang as a library:

`$ go get -u github.com/zachorosz/slang`

To install the slang library and command:

`$ go get -u github.com/zachorosz/slang/cmd/slang`

The above will install the library and the `slang` binary to `$GOPATH/bin`. Installing the `slang` command will also install its dependency, [go-prompt](github.com/c-bata/go-prompt), for the REPL environment.

## Slang

`usage: slang [[-e expression | filename] [arguments]]`

slang can evaluate and print a single expression from the command line using the `-e` flag.

```
$ slang -e "(+ 1 2)"
3
```

Arguments can be passed along with the flag. Evaluating the `*ARGV*` symbol will print a Vector of your arguments as strings.

```
$ slang -e "*ARGV*" 1 2 3
["1" "2" "3"]
```

Executing slang without any flags or arguments starts the REPL.

```
$ slang
slang>
```

To exit the REPL, evaluate `(exit)`.

```
$ slang
slang> (exit)
```

slang can also read and evaluate a file as a program. Pass the path to the file you wish to evaluate as the first argument:

```
$ slang /path/to/my-program.sl
```

Any additional arguments are treated as program arguments. Using `*ARGV*` in your program will allow you to interact with the arguments Vector.

## Some very useful resources

1. [Structure and Interpretation of Computer Programs](https://mitpress.mit.edu/sicp/full-text/book/book.html) by Gerald Jay Sussman and Hal Abelson
2. [kanaka's Make a Lisp](https://github.com/kanaka/mal) - parsing and tail-call optimization
3. Rob Pike's "Lexical Scanning in Go" talk - [video](https://www.youtube.com/watch?v=HxaD_trXwRE), [slides](https://talks.golang.org/2011/lex.slide#1)