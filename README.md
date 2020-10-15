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

Task List
- [x] Write a method to be able to read a file 
- [x] return lines in file as an array