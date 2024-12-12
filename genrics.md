#### Genric 

- Type Parameters 

```go 
func Index[T comparable](s []T,x T){

}
```

**Note**: The genric or the type provide should always implement the comparable type providing the `==` , `!=` etc 

- Genric types in go 

**Note**:  Go has `nil` not `null`

this are declared with the help of `any` 

```go 
type List[T any] struct{
    next *List[T]
    val T
}
```