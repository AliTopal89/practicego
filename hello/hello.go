// With import "fmt" we are importing a package 
// which contains the Println function that we use to print.
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world"))
}

