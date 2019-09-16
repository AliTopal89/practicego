package split

import (
	"testing"
	"reflect"
)

/*
  for *reflect.DeepEquals* Map case, you'll see that 
  it first checks if both maps are nil, then it 
  checks if they have the same length before finally 
  checking to see if they have the same set of (key, value) pairs.

  As DeepEqual traverses the data values it may find a cycle. 
  The second and subsequent times that DeepEqual compares two 
  pointer values that have been compared before, it treats the 
  values as equal rather than examining the values to which they point. 
  This ensures that DeepEqual terminates.
*/
func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a","b","c"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestSplitWrongSep(t *testing.T){
	got := Split("a/b/c", ",")
	want := []string{"a/b/c"}
	if !reflect.DeepEqual(got, want){
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

/* There is a lot of duplication in our tests. 
   For each test case only the input, 
   the expected output, and name of the 
   test case change.
*/

func TestSplitNoSep(t *testing.T){
	got := Split("abc", "/") 
	want := []string{"abc"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

/*
  To set up all the inputs and expected outputs and 
  feed them to a single test harness, Its time 
  to introduce table driven testing. 
 
  We can even use an anonymous 
  struct literal to reduce the boilerplate
*/
func TdTestSplit(t *testing.T){
	tests := []struct{
		input string
		s string //seperation string
		want []string

	}{
		{input: "a/b/c", s: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", s: ",", want: []string{"a/b/c"}},
		{input: "abc", s: "/", want: []string{"abc"}},
	}
	
	/*  
	   _, ~ It avoids having to declare all the variables 
	   for the returns values. You can see it in a loop too.

	   If you only need the second item in the range (the value), 
	   use the blank identifier, an underscore, to discard the first:

	   'sum := 0
		for _, value := range array {
    		sum += value
		}'
	*/

	for _, tc := range tests { // tc - test cases
		got := Split(tc.input, tc.s)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("want: %v, got : %v", tc.want, got)
		}	
	}
}