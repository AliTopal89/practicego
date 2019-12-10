# practicego  [![Build Status](https://travis-ci.org/AliTopal89/practicego.svg?branch=master)](https://travis-ci.org/AliTopal89/practicego)

## this README and curriculum

is mostly by the work that I followed from:

Chris James [quii](https://github.com/quii/learn-go-with-tests)

Go is a statically typed language meaning the type of a variable is known at compile time. To elaborate:
 - Strongly typed languages, such as Go, don't allow for "type coercion" - `("3" + 5 )`: the ability for a value to change type implicitly in certain contexts (e.g. merging two types using +).
 - Types are checked before run time
 - Static typed languages are those in which type checking is done at compile-time(`Code translated into something the computer can read before run-time`), whereas dynamic typed languages (`Code translated during execution`) are those in which type checking is done at run-time(`Period when program is executing commands`).

## Important notes

### Declaring Variables, Constants, Switch

 Declaring functions, with arguments and return types:

```go
import (
    "fmt"
)
```

With import `"fmt"` we are importing a package 
which contains the Println function that we use to print.

An argument of type string - means "an array, 
where every element inside of it is at least a String, 
which we will name arguments".

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}
```

variable Prefix will be assigned the "zero" value. 
This depends on the type, for example ints are 0 and for strings it is "".

```go
func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case turkish:
        prefix = turkishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}
```

A switch statement is a shorter way to write a sequence of `if - else` statements.
Go only runs the selected case, not all the cases that follow. 
In effect, the break statement that is needed at the end of each case in 
those languages is provided automatically in Go.

A "return" statement without arguments returns the named return values. 
Named return parameters are initialized to their zero values. This is known as a "naked" or "bare" return.

Example:

```go
func f() (i int, s string) {
    i = 17
    s = "abc"
    return // same as return i, s
}
```

Naked return statements should be used only in short functions, as with the example shown here. 
In the longer functions They can harm readability.

### Integers

<!-- https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc -->

```go
package integers
```
Go source files can only have one package per directory. The name of the package should match the name of the directory in which its source is located. If you package is called `integers`, then its source files may be located in

```$GOPATH/src/github.com/yourname/integers```


```go
...

import "testing"

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}
```

Using `%d` vs `%s`. That's because we want it to print an integer rather than a string.

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```
Please note that the example function will not be executed if you remove the comment `//Output: 6`. 
Although the function will be compiled, it won't be executed.

By adding this code `ExampleAdd()`, the example will appear in the documentation inside `godoc`.
Godoc examples are snippets of Go code that are displayed as package documentation 
and that are verified by running them as tests. To try this out, run `godoc -http=:6060` 
and navigate to `http://localhost:6060/pkg/`

### Iteration

To do stuff repeatedly in Go, you'll need for. In Go there are no "while", "do", "until"

```go
func BenchmarkRepeat(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Repeat("a", 20)  
    }
}
```
When the benchmark code is executed, 
it runs b.N times and measures how long it takes.

```go
func main(){
    ...
    f := "apple"
    fmt.Println(f)
}
```

The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

Unlike regular variable declarations, is a short variable declaration may redeclare variables provided they were originally declared earlier in the same block

### Arrays and Slices

```go
package arrays

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```
Range lets you iterate over an array. Every time it is called it returns two values, the index and the value. We are choosing to ignore the index value by using _ blank identifier, It's a bit like writing to the Unix /dev/null file.

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. 
Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. Slices are indexable and have a length. But unlike arrays, they can be resized.

The type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.


```go
// Slice of type `int`
var s []int
```
The slice is declared just like an array except that we do not specify any size in the brackets `[]`.

```go
func SumAll(numbersToSum ...[]int) []int {
    ...
    return
}
```
A variadic function accepts variable number of input values — zero or more. Ellipsis (three-dots) prefix in front of an input type makes a function variadic. 
Variadic functions can be called with any number of trailing arguments. Variadic function is a function which accepts a variable number of arguments. Variadic functions are also functions but they can take an infinite or variable number of arguments.

For example the signature append function:

`func append(slice []Type, elems ...Type) []Type`

You will see `elems ...Type` which means pack all incoming arguments into elems slice after the first argument.

One important thing to notice is that only the last argument of a function is allowed to be variadic.

So the first argument to append function will be a slice because it demands a slice but later arguments will be packed into one argument elems. When you do not pass any argument in the variadic function, then the slice inside the function is nil. The variadic functions are generally used for string formatting.

```go
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1,2}, []int{0,9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```
Use `reflect.DeepEqual` which is useful for seeing if any two variables are the same. It's important to note that reflect.DeepEqual is not "type safe", you can even compare a slice integer with a string.

```go
func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)
    ...
```
make allows you to create a slice with a starting capacity of the len of the numbersToSum  The make built-in function allocates and initializes an object of type slice, map, or chan (only).

```go
package main
import "fmt"
func main() {
  s := make([]string, 3)
  fmt.Println("emp:", s)
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])
}
// OUTPUT:
// emp: [  ]
// set: [a b c]
// get: c
```
To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).

```go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```

A slice can also be formed by "slicing" an existing slice or array. Slicing is done by specifying a half-open range with two indices separated by a colon. For example, the expression b[1:4] creates a slice including elements 1 through 3 of b above. 

A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array.

```go
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:] 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```

Slicing does not copy the slice's data. It creates a new slice value that points to the original array. This makes slice operations as efficient as manipulating array indices. Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.

### Structs, methods & table driven tests

A struct is a sequence of named elements, called fields, each of which has a name and a type.

A method is a function with a receiver. A method declaration binds an identifier, the method name, 
to a method, and associates the method with the receiver's base type. Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on "things".

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. In many other programming languages this is done implicitly and you access the receiver via this.

```go
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```
Given defined type Point, the declarations bind the methods `Length` and `Scale`, with receiver type `*Point`, to the base type Point.

#### Useful Resources:
1. [GoLang Guide](https://golang.org/doc/)
1. [Static vs. Dynamic](https://hackernoon.com/i-finally-understand-static-vs-dynamic-typing-and-you-will-too-ad0c2bd0acc7)
1. [Variadic Functions](https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085)
1. [GoLang Slices](https://www.callicoder.com/golang-slices/)
1. [Reflections in Go](https://golangbot.com/reflection/)
1. [Slices Usage](https://blog.golang.org/go-slices-usage-and-internals)


