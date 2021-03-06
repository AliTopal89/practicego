package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

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

		{
			"Struct with two string fields",
			struct {
				Name      string
				OtherName string
			}{"Chris", "Leroy Jenkins"},
			[]string{"Chris", "Leroy Jenkins"},
		},

		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Franklin",
				Profile{23, "Los Santos"},
			},
			[]string{"Franklin", "Los Santos"},
		},
		{
			"Pointers to things",
			&Person{
				"Franklin",
				Profile{23, "Los Santos"},
			},
			[]string{"Franklin", "Los Santos"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
				{32, "Istanbul"},
			},
			[]string{"London", "Reykjavík", "Istanbul"},
		},
		{
			"Arrays",
			[2]Profile{
				{22, "London"},
				{45, "Istanbul"},
			},
			[]string{"London", "Istanbul"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Qux",
			},
			[]string{"Bar", "Qux"},
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

		t.Run("with maps", func(t *testing.T) {
			aMap := map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			}

			var got []string
			walk(aMap, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "Bar")
			assertContains(t, got, "Boz")
		})

		t.Run("with channels", func(t *testing.T) {
			aChannel := make(chan Profile)

			go func() {
				aChannel <- Profile{33, "Berlin"}
				aChannel <- Profile{34, "Katowice"}
				close(aChannel)
			}()

			var got []string
			want := []string{"Berlin", "Katowice"}

			walk(aChannel, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("with functions", func(t *testing.T) {
			aFunnction := func() (Profile, Profile) {
				return Profile{33, "Namek"}, Profile{34, "Konoha"}
			}

			var got []string
			want := []string{"Namek", "Konoha"}

			walk(aFunnction, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
