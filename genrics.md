# Go Generics  

### **Type Parameters**  
- Functions can accept type parameters to enable generic programming.  
- The type parameter is specified in square brackets (`[T]`).  

Example:  
```go  
func Index[T comparable](s []T, x T) {  
    // Implementation  
}  
```  
**Note:** The type parameter must satisfy constraints, such as `comparable` for types supporting `==`, `!=`.  

---

### **Generic Types in Go**  
- Generic types are declared using the `any` constraint.  
- `any` is a built-in alias for `interface{}`.  

Example:  
```go  
type List[T any] struct {  
    next *List[T]  
    val  T  
}  
```  

**Note:** Go uses `nil`, not `null`, for representing zero values or absence of value.  

---