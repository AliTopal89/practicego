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

The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for `var f string = "apple"` in this case.

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
Range lets you iterate over an array. Every time it is called it returns two values, the index and the value. We are choosing to ignore the index value by using `_` blank identifier, It's a bit like writing to the Unix /dev/null file.

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. 
Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. Slices are indexable and have a length. But unlike arrays, they can be resized.

The type specification for a slice is `[]T`, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.


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
A variadic function accepts variable number of input values — zero or more. **Ellipsis** (three-dots) prefix in front of an input type makes a function variadic.
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
To create an empty slice with non-zero length, use the built-in make. Here we make a slice of strings of length 3 (initially zero-valued).

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
to a method, and associates the method with the receiver's base type. 

```go
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```
Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on "things".

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. In many other programming languages this is done implicitly and you access the receiver via `this`.

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

Interfaces allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety. 

```go
type Shape interface {
    Area() float64
}
```

This is creating a new type just like it did with Rectangle and Circle but this time it is an `interface` rather than a `struct.`

In our case:

 - `Rectangle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface.
 - `Circle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface.

 Like a `struct` an `interface` is created using the type keyword, followed by a name and the keyword `interface`. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.


Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.

```go
func TestArea(t *testing.T) {

   areaTests := []struct {
       name    string
       shape   Shape
       hasArea float64
   }{

       {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
       {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
       {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
   }
```

What we’d like to do is to set up all the inputs and expected outputs and feel them to a single test harness - `"tests := []test{name rectangle..."`
This is a great time to introduce table driven testing. We can use an anonymous struct literal `[]struct` to reduce the boilerplate so  now you can name the fields of instances of structs in the test fixture.


```go
   for _, tt := range areaTests {
       t.Run(tt.name, func(t *testing.T) {
           got := tt.shape.Area()
           if got != tt.hasArea {
               t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
           }
       })

   }

}
```
The `%#v` format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.

By wrapping each case in a `t.Run` you will have clearer test output on failures as it will print the name of the case for example:

```
--- FAIL: TestArea (0.00s)
    --- FAIL: TestArea/Rectangle (0.00s)
        shapes_test.go:33: main.Rectangle{Width:12, Height:6} got 72.00 want 72.10
```  

### Pointers & Errors

```go
type Wallet struct {
    balance int
}
```
In Go if a symbol (variables, types, functions etc.) starts with a lowercase symbol then it is private outside the package it's defined in.

```go
func (w Wallet) Deposit(amount int) {
    w.balance += amount
}

func (w Wallet) Balance() int {
    return w.balance
}
```
Now we can access the internal `balance` field in the struct using the **receiver** variable.

In Go, **when you call a function or a method the arguments are copied.**

When calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.

```go
...
 wallet.Deposit(10)

got := wallet.Balance()

fmt.Printf("origin of balance in test is %v \n", &wallet.balance)
...
```
you create a value - like a wallet, it is stored somewhere in memory. You can find out what the origin of that bit of memory with `&myVal.` We get the pointer to a thing(method) with the origin of symbol; `&`.

A variable is just a convenient, alphanumeric pseudonym for a memory location; a label, or nickname.

Now, rather than talking about memory locations, we can talk about variables, which are convenient names we give to memory locations.

Now that we know that memory is just a series of numbered cells, and variables are just nicknames for a memory location assigned by the compiler, what is a pointer?

>A pointer is a value that points to the memory address of another variable.

The pointer points to memory address of a variable, just as a variable represents the memory address of value.

>**Prereq of method set, interface type, interface** 

>The `method set` of an `interface type` is its interface. The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T. An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.

> ```go
>  type Stringer interface {
>      String() string
>  }
 > ```

>something satisfies/implements this interface if it has a method with the exact signature `String() string`.

>```go
>// Declare a Book type which satisfies the fmt.Stringer interface.
>type Book struct {
>    Title  string
>    Author string
>}
>
>func (b Book) String() string {
>    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
>}
>
>// Declare a Count type which satisfies the fmt.Stringer interface.
>type Count int
>
>func (c Count) String() string {
>    return strconv.Itoa(int(c))
>}
>```

>we have two different "interface types", `Book` and `Count`, which do different things. But the thing they have in common is that they both satisfy the `fmt.Stringer` *interface*.

>Dereferencing a pointer means using the `*` operator to retrieve the value from the memory address that is pointed by the pointer.

```go
type Wallet struct {
	balance int
 }

func(w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func(w *Wallet) Balance() int {
    return w.balance
```

create a new type from original and be more descriptive, which also allows to declare methods on them

```go
type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
```

