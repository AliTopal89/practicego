// %s	the uninterpreted bytes of the string or slice

// For now it's enough to know that your `t` of type `*testing.T` 
// is your "hook" into the testing framework so you can do 
// things like `t.Fail()` when you want to fail

// f in Errorf stands for `format`
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Ali")
    want := "Hello, Ali"

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}
