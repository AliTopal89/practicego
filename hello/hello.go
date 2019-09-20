package main

import (
    "fmt"
)

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

