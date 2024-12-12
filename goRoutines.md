# Go Concurrency  

### **Go Routines**  
- A Go routine is a lightweight thread, created with the `go` keyword.  
- Go routines are much more lightweight than traditional threads, allowing for a large number (up to millions) to run concurrently.  

Example:  
```go  
go f(x, y, z)  
```  

---

### **Channels**  
- Channels allow Go routines to communicate by sending and receiving data.  
- The `<-` operator is used to receive data from a channel.  
- Channels must be created before use:  
  ```go  
  ch := make(chan int)  
  ```

Example:  
```go  
func sum(s []int, c chan int) {  
    sum := 0  
    for _, v := range s {  
        sum += v  
    }  
    c <- sum  // Send sum to channel  
}

func main() {  
    s := []int{7, 2, 8, -9, 4, 0}  
    c := make(chan int)  
    go sum(s, c)  
    fmt.Println(<-c)  // Receive result from channel  
}  
```  

---

### **Buffered Channels**  
- Buffered channels allow sending data into a channel with a defined buffer capacity.  
- Example:  
  ```go  
  ch := make(chan int, 100)  // Buffered channel with capacity 100  
  ```  

---

### **Range and Close with Channels**  
- Channels should be closed to signal completion to the receiver.  
- When using `range` over a channel, it stops when the channel is closed.  
- Failure to close a buffered channel can lead to a panic if the capacity is exceeded.  

Example:  
```go  
func fibonacci(n int, c chan int) {  
    x, y := 0, 1  
    for i := 0; i < n; i++ {  
        c <- x  
        x, y = y, x + y  
    }  
    close(c)  // Close the channel to signal completion  
}

func main() {  
    c := make(chan int, 10)  
    go fibonacci(cap(c), c)  
    for i := range c {  
        fmt.Println(i)  // Read from channel until closed  
    }  
}  
```  

---

### **Select in Go**  
- The `select` statement allows waiting for multiple channel operations.  
- It blocks until one of the channels is ready for communication. If multiple channels are ready, it selects one randomly.  

Example:  
```go  
select {  
    case <-tick:  
        fmt.Println("tick.")  
    case <-boom:  
        fmt.Println("BOOM!")  
        return  
    default:  
        fmt.Println(".")  
        time.Sleep(50 * time.Millisecond)  
}  
```  

---  
