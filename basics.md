# Go Programming Concepts  

### **Packages**  
- Every Go program is composed of packages.  
- Programs start execution in the `main` package.  
- Import paths specify the packages being used (e.g., `"fmt"`, `"math/rand"`).  
- **Naming Convention:** The package name matches the last element of its import path (e.g., `"math/rand"` has the package name `rand`).  

---

### **Exported Names**  
- A name is **exported** if it begins with a capital letter:  
  - Example: `Pizza` and `Pi` are exported, while `pizza` and `pi` are not.  
- Only exported names can be accessed from outside a package.  

**Fixing Unexported Name Error:**  
Rename `math.pi` to `math.Pi` to access the exported name.  

---

### **Named Return Values**  
- Named return values act as variables defined at the top of a function.  

Example:  
```go  
func split(sum int) (x, y int) {  
    x = sum * 4 / 9  
    y = sum - x  
    return // Naked return  
}  
```  

---

### **Declaring Variables**  
- Example:  
  ```go  
  var c, python, java bool  
  ```  
- **Note:** The `:=` construct is unavailable outside functions since all statements must start with a keyword (`var`, `func`, etc.).  

---

### **Type Conversion**  
- Converts a value `v` to type `T` using `T(v)`.  

Examples:  
```go  
// Package-level  
var i int = 42  
var f float64 = float64(i)  
var u uint = uint(f)  

// Function-level  
i := 42  
f := float64(i)  
u := uint(f)  
```  

---

### **Type Inference**  
- Go infers types based on the initial value assigned.  

Example:  
```go  
var x = 42    // x is inferred as int  
y := "hello"  // y is inferred as string  
```  

---

### **Constants**  
- Constants are declared using the `const` keyword and cannot use `:=`.  

Example:  
```go  
const world = "world"  
```  

---

### **Format Specifiers**  
| Specifier | Description |  
|-----------|-------------|  
| `%v`      | Value representation |  
| `%t`      | Type representation |  
| `%g`      | General for floating-point/scientific notation |  
| `%e`      | Scientific notation |  
| `%f`      | Floating-point notation |  

---  