```go
assertError := func(t *testing.T, got error, want string) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got.Error() != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

We've introduced `t.Fatal` which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around

```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }
```

We don't really care what the exact wording is, just that some kind of meaningful error around withdrawing is returned given a certain condition.

In Go, errors are values, so we can refactor it out into a variable and have a single source of truth for it.

### Maps

```go 
func main() {
   var m = make(map[string]int)
}
```
Initializing a map using the built-in make() function. The function will return an initialized and ready to use map -

```go
// Initialize an empty map
var m = map[string]int{}
```

You can also create an empty map using a map literal by leaving the curly braces empty like above.

`m[key] = value`

You can add new items to an initialized map and you can retrieve the value assigned to a key in a map using the same syntax above.

```go
name, ok := employees[1010]  // "", false
```

If you try to access a key that doesn’t exist, then the map will return an empty string "" (zero value of strings), and false -

If you just want to check for the existence of a key without retrieving the value associated with that key, then you can use an _ (underscore) in place of the first value so it reads `_ok := employees[1010]`

```go
func TestSearch(t *testing.T) {
    dictionary := map[string]string{"test": "this is just a test"}
```
starts with the map keyword and requires two types. The first is the key type, which is written inside the `[]`. The second is the value type, which goes right after the `[]`.

```go
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
```
`m[key] = value`You can add new items to an initialized map and you can retrieve the value assigned to a key in a map using the same syntax above.

Assertions allow you to easily write test code, and are global funcs in the `assert` package. All assertion functions take, as the first argument, the `*testing.T` object provided by the testing framework. This allows the assertion funcs to write the failings and other details to the correct place

```go
    assertStrings(t, err.Error(), want)
```

Errors can be converted to a string with the `.Error()` method, which we do when passing it to the assertion. We are also protecting `assertStrings` with **if** to ensure we don't call `.Error()` on `nil`.

A two-value assignment tests for the existence of a key:

```i, ok := m["route"]```

In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn't exist, (i) is the value type's zero value (0). The second value (ok) is a boolean that is true if the key exists in the map, and false if not.

```_, ok := m["route"]```

To test for a key without retrieving the value, use an underscore in place of the first value:

You can modify maps without passing them as a pointer. This is because map is a reference type. Meaning it holds a reference to the underlying data structure, much like a pointer. The underlying data structure is a `hash table`, or `hash map`.

`var m map[string]string`

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function

`var dictionary = make(map[string]string)`

or 

`var dictionary = map[string]string{}`

so you can have an empty hash map and point dictionary at it, without runtime panic. 

```go
func (d Dictionary) Delete(word string) {
    delete(d, word)
}
```

Go has a built-in function `delete` that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.

### Dependency Injection

```go
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

Under the hood Printf just calls Fprintf passing in `os.Stdout`.

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error){

}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
}
```

`Fprint` formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered. 

`Fprintf` formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.

`Fprintf` is like `printf` again. Here instead on displaying the data on the monitor, or saving it in some string, the formated data is saved on a file which is pointed to by the file pointer which is used as the first parameter to fprintf. 

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Write writes `len(p)` bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.

```go
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.FPrintf(w, "hello/n")
}

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":8090", nil)
}
```
A fundamental concept in net/http servers is handlers. A handler is an object implementing the **http.Handler** interface. A common way to write a handler is by using the `http.HandlerFunc` adapter on functions with the appropriate signature.

Functions serving as handlers take a `http.ResponseWriter` and a `http.Request` as arguments. The response writer is used to fill in the HTTP response. Here our simple response is just “hello\n”. We register our handlers on server routes using the `http.HandleFunc` convenience function. It sets up the default router in the net/http package and takes a function as an argument.

### Mocking

`buffer.bytes` it’s an adaptor that lets you use a byte slice as an io.Writer and turn strings/byte slices into io.Readers.

```go
func (b *Buffer) Bytes() []byte
```
byte is an alias for `uint8` and is equivalent to `uint8` in all ways. It is used, by convention, to distinguish byte values from `8-bit` unsigned integer values. Bytes returns a slice of length `b.Len()` holding the unread portion of the buffer. The slice is valid for use only until the next buffer modification(that is, only until the next call to a method like Read, Write, Reset, or Truncate).

`Bytes.Buffer`: Often we want to build up a long sequence of bytes. With bytes.Buffer we can write bytes into a single buffer, and then convert to a string when we are done. For performance, using bytes.Buffer is an ideal choice. And it can simplify some situations where we want to append many values together.

```go
func main(){
    fmt.Println("Using 'bytes standard package'")

    var b bytes.Buffer
    b.Write([]byte("Hello, World!\n"))
    fmt.Fprintf(b, "Holly %s\n, smokes Batman!")
}
```
we have this bytes buffer and we can use `b.Write` and thats fine because even if you look at implementation documentation you can see that all its a pointer to a buffer is the receiver for that write method, but the reason this will work when we are talking about methods, if you remember `b.Write` is shorthand for enclosing b and taking the address of it and then calling a write method on it, so we can do that. The reason why line `fmt.Fprintf(b,..)` isn't going to works is because now you are passing it to a function a copy of that buffer, and inside that function it wants to take a pointer, so you can get it working by passing a pointer in our buffer `fmt.Fprintf(&b), "Holy %s\n"`

###### bytes vs strings
Byte slices represent a mutable, resizable, contiguous list of bytes
  - Given a slice of bytes:
 `buf := []byte{1,2,3,4}`
  - It’s mutable so you can update elements:
  `buf[3] = 5  // []byte{1,2,3,5}`
  - It’s resizable so you can shrink it or grow it:
    ```go
    buf = buf[:2]           // []byte{1,2}
    buf = append(buf, 100)  // []byte{1,2,100}`
    ```
Strings, on the other hand, represent an immutable, fixed-size, contiguous list of bytes. That means that you can’t update a string — you can only create new ones.

```go
func NewReader(b []byte) *Reader
func NewReader(s string) *Reader
```

The in memory reader functions above return `io.Reader`*(Reader is the interface that wraps the basic Read method.*
*Read reads up to len(p) bytes into p. It returns the number of bytes read `(0 <= n <= len(p))` and any error encountered.)* implementation that wraps around your in-memory byte slice or string, which also implement all the read-related interfaces in `io`.

```go
got := buffer.String()
	want := `3
			 2
			 1
			 Go!`
```

The backtick syntax is another way of creating a string but lets you put things like newlines.

The Go `==` operator compares not just the time instant but also the Location and the *monotonic clock*(is for measuring time) reading.
Dependency injection is a technique whereby one struct supplies the dependencies of another struct, by explicitly providing components with all of the dependencies they need to work. Dependency injection is a useful tool for decoupling logical entities.

```go
type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}
```

Spies are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc. In our case, we're keeping track of how many times `Sleep()` is called so we can check it in our test.

In Go increment and decrement operations can’t be used as expressions, only as statements. Without pointer arithmetic, the convenience value of pre- and postfix increment operators drops. By removing them from the expression hierarchy altogether, expression syntax is simplified and the messy issues around order of evaluation of `++` and `--` (consider `f(i++)` and `p[i] = q[++i]`) are eliminated as well.

```go
type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)
}
```

We are using duration to configure the time slept and sleep as a way to pass in a sleep function. The signature of `sleep` is the same as for `time.Sleep` allowing us to use `time.Sleep` in our real implementation and a *spy* in our tests.

### Concurrency

Concurrency means multiple computations are happening at the same time - *having more than one thing in progress.*

Normally in Go when we call a function `doSomething()` we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is *blocking* - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a *goroutine*. 

Because the only way to start a goroutine is to put go in front of a function call, we often use "anonymous functions" when we want to start a *goroutine*.

```go
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        go func() {
            results[url] = wc(url)
        }()
    }

    return results
}
```

Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that the're declared - this is what the `()` at the end of the anonymous function is doing. Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

 >Lexical scoping (sometimes known as static scoping ) is a convention used with many programming languages that sets the scope (range of functionality) of a variable so that it may only be called (referenced) from within the block of code in which it is defined.

The body of the anonymous function above is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function) each of which will add its result to the results map.

```go
for _, url := range urls {
        go func(u string) {
            results[u] = wc(u)
        }(url)
    }
