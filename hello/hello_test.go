// %s	the uninterpreted bytes of the string or slice

// For now it's enough to know that your `t` of type `*testing.T` 
// is your "hook" into the testing framework so you can do 
// things like `t.Fail()` when you want to fail

// f in Errorf stands for `format`

// Added assertion into a function. This reduces duplication 
// and improves readability of our tests. In Go you can declare functions 
// inside other functions and assign them to variables.
package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Ali")
        want := "Hello, Ali"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
