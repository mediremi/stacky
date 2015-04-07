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
; Let's divide the two numbers on the stack
Div
; Since literal values are integers, we will see [0]
DBGSTK