```
By giving each anonymous function a parameter for the url - `u` - and then calling the anonymous function with the `url` as the argument, we make sure that the value of `u` is fixed as the value of `url` for the iteration of the loop that we're launching the goroutine in. `u` is a copy of the value of `url`, and so can't be changed.

```go
for _, str := range []string{"a", "b", "c"} {
    // Anonymous function
	go func() {
		fmt.Println(str)
	}()
}
<-time.After(1 * time.Second)
```

Before `<-time.After(1 * time.Second)` goroutine exits just after finishing the for loop, this means that the program exits and the goroutines that print each of the values are not scheduled, so we get no ouput.

A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by a function invocation

Race condition [example](concurrency/race_detector/README.md)

Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes. Channels by default are **pointers**. Mostly, when you want to communicate with a goroutine, you pass the channel as an argument to the function or method. Hence when goroutine receives that channel as an argument, you don’t need to dereference it to push or pull data from that channel.

```go
// Send statement
resultChannel <- result{u, wc(u)}
/*
    instead of writing to the "map" directly we're sending a "result" struct for each call 
    to "wc" to the "resultChannel" with a *send statement*. This uses the "<-" operator,

*/
...
// Receive expression
result := <-resultChannel

/*
  The next "for" loop iterates once for each of the "urls". Inside we're using a *receive expression*, 
  which assigns a value received from a `channel` to a variable. This also uses 
  the "<-" operator, but with the two operands now reversed: the channel is now on the right 
  and the variable that we're assigning to is on the left
*/
```

### Select

Go has a special statement called `select` which works like a `switch` but for `channels`.

`http.HandlerFunc` is a type that looks like this: `type HandlerFunc func(ResponseWriter, *Request)`

All it's really saying is it needs a function that takes a **ResponseWriter** and a **Request**, which is not too surprising for an HTTP server.

It turns out there's really no extra magic here, this is also *** how you would write a real HTTP server in Go. *** The only difference is we are wrapping it in an `httptest.NewServer` which makes it easier to use with testing, as it finds an open port to listen on and then you can close it when you're done with your test.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
The main function begins with a call to `http.HandleFunc`, which tells the http package to handle all requests to the web root ("/") with handler.

It then calls `http.ListenAndServe`, specifying that it should listen on port 8080 on any interface (":8080"). This function will block until the program is terminated.`ListenAndServe` always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with log.Fatal.

The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.

An `http.ResponseWriter` value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.

An `http.Request` is a data structure that represents the client HTTP request. `r.URL.Path` is the path component of the request URL, in this case `/monkeys`. The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name.

If you run this program and access the URL:

```http://localhost:8080/monkeys```

the program would present a page containing:

```Hi there, I love monkeys!```

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
	fmt.Println("hello")
}

/*
hello
world
*/
```
A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

##### Buffered Channels
It's also possible to pass a second parameter to the make function when creating a channel:

```go 
c := make(chan int, 1)
```

This creates a buffered channel with a capacity of 1. Normally channels are synchronous; both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.

##### select

Introducing new construct called `select` which helps us synchronise processes really easily and clearly.

```go
func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
```

If you recall from the concurrency chapter, you can wait for values to be sent to a channel with `myVar := <-ch`. This is a *blocking call*, as you're waiting for a value.

What `select` lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.

We use `ping`  in our select to set up two channels for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).

```go

func ping(url string) chan struct{} {
    ch := make(chan struct{})
    go func() {
        http.Get(url)
        close(ch)
    }()
    return ch
}

```
A `chan struct{}` is the smallest data type available from a memory perspective so you get no allocation versus a bool.


**Always `make` channels**

Notice how we have to use make when creating a channel; rather than say `var ch chan struct{}`. When you use `var` the variable will be initialised with the "zero" value of the type. So for string it is `""`, int it is `0`, etc.

For channels the zero value is `nil` and if you try and send to it with `<-` it will block forever because you cannot send to nil channels

### Reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types. Reflection builds on the type system.
Reflection is the ability of a program to inspect its variables and values at run time and find their type.

Why do we even need to inspect a variable and find its type at runtime when each and every variable in our program is defined by us and we know its type at compile time itself? In cases where we don’t know of the concrete types in compile time, reflections come into play.

The reflect package implements run-time reflection in Go. The reflect package helps to identify the underlying concrete type and the value of a `interface{}` variable.

##### Types and interfaces

> One important category of type is `interface types`, which represent fixed sets of methods. An interface variable can store any concrete (non-interface) value as long as that value implements the interface's methods.
>
>io.Reader and io.Writer, the types Reader and Writer from the io package:
>
>```go
>// Reader is the interface that wraps the basic Read method.
>// Any type that implements a Read method with this signature is said to implement io.Reader. 
>
>// that means that a variable of type io.Reader can hold any value whose type has a Read method
>type Reader interface {
>    Read(p []byte) (n int, err error)
>}
>
>// Writer is the interface that wraps the basic Write method.
>// Any type that implement a Write method with this signature is said to implement io.Writer
>
> // that means that an "interface" variable of type io.Writer can hold any value whose "interface" type  has a Write "interface" method.
>type Writer interface {
>    Write(p []byte) (n int, err error)
>} 
>```
>
>```go
>
>// Whatever concrete "non interface" value "r" may hold "r's" type is always "io.Reader"
>var r io.Reader
>r = os.Stdin
>r = bufio.NewReader(r)
>r = new(bytes.Buffer)
>```

Another example:

>```go
> // concrete type
>type file struct {
>    name string
>}
> // concrete type
>type pipe struct {
>    name string
>}
>
>func main() {
>    var f file // you now have a value of type file
>    var p pipe // you now have a value of type pipe
>    fmt.Println(f, p)
>}
>```

> ```go
> // interface type
> type reader interface {
>   read(b []byte) (int, error)
> }
>```
>
> An `interface type` is the opposite of a `struct` type (concrete type). 
> An interface type can only declare a method set of behavior. This means there
> is nothing concrete about an interface type.
>```go
>  func retrieve(r reader) error {
>    data := make([]byte, 100)
>
>    len, err := r.read(data)
>    if err != nil {
>        return err
>    }
>    fmt.Println(string(data[:len]))
>    return nil
>}
>```
>When you look at the function declaration for retrieve, the function
>seems to say, pass me a value of type reader. But you know that’s impossible
>because there is no such thing as a value of type reader. Values of type reader
>does not exists because reader is an interface type. You know interface values are
>valueless.
>
>Then what is the function declaration saying? It’s saying the following:
>
>*I will accept any piece of concrete data (any value or pointer) that implements*
>*the reader contract. That implements the full method set of behavior defined by the reader interface.*


important example of an interface type is the empty interface empty interface: `interface{}`

represents the empty set of methods and is satisfied by any value at all, since any value has zero or more methods.

`interface{}` An empty interface can be used to hold any data and it can be a useful parameter since it can work with any type.

```go
type emptyInterface struct {
   typ  *rtype            // word 1 with type description
   word unsafe.Pointer    // word 2 with the value
}
```

```go
type rtype struct {
   size       uintptr
   ptrdata    uintptr
   hash       uint32
   tflag      tflag
   align      uint8
   fieldAlign uint8
   kind       uint8
   alg        *typeAlg
   gcdata     *byte
   str        nameOff
   ptrToThis  typeOff
}
```

##### Representation of interfaces

A variable of interface type stores a pair: the concrete value assigned to the variable and the value's type descriptor. To be more precise, the value is the underlying concrete data item that implements the interface and the type describes the full type of that item

>```go
> var r io.Reader
>// r contains, schematically, the (value, type) pair, (tty, *os.File)
>// even though the interface value provides access only to the Read method, the value inside carries all the type information about that value.
>//               type                methods
> tty, err := os.Openfile("/dev/tty", os.O_RDWR, 0)
>// os.O_RDWR   int = syscall.O_RDWR   // open the file read-write. 
>// type *os.File implements methods other than Read; even though the interface value provides access only to the Read method
> if err != nil {
>    return nil, err
>}
>    //concrete val
>  r = tty
>
>```


```go
var w io.Writer
w = r.(io.Writer)

