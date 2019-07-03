// With import "fmt" we are importing a package 
// which contains the Println function that we use to print.
package main

import "fmt"

const turkish = "Turkish"
const turkishHelloPrefix = "Napiyon, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == turkish {
    	return turkishHelloPrefix + name
    }
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}

