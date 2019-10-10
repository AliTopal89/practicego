# practicego  [![Build Status](https://travis-ci.org/AliTopal89/practicego.svg?branch=master)](https://travis-ci.org/AliTopal89/practicego)

## this README and curriculum

is mostly by the work that I followed from:

Chris James [quii](https://github.com/quii/learn-go-with-tests)

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