```

The expression in this assignment is a type assertion; what it asserts is that the item inside `r` also implements `io.Writer`, and so we can assign it to `w`. After the assignment, `w` will contain the pair `(tty, *os.File)`, same pair as in `r`.

The static `type(io.Reader)` of the interface determines what methods may be invoked with an interface `variable(r)`, even though the concrete `value(tty)` inside may have a larger set of `methods(("/dev/tty", os.O_RDWR, 0)`.

Then you can:

```go
var empty interface{}
empty = w

```

and our empty interface value empty will again contain that same pair, `(tty, *os.File)`. That's handy: an empty interface can hold any value and doesn't need explicit type assertion

A `concrete type` is a regular type, it specifies the exact representation of the data and the methods, data specifically, but also methods that are used in the type of the receiver type. And `interface type` just specifies some method signatures. So no data is specified, just the methods.

```go
...
func main(){
    Println(true)
	Println(2010)
	Println(9.15)
	Println(7 + 7i)
}

func Println(x interface{}){
    fmt.Printf("type is '%T', value: %v\n",x,x)
}

// type is 'bool', value: true
// type is 'int', value: 2010
// type is 'float64', value: 9.15
// type is 'complex128', value: (7+7i)
```

On `x` we get the type of the value that is stored in `x`. It looked like type of `x` was changing but type of `x` remained the same as an empty interface, what was changing was the information that `x` contain and thats what was dynamic.  Type of the varial is not changing but its the the informations about values that `x` is storing is changing. 

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i:=0; i< val.NumField(); i++ {
        field := val.Field(i)
        fn(field.String())
    }
}
```
The `reflect` package has a function `ValueOf` which returns us a Value of a given variable.
`Field `returns the `i'th` field of the struct value
`val` has a method `NumField` which returns the number of fields in the value. 

Note: A literal of a value is a text representation of the value in code. `Composite literal` is something that can make our lives easier when we want to assign values upon a variable initialization. Literals in source code allow to specify fixed values like numbers, strings or booleans In go Keys and values from `composite literals` must be assignable to respective keys, elements or struct fields.

Floating point literal example

```go
1.23e2  // == 123.0
123E2   // == 12300.0
123.E+2 // == 12300.0
```
numeric literal example

```go
0_33_77_22   // == 0337722
0x_Bad_Face  // == 0xBadFace
0X_1F_FFP-16 // == 0X1FFFP-16
```

string literals example

```go
// The following interpreted string literals are equivalent.
"\141\142\143"
"\x61\x62\x63"
"\x61b\x63"
"abc"
```

The concrete type of `interface{}` is represented by reflect.Type and the underlying value is represented by reflect.Value. There are two functions `reflect.TypeOf()` and `reflect.ValueOf()` which return the reflect.Type and reflect.Value respectively.

```go
func (v Value) Call(in []Value) []Value
```

Call calls the function v with the input arguments in. For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]). Call panics if v's Kind is not Func. It returns the output results as Values.


### Sync

Package sync provides basic synchronization primitives such as mutual exclusion locks.

A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time. This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex

 - Lock
 - Unlock

```go
mutex.Lock() 

x = x + 1 // this statement be executed
          // by only one Goroutine 
          // at any point of time  

mutex.Unlock()
```

`sync.WaitGroup` which is a convenient way of synchronising concurrent processes.

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set
> the number of goroutines to wait for. Then each of the goroutines runs and calls Done when
> finished. At the same time, Wait can be used to block until all goroutines have finished.

Functions of the form below

```func BenchmarkXxx(*testing.B)```

are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

#### Goroutines

Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as light weight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence it's common for Go applications to have thousands of Goroutines running concurrently.

#### Defer

```go
package main

import "fmt"

func foo() {
    defer fmt.Println("Goku: No Vegeta")
    fmt.Println("Vegeta: I will destroy Earth Kakarot hahaha")
}

func main() {
    foo()
```

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.


#### Context

A package that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.

```go
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

A *Context* does not have a `Cancel` method for the same reason the `Done` channel is *receive-only*: the function receiving a cancelation signal is usually not the one that sends the signal. In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the `WithCancel` function (described below) provides a way to cancel a new Context value.

```go
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()
```
context has a method `Done()` which returns a channel which gets sent a signal when the context is "done" or "cancelled". We want to listen to that signal and call `store.Cancel` if we get it but we want to ignore it if our Store manages to Fetch before it.

To manage this we run `Fetch` in a goroutine and it will write the result into a new channel `data`.

```go

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

