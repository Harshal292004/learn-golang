Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.



Exported names
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again



- Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

- Naked return 

return without args 

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```


- Declaring variables 

```go
var c,python,java bool
```

**Note**

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available

- Type Conversion

The expression T(v) converts the value v to the type T.

```go 
//package level
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
// function level
i := 42
f := float64(i)
u := uint(f)
```

- Type inference 

```
```

- Constants 

Constants cannot be declared using the := syntax.

```go
const world="world"
```


- format specifiers

`%v`=> for value 
`%t`=> for type 
`%g`=> general for the floating and scientific notation
`%e`=> Scientific notation
`%f`=> floating notation 