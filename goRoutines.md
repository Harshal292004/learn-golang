- Go Routines

`go f(x,y,z)` go Routines are like thread but lighter than threads making a difference of upto a thousand or even a million threads 

- Channels 

`<-` Channel operator 

Like maps and slices, channels must be created before use:

`ch := make(chan int)`

declared as below 

```go 
func sum(s []int , c chan int){
    sum:=0
    for _,v:= range s{
        sum+=v
    }

    c <- sum //send sum to c 
}

func main(){
    s:= []int{7,2,8,-9,4,0}
    c:=
}
```


- Buffered Channels

`ch := make(chan int, 100)`

Buffer value :->100


- Range and Close with channels 

when usgin range over a buffered channel it has a sepcific capacity whcich when exceeded can cuase panic 
thus we need to close the channel so that range knows that the routine cant foward any more data and it needs to stop 

```go 
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```


- Select in Go 

The select statement in Go is used to wait for multiple channels to be ready for communication. Think of it like a "switch" for channels — it waits for whichever channel is ready and then performs the associated action.

Blocking Behaviour : When using `select` it blocks the execution until one of the channels involved is reasy for communication. if multiple are ready it picks at random 


```go 
select {
    case <-tick:
        fmt.Println("tick.")
    case <-boom:
        fmt.Println("BOOM!")
        return
    default:
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)
}
```