```

The function Server takes a Store and returns us a `http.HandlerFunc`. Store is defined as the returned function calls the store's Fetch method to get the data and writes it to the response.


```go
type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
```

simulating a slow process where we build the result slowly by appending the string, character by character in a goroutine. When the goroutine finishes its work it writes the string to the `data` channel. The goroutine listens for the `ctx.Done` and will stop the work if a signal is sent in that channel.

Finally the code uses another `select` to wait for that goroutine to finish its work or for the cancellation to occur.


The problem with `context.Values` is that it's just an untyped map so you have no type-safety and you have to handle it not actually containing your value. You have to create a coupling of map keys from one module to another and if someone changes something things start breaking.

In short, if a function needs some values, put them as typed parameters rather than trying to fetch them from `context.Value`. This makes it statically checked and documented for everyone to see.

`"context"` package does solve these problems, better than anything else for Go:

- The cancelation channels are usually not accepted by other libraries and functions and thus cancelation is only possible ‘in-between’ the slow operations.

- Considering a ‘tree of goroutines’ (where children goroutines are the ones spawned by the parent goroutines), it’s easy to cancel the whole tree (just close the cancelation channel), but it’s harder to cancel a sub-tree (you need to introduce another channel for that, or some other solution).

#### Porperty Based Tests 

`strings.Builder` - A Builder is used to efficiently build a string using Write methods.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

}
// Output:
//3...2...1...ignition
```
In addition to providing the `WriteString`, `WriteRune`, and `WriteByte` methods, the string builder also implements the `io.Writer` interface. At first this may not seem very important - why would we want to write a byte slice when we could just write a string? - but because the string builder implements the io.Writer interface it means that we can use functions like `fmt.Fprintf` along with the string builder. 

Property based tests help you exercise them against our code by throwing random data at your code and verifying the rules you describe always hold true. The real challenge about property based tests is having a good understanding of your domain so you can write these properties.

provided - `quick.Check` a function that it will run against a number of random inputs, if the function returns false it will be seen as failing the check.

```go
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num int) bool {
    roman := ConvertToRoman(num)
    fromRoman := ConverttoNum(roman)
	return fromRoman == num
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

```
Given random number (e.g 4).
Call `ConvertToRoman` with random number (should return `IV` if `4`).
Take the result of above and pass it to `ConverttoNum`.
The above should give us our original input (4).

**`uint16`**
Go has types for unsigned integers, which means they cannot be negative; so that rules out one class of bug in our code immediately. By adding 16, it means it is a 16 bit integer which can store a max of 65535, which is still too big but gets us closer to what we need.

``` go
func quick.Check(f interface{}, config *quick.Config) error
```
Check looks for an input to f, any function that returns bool, such that f returns false. It calls f repeatedly, with arbitrary values for each argument. If f returns false on a given input, Check returns that input as a *CheckError. For example:

```go
func TestOddMultipleOfThree(t *testing.T) {

    f := func(x int) bool {

        y := OddMultipleOfThree(x)

        return y%2 == 1 && y%3 == 0
    }
    if err := quick.Check(f, nil); err != nil {

        t.Error(err)

    }

}
```

#### Mathematics

Alas trigonometry and vectors will have some use, can't say "why did I learn this never gonna use" anymore

Go `math` package has both `sin` & `cos`, with one small snag we'll need to get our heads around; if we look at the description of `math.Cos`:

> Cos returns the cosine of the radian argument x.

It wants the angle to be in radians. So what's a radian? Instead of defining the full turn of a circle to be made up of 360 degrees, we define a full turn as being 2π radians.

`math.Sin`:
> Sin returns the sine of the radian argument x

```go
func secondsInRadians (t time.Time) float64 {
  return float64(t.Second()) * (math.Pi / 30)
}

// "go test" Output:
// clockface_test.go:24: Wanted 3.141592653589793 radians, but got 3.1415926535897936
```

Programs using times should typically store and pass them as values, not pointers. That is, time variables and struct fields should be of type `time.Time`, not `*time.Time`

Floating point arithmetic is notoriously inaccurate. Computers can only really handle integers, and rational numbers to some extent. Decimal numbers start to become inaccurate, especially when you factor them up and down as shown in the `secondsInRadians` function above.

Due to rounding errors, most `floating-point` numbers end up being slightly imprecise. As long as this imprecision stays small, it can usually be ignored. However, it also means that numbers expected to be equal (e.g. when calculating the same result through different correct methods) often differ slightly, and a simple equality test fails

When printing a struct, the modified format `%+v` annotates the fields of the structure with their names

Computers often don't like dividing by zero because infinity is a bit strange.

`math.Inf()` - The Inf() function is an inbuilt function of the math package which is used to get the positive infinity or negative infinity. It returns the positive infinity (+Inf) if sign >= 0, the negative infinity (-Inf) if sign < 0. Where the sign is the given parameter.

It accepts a parameter (sign), and returns the positive infinity if sign >= 0, negative infinity if sign < 0.

```go
func Inf(sign int) float64
```
- sign : The value to be used to get the positive or negative infinity. The return type of Inf() function is a float64, it returns the positive infinity if sign >= 0, negative infinity if sign < 0.

### Reading Files

Iterative test driven development requires us to break up our work, but we should be careful not to fall into the trap of taking a `"bottom up"` approach.

In  a `Bottom Up` - approach the individual base elements of the system are first specified in great detail. These elements are then linked together to form larger subsystems, which then in turn are linked, sometimes in many levels, until a complete top-level system is formed. This strategy often resembles a "seed" model, by which the beginnings are small but eventually grow in complexity and completeness.


`Loose coupling` is an approach to interconnecting the components in a system or network so that those components, also called elements, depend on each other to the least extent practicable.

```go
  pacakage blogposts_test
```

By appending `_test` to our intended package name, we only access exported members from our package - just like a real user of our package.

Now, in Go 1.16, there’s a new `io/fs` package that provides a common filesystem interface: `fs.FS`. At first glance, the FS interface is puzzlingly small:

```go
type FS interface {
    Open(name string) (File, error)
}
```

You can read this as “the most atomic type of filesystem is just an object that can open a file at a path, and return a file object”.

An `FS` provides access to a hierarchical file system.

The Go library allows for more complex behavior by providing other filesystem interfaces that can be composed on top of the base `fs.FS` interface, such as `ReadDirFS`, which allows you to list the contents of a directory. The `FS` interface is the minimum implementation required of the file system. `fs.FS` gives us a way of opening a file within it by name with its `Open` method. 

From there we read the data from the file and, for now, we do not need any sophisticated parsing, just cutting out the `Title`: text by slicing the string.

