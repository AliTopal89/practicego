package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	// blank identifier so you dont declare any variable at all
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			/*
			  Array values are deeply equal when their corresponding elements are deeply equal.
			  Struct values are deeply equal if their corresponding fields, both exported and unexported, are deeply equal.
			  Func values are deeply equal if both are nil; otherwise they are not deeply equal.
			  Interface values are deeply equal if they hold deeply equal concrete values.
			*/

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	// expected := "Chris"
	// var got []string

	// // concrete type ~ struct
	// x := struct {
	// 	Name string
	// }{expected}

	// walk(x, func(input string) {
	// 	got = append(got, input)
	// })

	// if len(got) != 1 {
	// 	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	// }
	// if got[0] != expected {
	// 	t.Errorf("got %q, want %q", got[0], expected)
	// }
}
