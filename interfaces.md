# Go Concepts  

### **Interfaces**  
- An **interface** defines a set of methods. Any type that implements all the methods of the interface is said to implement that interface.  
- Example of interface declaration:  
  ```go  
  type inter interface {  
      meth1()  
      meth2()  
      meth3()  
  }  
  ```  

- A type implements an interface by providing method definitions. No explicit declaration of implementation is required.  
  Example:  
  ```go  
  type struc struct{}  

  func (s struc) meth1() {}  
  func (s struc) meth2() {}  
  func (s struc) meth3() {}  

  // struc automatically implements the inter interface  
  func getdetails(i inter) {}  

  var s = struc{}  
  getdetails(s)  
  ```  

- **Key Point:** A type implements an interface automatically if it has the required methods.  
- **Empty Interface:** Every type implements the empty interface `interface{}` because it has no method requirements.  

---

### **Type Switches**  
- Type switches allow handling multiple types in a single statement.  

Example:  
```go  
switch v := i.(type) {  
case int:  
    fmt.Println("int", v)  
case string:  
    fmt.Println("string", v)  
default:  
    fmt.Println("unknown type")  
}  
```  

---

### **Named Interfaces**  
- Interfaces can be named and used to define behaviors.  

Example:  
```go  
type Copier interface {  
    Copy(sourceFile, destinationFile string) (bytesCopied int)  
}  
```  

---

### **Defer in Go**  
- The `defer` statement schedules a function to be executed after the surrounding function completes. It’s commonly used for cleanup tasks like closing files or releasing resources.  
- Deferred function calls are pushed onto a stack and executed in Last-In-First-Out (LIFO) order.  

Example:  
```go  
func main() {  
    defer fmt.Println("world")  
    fmt.Println("hello")  
}  
// Output:  
// hello world  
```  

- Example of multiple `defer` statements:  
```go  
func main() {  
    fmt.Println("counting")  

    for i := 0; i < 10; i++ {  
        defer fmt.Println(i)  
    }  

    fmt.Println("done")  
}  
// Output:  
// counting  
// done  
// 9 8 7 6 5 4 3 2 1 0  
}  
```  

---
