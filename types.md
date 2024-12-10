- Structs 

```go
type car struct{
    Make string 
    Model string 
    Height int 
    Width int 
}
```

- Anonymous Struct 

they are meant to be used a single time eg.http header 

```go 
myCar:=struct{
    brand string 
    model string 
}{
    brand:"tesla",
    model:"model 3"
}

```

- Embedded Struct 

Go not a oop language but can have a data inherited using 

```go 
type car struct{
    brand string 
    model string 
}
//car's all props are provided to the truck 
type truck struct{
    car 
    window int 
}
```

assignment

```go 
lanestruck := truck{
    windo : 10,
    car:car{
        brand:"toyota"
        model:"whatefver"
    }
}
```


- struct method in go 

```go 
type rect structure{
    
}
```