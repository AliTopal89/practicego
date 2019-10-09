# practicego

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

an argument of type string - means "an array, 
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

A switch statement is a shorter way to write a sequence of `if - else` statements.
Go only runs the selected case, not all the cases that follow. 
In effect, the break statement that is needed at the end of each case in 
those languages is provided automatically in Go.

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

A "return" statement without arguments returns the named return values.
This is known as a "naked" return.
Naked return statements should be used only in short functions, 
as with the example shown here. They can harm readability in longer functions.
