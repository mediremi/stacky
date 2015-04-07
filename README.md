# Stacky - A stack-based language
---

Stacky is a stack-based language implemented in [go](https://golang.org).

```
; (This is a comment)
; Let's push a number to the stack
Literal
42
; Now, let's add two numbers
Literal
2
Literal
2
Add
; Let's view the stack
DBGSTK
; [42 4] should have been printed to your screen
```

You can run the above example with `stacky examples/tutorial.st` (after
[installing](#install) stacky).

## Install
`go get github.com/medimatrix/stacky`

You can now run `stacky <stacky_file>`.

## Reference

### Comments
Any line that starts with a semicolon (`;`) is considered to be a comment and
will be ignored. Comments must be on their own line.

```
; This is a comment
;;??;So is this
Literal ; This does not work (currently ☺)
```

### Instructions

#### Print
The next instruction is the length of the string. The last n (where n is the
length provided) values on the stack are popped off and printed to `stdout`.

```
Literal
'H'
Literal
'i'
Print
2
```

#### Add
The last 2 values on the stack are added, popped off, and then the result of the
addition is pushed on to the stack.

#### Mult
The last 2 values on the stack are multiplied, popped off, and then the result
of the multiplication is pushed on to the stack.

#### Literal
The next instruction is interpreted as a value. Use `Literal` to push arguments
to the stack. If a character is provided, it is converted to its ASCII value.

```
Literal
42
Literal
2
Add
; The stack is now [4]
```

#### DBGSTK
The current state of the stack is printed to `stdout`. (DBGSTK -> *D*e*b*u*g* *ST*ac*K*)

## Examples
See the [examples folder](examples).

## License
Licensed under the [MIT license](LICENSE.md).
