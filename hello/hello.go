// With import "fmt" we are importing a package 
// which contains the Println function that we use to print.
package main

import "fmt"

func Hello(name string) string {
    return "Hello, " + name
}

func main() {
    fmt.Println(Hello("world"))
}

