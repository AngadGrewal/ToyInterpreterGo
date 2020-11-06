### ToyInterpreterGo

A Go program that is a mini-interpreter for a toy programming language that allows the following: 

1. Reading of a .txt file
2. The use of variables that consist of a single letter (e.g. A, a, ...) 
3. The use of whole numbers: ( e.g. -1, -20, 0, 1, 200)
4. Assignment (=): ( e.g. A = B, A = 10, ) 
5. Addition of exactly two elements (variables or constants) (+) : ( e.g. C = A + B, D = 1 + A, ... )
6. The ability to "return" a value when a single variable or constant is on a line by itself ( e.g. A, B, 10)

The input in the toy interpreter will be of the following variety. 

```
A = 2
B = 8
C = A + B
C
```
and return the return value of the toy program, in this case, the value 10. 

It will be tested with the following 

```
A = 2
B = 22
Z = 91
K = A + B
Z = K + A
Z
```
Expecting 26

and 

```
A = 2 + 1
B = A + 9
C = A + B
C 
``` 
Expecting 15

Task List
- [x] Write a method to be able to read a file 
- [x] return lines in file as an array
- [x] recognise whether a statement is an assignment, addition or return
- [x] if a line has assignment and addition, recognise it as an assignment
- [x] work out how assignment would work in terms of reading from instructions.
- [x] Create a map that acts as a variable store
- [x] Make the maths work (make sure that the program can read multiple digits)
