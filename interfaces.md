- Interface


declared as below :


```go 
type inter interface{
    meth1(),
    meth2(),
    meth3()
}
```

for implementing the interface you just need to have the methods added to the strucutre and it will atomatically implement the interface 

```go
type struc struct{

}

func (s struc) meth1(){

}

func (s struc) meth2(){
}


func (s struc) meth3(){
}

// here the struc has implemented the iteer interface now soo can use struc as a interface inter

func getdetails(i inter){

}
var s=struc{

}

getdetails(s)
```

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements



-Tye Switches


- Named Interfaces 

```go 
type Copier interface{
    Copy(sourceFile string , destinationFile string) (bytesCopied int)
}
```

- defer in go 

Purpose: The defer statement in Go is used to schedule a function call to be executed after the surrounding function completes. It is commonly used for cleanup tasks, such as closing files, releasing resources, or unlocking mutexes.

```go 
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
output:
`hello world`


the defer functions output are pushed onto stack  and then pop accordingly 

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

Output :

```powershell
counting
done
9
8
7
6
5
4
3
2
1
0
```

