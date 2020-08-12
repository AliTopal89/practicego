package main

import "reflect"

/* The reflect package has a function ValueOf which
   returns us a Value of a given variable. This has ways
   for us to inspect a value, including its fields which
   we use on the next line.
*/
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}

}

/* We look at the first and only field, there may be
   no fields at all which would cause a panic,We then call
   `String()` which returns the underlying value as a string
   but we know it would be wrong if the field was something
   other than a string.
*/

func main() {

}