Separating the 'opening file code' from the 'parsing file contents code' will make the code simpler. 

```go
func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile fs.File) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
```

> Does `newPost` have to be coupled to an `fs.File` ? Do we use all the methods
> and data from this type? What do we really need?

In our case we only use it as an argument to `io.ReadAll` which needs an 
`io.Reader`. So we should loosen the coupling in our function and ask for an `io.Reader`.


```go
func ReadAll(r io.Reader) ([]byte, error)
```
`ReadAll` reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because `ReadAll` is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported. 

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends. 

Also, while testing you can use `MapFS` to

```go
func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
    ...}
}
// A MapFS is a simple in-memory file system for use in tests,
// represented as a map from path names to information about 
// the files or directories they represent.
```

`bufio.Scanner`: 
 - Package `bufio` implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O. 
 - `Scanner` provides a convenient interface for reading data such as a file of newline-delimited lines of text. Successive calls to the Scan method will step through the 'tokens' of a file, skipping the bytes between the tokens.

`strings.TrimPrefix`
- Package `strings` implements simple functions to manipulate UTF-8 encoded strings. 
- ```go
    func TrimPrefix(s, prefix string) string
  ```
- ```go
   package main

    import (
        "fmt"
        "strings"
    )

    func main() {
        var s = "¡¡¡Hello, Gophers!!!"
        s = strings.TrimPrefix(s, "¡¡¡Hello, ")
        s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
        fmt.Print(s)
    }

    // Output :
    // Gophers!!!
  ```
  - TrimPrefix returns s without the provided leading prefix string. If `s` doesn't start with prefix, `s` is returned unchanged. 

`scanner.Scan()`: 
- returns a `bool` which indicates whether there's more data to scan, so we can use that with a `for` loop to keep reading through the data until the end.

`strings.TrimSuffix`
- TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
- ```go
    package main

    import (
        "fmt"
        "strings"
    )

    func main() {
        var s = "¡¡¡Hello, Gophers!!!"
        s = strings.TrimSuffix(s, ", Gophers!!!")
        s = strings.TrimSuffix(s, ", Marmots!!!")
        fmt.Print(s)
    }

    // Output:
    // ¡¡¡Hello
  ```

The state of the art for filesystem abstraction (prior to Go 1.16) has been the `afero` library, which contains an interface type for filesystems and a number of common implementations that provide this interface. For example, `afero.OsFs` wraps the os package and `afero.MemMapFs` is an in-memory simulated filesystem that’s useful for testing. Since afero.Fs is just an interface, you can theoretically write any type of client that provides filesystem like behavior (e.g. S3, zip archives, SSHFS, etc.), and use it transparently by anything that acts on an `afero.Fs`.

One downside, similar to the `io `package, is that not all combinations of interface types are covered, so you may need to sprinkle some helper interfaces throughout library code. For example, if I want a `fs.FS` that supports `ReadDir` and `Stat`, I’d need to write my own interface like this:

```go
type readDirStatFS interface {
    fs.ReadDirFS
    fs.StatFS
}
```
There’s one big caveat that you’ll notice if you look at what’s conspicuously absent from the `fs.File` interface: any ability to *write* files. The `fs` package provides a *read-only* interface for filesystems.


So, **to conclude**: out-of-the-box with Go `1.16` you can use `fs.FS` in place of `afero.Fs` for testing and in cases when you’re only performing *read-only* operations

### HTML Templates

Many websites do not need to be a `Single-Page Application`. HTML and CSS are fantastic ways of delivering content and you can use Go to make a website to deliver HTML.

You can generate your HTML in Go with elaborate usage of 
`fmt.Fprintf` (`func Fprintf(w io.Writer, format string, a ...any) (n int, err error)`), but in this chapter you'll learn that Go's standard library has some tools to generate HTML in a simpler and more maintainable way.

We'll design our code so it accepts an io.Writer. This means the caller of our code has the flexibility to:

- Write them to an `os.File` , so they can be statically served
- Write out the HTML directly to a `http.ResponseWriter`
- Or just write them to anything really! So long as it implements `io.Writer`
the user can generate some HTML from a Post

Go has two templating packages `text/template` and `html/template` and they share the same interface. What they both do is allow you to combine a template and some data to produce a string.

Execution of the template walks the structure and sets the cursor, represented by a period '.' and called "dot", to the value at the current location in the structure as execution proceeds.

The input text for a template is UTF-8-encoded text in any format. `"Actions"`--data evaluations or control structures--are delimited by `"{{" and "}}"`

In trim markers, the white space must be present: `"{{- 3}}"` is like `"{{3}}"` but trims the immediately preceding text, while `"{{-3}}"` parses as an action containing the number `-3`.

For example:

```go
"{{23 -}} < {{- 45}}"
// Outputs:

//"23<45"
```

`text/template` example

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}
// Output:
// 17 items are made of wool
```


What's the difference with the HTML version?

`Package template` (html/template) implements data-driven templates for generating HTML output safe against code injection. 

`Package template` implements data-driven templates for generating textual output. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

For Actions - "Arguments" and "pipelines" are evaluations of data.

```
{{range pipeline}} T1 {{end}}
	The value of the pipeline must be an array, slice, map, or channel(pipes that connect concurrent goroutines). If the value of the pipeline has length zero, nothing is output; Otherwise, dot is set to the successive elements of the array, slice, or map and T1 is executed. If the value is a map and the keys are of basic type with a defined order, the elements will be visited in sorted key order.
```

`package/template` example:

```go
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Items}}<div>{{.}}</div>
        {{else}}<div><strong>no rows</strong></div>{{end}}
    </body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "Disney plus Programs",
		Items: []string{
			"The Mandalorian",
			"Avengers Endgame",
			"Lock Stock and Two Smoking Barrels",
		},
	}
	err = t.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "Upcoming Programs",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)

}
// Output: 
// <!DOCTYPE html>
// <html>
//     <head>
//         <meta charset="UTF-8">
//         <title>Disney plus Programs</title>
//     </head>
//     <body>
//         <div>The Mandalorian</div>
//         <div>Avengers Endgame</div>
//         <div>Lock Stock and Two Smoking Barrels</div>
        
