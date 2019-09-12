package split
 
import (
    "strings"
)

/*
Split slices s into all substrings separated by sep and returns a slice of
the substrings between those separators.

If s does not contain sep and sep is not empty, Split returns a
slice of length 1 whose only element is s.

If sep is empty, Split splits after each UTF-8 sequence. If both s
and sep are empty, Split returns an empty slice.

It is equivalent to SplitN with a count of -1.
*/
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}