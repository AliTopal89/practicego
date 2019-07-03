// With import "fmt" we are importing a package 
// which contains the Println function that we use to print.

// variable Prefix will be assigned the "zero" value. 
// This depends on the type, for example ints are 0 and for strings it is "".

// A "return" statement without arguments returns the named return values.
// This is known as a "naked" return.
// Naked return statements should be used only in short functions, 
// as with the example shown here. They can harm readability in longer functions.

package main

import "fmt"

const turkish = "Turkish"
const french = "French"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const turkishHelloPrefix = "Napiyon, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

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

func main() {
    fmt.Println(Hello("world", ""))
}