//     </body>
// </html>
// <!DOCTYPE html>
// <html>
//     <head>
//         <meta charset="UTF-8">
//         <title>Upcoming Programs</title>
//     </head>
//     <body>
//         <div><strong>no rows</strong></div>
//     </body>
// </html>
```

**However, it is hard to read and any IDE would not have the correct syntax recognition like it would use package embded!**

```go
import "embed"
var (
	//go:embed "templates/*"
	postTemplates embed.FS
    // An FS is a read-only collection of files, usually initialized with 
    //"//go:embed" directive.
)
```

Package embed provides access to files embedded in the running Go program.

Go source files that import "embed" can use the `//go:embed` directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time. 

```go
package server

import "embed"

// content holds our static web server content.
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS
```

A `//go:embed` directive above a variable declaration specifies which files to embed, using one or more path.Match patterns.

The directive must immediately precede a line containing the declaration of a single variable. Only blank lines and ‘//’ line comments are permitted between the directive and the declaration. 

`ApprovalTests` allows for easy testing of larger objects, strings and anything else that can be saved to a file (images, sounds, csv, etc...), the approval tool can compare the output for you with an "approved" file you created.

 - Add a dependency to `"github.com/approvals/go-approval-tests"` to your project and edit the test
 - The first time you run it, it will fail because we haven't approved anything yet
 - It will have created two files, that look like the following
     `renderer_test.TestRender.it_converts_a_single_post_into_HTML.received.txt`
     `renderer_test.TestRender.it_converts_a_single_post_into_HTML.approved.txt`
   The received file has the new, unapproved version of the output. Copy that into the empty approved file and re-run the test.

Store approved files in testdata subfolder:

```go
func TestMain(m *testing.M) {
    //  To configure this, add a call to UseFolder to your TestMain
	UseFolder("testdata")
	os.Exit(m.Run())
}
```

**Benchmarking:**

`b.N` specifies the number of iterations; the value is not fixed, but dynamically allocated, ensuring that the benchmark runs `for` at least one second by `default`.

the reason behind:

`panic: runtime error: invalid memory address or nil pointer dereference` - it's didn't give a memory address to the new variable.To avoid this error, just new the variable, and it will be grant a memory address.

`Render` is a package that provides functionality for easily rendering JSON, XML, text, binary data, and HTML templates.

After adding `RenderIndex` method to our `PostRenderer` that again takes an `io.Writer` and a slice of `Post`

`VerifyString` stores the passed string into the received file and confirms that it matches the approved local file. On failure, it will launch a reporter.

Package reporters provides types to report comparison results.
Reporters launch programs on failure to help you understand, fix and approve results.

```go
type Reporter interface {
	// Report is called when the approved and received file do not match.
	Report(approved, received string) bool
}
```
Reporter are called on failing approvals.

**Passing functions in templates**

Going back to the principles of Mustache and logic-less templates, why did they advocate for logic-less? What is wrong with logic in templates?

Even though the approval tests technique has reduced the cost of maintaining these tests, they're still more expensive to maintain than most unit tests you'll write. They're still sensitive to any minor markup changes you might make, it's just we've made it easier to manage. Instead embrace the idea of `view model`, where you construct specific types that contain the data you need to render, in a way that's convenient for the templating language.

This way, whatever important business logic you use to generate that bag of data can be unit tested separately, away from the messy world of HTML and templating.

A `view model` represents the data that you want to display on your view/page, whether it be used for static text or for input values (like textboxes and dropdown lists) that can be added to the database (or edited). It is something different than your domain model.

Body render had the HTML escaped This is a security feature of Go's html/template package to stop malicious 3rd-party HTML being outputted, to circumvent this `template.HTML` package can be used

```golang
type HTML string
```

HTML encapsulates a known safe HTML document fragment. It should not be used for HTML from a third-party, or HTML with unclosed tags or comments. The outputs of a sound HTML sanitizer and a template escaped by this package are fine for use with HTML.

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

```golang

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
    // parse the Body into HTMLBody and then use that field in the template to render the HTML
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
```

### Generics

The Go compiler expects you to write your functions, structs e.t.c. by describing what types you wish to work with.

You can't pass a `string` to a function that expects an `integer`.
 - Make function implementation simpler. By describing to the compiler what types you work with, you constrain the number of possible valid implementations. You can't "add" a `Person` and a `BankAccount`. You can't `capitalise` an `integer`. In software, constraints are often extremely helpful.

As we may know from previous chapters - Go offers you a way to be more abstract with your types with `interfaces`, so that you can design functions that do not take concrete types but instead, types that offer the behaviour you need.

By using `interface{}` the compiler can't help us when writing our code, because we're not telling it anything useful about the types of things passed to the function. Try comparing two different types.

```golang
func (is *I) Equal(a, b interface{})
//...
AssertNotEqual(1, "1")
// the test compiles but fails with "got 1, want 1" 
```
the error message got 1, want 1 is unclear. Writing functions that take `interface{}` can be extremely challenging and bug-prone because we've lost our constraints, and we have no information at compile time as to what kinds of data we're dealing with.

Generics offer us a way to make abstractions (like interfaces) by letting us **describe our constraints**. They allow us to write functions that have a similar level of flexibility that `interface{}` offers but retain type-safety and provide a better developer experience for callers.

In our case the type of our type parameter is comparable and we've given it the label of T. This label then lets us describe the types for the arguments to our function (got, want T).

```golang
// change from interface to type Parameter
func AssertEqual(t*testing.T, got, want, interface{})
//..
func AssertEqual[T comparable](t *testing.T, got, want T) {}
// change from interface to type Parameter
func AssertEqual(t*testing.T, got, want, interface{})
// ...
func AssertNotEqual[T comparable](t *testing.T, got, want T) {}
```
In our greetings test case we want to use `comparable` with label `T`. This label then lets us describe the types for the arguments to our function (got, want T).

In this case we are using `comparable` because we want to tell the compiler that we wish to use `==` & `!=` operators on things of type `T`.

##### `T any` vs. `interface{}`

```golang
func GenericFoo[T any](x, y T)

func InterfaceyFoo(x, y interface{})
```

In terms of constraints, `any` does mean "anything" and so does `interface{}` but the difference is that with the generic version *you are still describing a specific type*, there is still constraint that this will be work with "one" type.

