package main

import "reflect"

/* The reflect package has a function ValueOf which
   returns us a Value of a given variable. This has ways
   for us to inspect a value, including its fields which
   we use on the next line.
*/
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	/*
	  panic: reflect: call of reflect.Value.NumField on ptr Value
	  In order to resolve this, we can check val.Kind()
	  for a pointer result, and then get the actual value,
	  resolving the pointer:
	*/

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
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

// case reflect.Ptr:
// 	ft := reflect.TypeOf(x)
// 	fv := reflect.ValueOf(ft)
// 	fp := fv.Elem()
// 	walk(field.Interface())

/* We look at the first and only field, there may be
   no fields at all which would cause a panic,We then call
   `String()` which returns the underlying value as a string
   but we know it would be wrong if the field was something
   other than a string.
*/

func main() {

}
