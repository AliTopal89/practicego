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

In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn't exist, i is the value type's zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.

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

Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.

```go
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.FPrintf(w, "hello/n")
}

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":8090", nil)
}
```
A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handler interface. A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.

Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response. Here our simple response is just “hello\n”. We register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in the net/http package and takes a function as an argument.

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
we have this bytes buffer and we can use `b.Write` and thas fine because even if you look at implementation documentation you can see att all its a pointer to a buffer is the receiver for that write method, but the reason this will work when we are talking about methods, if you remember `b.Write` is shorthand for enclosing b and taking the address of it and then calling a write method on it, so we can do that. The reason why line `fmt.Fprintf(b,..)` isn't going to works is because now you are passing it to a function a copy of that buffer, and inside that function it wants to take a pointer, so you can get it working by passing a pointer in our buffer `fmt.Fprintf(&b), "Holy %s\n"`

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

The Go == operator compares not just the time instant but also the Location and the *monotonic clock*(is for measuring time) reading.
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

Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that the're declared - this is what the () at the end of the anonymous function is doing.Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

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

All it's really saying is it needs a function that takes a ResponseWriter and a Request, which is not too surprising for an HTTP server.

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


#### Useful Resources:
1. [GoLang Guide](https://golang.org/doc/)
1. [Static vs. Dynamic](https://hackernoon.com/i-finally-understand-static-vs-dynamic-typing-and-you-will-too-ad0c2bd0acc7)
1. [Variadic Functions](https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085)
1. [GoLang Slices](https://www.callicoder.com/golang-slices/)
1. [Reflections in Go](https://golangbot.com/reflection/)
1. [Slices Usage](https://blog.golang.org/go-slices-usage-and-internals)
1. [Table Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
1. [Interfaces in Go](https://www.alexedwards.net/blog/interfaces-explained)
1. [Bytes & Strings Package](https://medium.com/go-walkthrough/go-walkthrough-bytes-strings-packages-499be9f4b5bd)
1. [Dependency Injection](https://appliedgo.net/di/)
1. [Anonymous Function Loops](https://zknill.io/posts/gos-anonymous-functions-loops/)
1. [Race condition detector](https://blog.golang.org/race-detector)
1. [Concurrency Patters](https://blog.golang.org/pipelines)