What this means is you can call `InterfaceyFoo` with any combination of types (e.g `InterfaceyFoo(apple, orange))`. However GenericFoo still offers some constraints because we've said that it only works with one type, T.

```golang
func GenericFoo[T any](x, y T)

// Valid:
// GenericFoo(apple1, apple2)
// GenericFoo(orange1, orange2)
// GenericFoo(1, 2)
// GenericFoo("one", "two")

// Not Valid:
// GenericFoo(apple1, orange1)
// GenericFoo(1, "1")
```

If your function returns the generic type, the caller can also use the type as it was, rather than having to make a type assertion because when a function returns interface{} the compiler cannot make any guarantees about the type.

```golang
func PrintAnything[T any](v T) {}
```
We can read this as saying:
- *For any type T, PrintAnything[T] takes a parameter of type T.*

Stacks are  a collection of items where you can `Push` items to the "top" and to get items back again you `Pop` items from the top (LIFO - last in, first out)

*Last In First Out: It is a method for handling data structures where the first element is processed last and the last element is processed first.* 

try interface routes and see type safety error:

 - We really want to capture the idea of a stack in one type, and have one set of tests for them.
 - ```go
    // stack.go refactored with interface
    type StackOfInts = Stack
    type StackOfStrings = Stack

    type Stack struct {
        values []interface{}
    }
    //...
    // generics_test
    func AssertEqual(t *testing.T, got, want interface{}) {
	    //test block for got != want
    }

    func AssertNotEqual(t *testing.T, got, want interface{}) {
       //test block for got == want
    }
   ```

 - We're aliasing our previous implementations of `StackOfInts` and `StackOfStrings` to a new unified type Stack

 - type safety is removed by by making it so `values` is a `slice` of `interface{}` [for example](https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d)

 - ```go
       t.Run("interface stack dx is horrid", func(t *testing.T) {
        myStackOfInts := new(StackOfInts)

        myStackOfInts.Push(1)
        myStackOfInts.Push(2)
        firstNum, _ := myStackOfInts.Pop()
        secondNum, _ := myStackOfInts.Pop()
        AssertEqual(firstNum+secondNum, 3)
      })
      // Output of "go test": invalid operation: operator + not defined on firstNum (variable of type interface{})
      // needs to be fixed with typeAssertion because compiler doesn't know what the data is
   ```
 - problem with type safety here comes from when `Pop` returns `interface{}` compiler doesn't know what the data is, doesn't let you use the operator `+` because can't know that it should be an `integer`.

 - A type assertion used in an assignment statement or initialization of the special form

 - ```go
    // get our ints from out interface{}
	reallyFirstNum, ok := firstNum.(int)
	AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

	reallySecondNum, ok := secondNum.(int)
	AssertTrue(t, ok) // and again!

	AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
   ```
 - asserts that *firstNum/secondNum* is an `integer` and that the value (in this case `firstNumIsInt` and `secondNumIsInt`) stored in firstNum/seconNum is of type int. The notation `firstNum.(int)`, `secondNum.(int)` in this case is called a type assertion



#### Troubleshooting
- ` go mod init` - initialize go module in your project
- `gopls -rpc.trace -v check ~/file_name.go`
 - ```log
     discovered missing identifiers: map[memRecordCycle:true pageBits:true]
	 package="runtime"
   ```


#### Useful Resources:
1. [GoLang Guide](https://golang.org/doc/)
1. [Static vs. Dynamic](https://hackernoon.com/i-finally-understand-static-vs-dynamic-typing-and-you-will-too-ad0c2bd0acc7)
1. [Variadic Functions](https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085)
1. [GoLang Slices](https://www.callicoder.com/golang-slices/)
1. [Reflections in Go](https://golangbot.com/reflection/)
1. [Slices Usage](https://blog.golang.org/go-slices-usage-and-internals)
1. [Table Driven Tests](https://dave.cheney.net/2019/07/prefer-table-driven-tests)
1. [Interfaces in Go](https://www.alexedwards.net/blog/interfaces-explained)
1. [Bytes & Strings Package](https://medium.com/go-walkthrough/go-walkthrough-bytes-strings-packages-499be9f4b5bd)
1. [Dependency Injection](https://appliedgo.net/di/)
1. [Anonymous Function Loops](https://zknill.io/posts/gos-anonymous-functions-loops/)
1. [Race condition detector](https://blog.golang.org/race-detector)
1. [Concurrency Patters](https://blog.golang.org/pipelines)
1. [Go Language Reflection part 1](https://www.youtube.com/watch?v=oiX7fAmOYX0&t=29s)
1. [Basic types and Value Literals](https://go101.org/article/basic-types-and-value-literals.html)
1. [Modules with vendor support Travis CI](https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/)
1. [Composite literals in go](https://medium.com/golangspec/composite-literals-in-go-10dc62eec06a)
1. [Breakdown of Reflections](https://golangbot.com/reflection/)
1. [Reflections Deep Dive](https://medium.com/swlh/go-reflections-deep-dive-from-structs-and-interfaces-e1931f0c99af)
1. [Interfaces and Reflect](https://blog.gopheracademy.com/advent-2018/interfaces-and-reflect/)
1. [Go Routines](https://golangbot.com/goroutines/)
1. [Basic defer tutorial](https://www.youtube.com/watch?v=aVDkuViaJfY)
1. [Context](https://faiface.github.io/post/context-should-go-away-go2/)
1. [Strings.builder](https://www.calhoun.io/concatenating-and-building-strings-in-go/)
1. [Comparing Floating Point Numbers](https://gopherlabs.kubedaily.com/StandardLib/Comparing_floating_point_numbers.html)
1. [Floating Point Guide](https://floating-point-gui.de/errors/comparison/)
1. [Zero Values in GoLang](https://dave.cheney.net/2013/01/19/what-is-the-zero-value-and-why-is-it-useful)
1. [Math.Inf walkthrough](https://www.includehelp.com/golang/math-inf-function-with-examples.aspx)
1. [Top-down Bottom-Up](https://en.wikipedia.org/wiki/Top-down_and_bottom-up_design)
1. [GoLang workspace folders](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)
1. [Tour of io/fs package](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/)
1. [Text/Template](https://pkg.go.dev/text/template)
1. [How to Use Package Template in Go](https://appdividend.com/2019/11/27/golang-template-example-package-template-in-golang/)
1. [Go Templates](https://zetcode.com/golang/template/)
1. [Approval Tests](https://approvaltests.com/)
1. [Approval Tests Importance](https://understandlegacycode.com/blog/characterization-tests-or-approval-tests/)
1. [GoLang Panic Runtme](https://blog.wuhsun.com/panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference/)
1. [View Model example](https://stackoverflow.com/a/11074506)
1. [Type Parameters](https://bitfieldconsulting.com/golang/type-parameters)
