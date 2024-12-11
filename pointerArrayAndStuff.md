- Pointers

Like C but no pointer arthemtic like below 

```go
//throws error 
func main(){
    i,j:=42,34
    p:= &i //point to i
    q:= &j //point to j 
    p=p+q //throws error 

    *p= *p/37 // no problem here 
}
```

- Pointer to struct 


```go 
type Vertex struct{
    x int 
    y int
}

func main(){
    v:= Vertex{1,2}
    p:= &v
    p.x=3435
    //or 
    (*p).x=3435
}
```

- Arrays 

```go 
var a [2] string 
a[0]= "Jello"
a[1]= "Wld"

primes:= [6]int{2,3,4,5}

```

- Slices 

```go 
primes := []int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
fmt.Println(s)
```

half open includes the start excludes the last 

**Changing the elements of a slice modifies the corresponding elements of its underlying array.**

- *Slices have a capacity and length*

the cpacity is the length of the underlying array and the length is the number of elements contained by the slice 

- Zero slice is nill

- The Make function for the slice making 

```go 
//creating dynamically sized array 
// a:= make([]int,len(a),cap(a))
a:= make([]int,0,5)
```

- Slice of slices or 2D array 

```go 
board:=[][]string{
    []string{"_","_","_"},
    []string{"_","_","_"},
    []string{"_","_","_"}
}

board[0][0]="X"
board[2][2]="O"
```

- Range 

```go 
var pow=[]int{1,2,4,8,16,32,64,128}
//  similar to js map((index,item)=>{})
for i,v:=range pow{
    fmt.Printf("2**%d = %d\n",i,v)
}
//skip the index or item 
for _,v:=range pow 
for i,_:= range pow
```

- Maps 

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

```go 
var m map[string]Vertex

m= make(map[string]Vertex)
m["BL"]=Vertex{
    40.683,-74.54
}
```

```go
//map literal 
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

- Mutating map 

Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]



- Passing functions as argument 

```go 
func compute(fn func(float64) float64){
    return fn(3)
}

func main(){
    fn:=func(x float64) float64{
        return x
    }
    compute(fn)
}
```