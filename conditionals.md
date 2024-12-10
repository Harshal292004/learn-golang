- Only for loop 

```go
for i:=0 ; i< 10 ; i++{
    sum+=i
}
```


- Go's While 

```go 
func main(){
    sum:=1
    for sum<1000{
        sum+=sum
    }

    fmt.Println(sum)
}
```

- infinity 

```go 
for{

}
```

- Conditionals 

```go 
if x<0{

}

if v := math.Pow(x, n); v < lim {
	return v
}

```


- Switch Case 

Notice that in Go, the break statement is not required at the end of a case to stop it from falling through to the next case. The break statement is implicit in Go.

```go 
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"

    // all three of these cases will set creator to "A Steve"
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