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