# Go Control Structures  

### **For Loop**  
- The only looping construct in Go.  
- Example:  
  ```go  
  for i := 0; i < 10; i++ {  
      sum += i  
  }  
  ```  

---

### **Go’s While Equivalent**  
- A `for` loop without initialization and post statements behaves like a `while`.  

Example:  
```go  
func main() {  
    sum := 1  
    for sum < 1000 {  
        sum += sum  
    }  
    fmt.Println(sum)  
}  
```  

---

### **Infinite Loop**  
- A `for` loop with no conditions creates an infinite loop.  

Example:  
```go  
for {  
    // Infinite loop body  
}  
```  

---

### **Conditionals**  
- Simple `if` statement:  
  ```go  
  if x < 0 {  
      // Code block  
  }  
  ```  

- Short variable declaration in condition:  
  ```go  
  if v := math.Pow(x, n); v < lim {  
      return v  
  }  
  ```  

---

### **Switch Case**  
- No `break` is needed; it’s implicit in Go.  
- Use `fallthrough` to execute subsequent cases.  

Example:  
```go  
func getCreator(os string) string {  
    var creator string  
    switch os {  
    case "linux":  
        creator = "Linus Torvalds"  
    case "windows":  
        creator = "Bill Gates"  

    // Multiple cases with fallthrough  
    case "Mac OS":  
        fallthrough  
    case "Mac OS X":  
        fallthrough  
    case "mac":  
        creator = "A Steve"  

    default:  
        creator = "Unknown"  
    }  
    return creator  
}  
```  

---