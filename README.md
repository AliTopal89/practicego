# practicego

## this README and curriculum is mostly by

Chris James [quii](https://github.com/quii/learn-go-with-tests)

## Important notes

```go
import (
    "fmt"
)
```

With import `"fmt"` we are importing a package 
which contains the Println function that we use to print.

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

A "return" statement without arguments returns the named return values.
This is known as a "naked" return.
Naked return statements should be used only in short functions, 
as with the example shown here. They can harm readability in longer functions.