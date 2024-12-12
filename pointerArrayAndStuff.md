# Go Concepts  

### **Pointers**  
- Go pointers are similar to C pointers, but pointer arithmetic is not allowed.  
- Example:  
  ```go  
  func main() {  
      i, j := 42, 34  
      p := &i  // Pointer to i  
      q := &j  // Pointer to j  
      p = p + q  // Error: pointer arithmetic not allowed  
      *p = *p / 37  // No problem: dereferencing is allowed  
  }  
  ```  

---

### **Pointers to Structs**  
- You can use pointers to structs, and access or modify their fields via the pointer.  
- Example:  
  ```go  
  type Vertex struct {  
      x, y int  
  }

  func main() {  
      v := Vertex{1, 2}  
      p := &v  
      p.x = 3435  // Modify using pointer  
      // Or use:  
      (*p).x = 3435  // Dereferencing pointer to access field  
  }  
  ```  

---

### **Arrays**  
- Arrays are fixed-size collections of elements.  
- Example:  
  ```go  
  var a [2]string  
  a[0] = "Hello"  
  a[1] = "World"  

  primes := [6]int{2, 3, 5, 7, 11, 13}  // Initializing array  
  ```  

---

### **Slices**  
- Slices are more flexible than arrays, as they allow dynamic resizing.  
- Example:  
  ```go  
  primes := []int{2, 3, 5, 7, 11, 13}  
  var s []int = primes[1:4]  // Slice from index 1 to 3 (half-open)  
  fmt.Println(s)  // Output: [3 5 7]  
  ```

- **Important points:**  
  - **Changing a slice** changes the corresponding elements in the underlying array.  
  - Slices have both **length** (number of elements) and **capacity** (size of underlying array).  
  - Zero-value of a slice is `nil`.  

- **Using `make` to create slices:**  
  ```go  
  a := make([]int, 0, 5)  // Create slice of length 0 and capacity 5  
  ```

---

### **Slice of Slices (2D Arrays)**  
- A slice of slices can be used to represent 2D arrays.  
- Example:  
  ```go  
  board := [][]string{  
      []string{"_", "_", "_"},  
      []string{"_", "_", "_"},  
      []string{"_", "_", "_"},  
  }  
  board[0][0] = "X"  
  board[2][2] = "O"  
  ```  

---

### **Range**  
- The `range` keyword is used to iterate over slices and maps.  
- Example with a slice:  
  ```go  
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}  
  for i, v := range pow {  
      fmt.Printf("2**%d = %d\n", i, v)  
  }
  ```

- To skip values:  
  ```go  
  for _, v := range pow  // Skip index  
  for i, _ := range pow  // Skip value  
  ```  

---

### **Maps**  
- Maps are unordered collections of key-value pairs.  
- Example of map declaration and initialization:  
  ```go  
  var m map[string]Vertex  // Declare map  
  m = make(map[string]Vertex)  // Initialize map  
  m["Bell Labs"] = Vertex{40.68433, -74.39967}  // Add an element  
  ```  

- **Map Literal:**  
  ```go  
  var m = map[string]Vertex{  
      "Google": {37.42202, -122.08408},  
      "Bell Labs": {40.68433, -74.39967},  
  }  
  ```  

- **Mutating a Map:**  
  - Insert or update: `m[key] = value`  
  - Retrieve: `value = m[key]`  
  - Delete: `delete(m, key)`  
  - Test presence: `value, ok = m[key]`  

---

### **Passing Functions as Arguments**  
- Functions in Go can be passed as arguments to other functions.  
- Example:  
  ```go  
  func compute(fn func(float64) float64) {  
      return fn(3)  
  }

  func main() {  
      fn := func(x float64) float64 {  
          return x * x  // Square the number  
      }  
      compute(fn)  
  }  
  ```  

---