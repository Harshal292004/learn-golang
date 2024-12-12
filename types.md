### **Structs**  
- A **struct** is a collection of fields, which can be of different types.  
- Example:  
  ```go  
  type car struct {  
      Make   string  
      Model  string  
      Height int  
      Width  int  
  }  
  ```

---

### **Anonymous Structs**  
- **Anonymous structs** are used for one-time use cases, such as in an HTTP header.  
- Example:  
  ```go  
  myCar := struct {  
      brand string  
      model string  
  }{  
      brand: "Tesla",  
      model: "Model 3",  
  }  
  ```

---

### **Embedded Structs**  
- Go is not a traditional OOP language, but it supports **embedding structs**, allowing a form of inheritance. The fields of the embedded struct are inherited by the outer struct.  
- Example:  
  ```go  
  type car struct {  
      brand string  
      model string  
  }

  // truck embeds car, so it inherits car's properties  
  type truck struct {  
      car    // embedded struct  
      window int  
  }
  ```

- **Assignment example:**  
  ```go  
  lanestruck := truck{  
      window: 10,  
      car: car{  
          brand: "Toyota",  
          model: "Whatever",  
      },  
  }  
  ```

---

### **Struct Methods in Go**  
- Methods can be defined on structs, allowing you to associate behavior with struct types.  
- Example of a method on a struct:  
  ```go  
  type rect struct {  
      width, height int  
  }

  func (r rect) area() int {  
      return r.width * r.height  
  }  

  func main() {  
      r := rect{width: 10, height: 5}  
      fmt.Println("Area:", r.area())  
  }  
  ```  

---