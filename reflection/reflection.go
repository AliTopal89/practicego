package main

import "reflect"

/* The reflect package has a function ValueOf which
   returns us a Value of a given variable. This has ways
   for us to inspect a value, including its fields which
   we use on the next line.
*/
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	/*
	  panic: reflect: call of reflect.Value.NumField on ptr Value
	  In order to resolve this, we can check val.Kind()
	  for a pointer result, and then get the actual value,
	  resolving the pointer:
	*/

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

	/*
	  if it is a slice you can call Len() and Index() on the value
	  to get the len of the slice and element at an index
	*/
}

func getValue(x interface{}) reflect.Value {
	// Get the reflect.Value of x
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

// case reflect.Ptr:
// 	ft := reflect.TypeOf(x)
// 	fv := reflect.ValueOf(ft)
// 	fp := fv.Elem()
// 	walk(field.Interface())

/*
for i := 0; i < val.Len(); i++ {
    walk(val.Index(i).Interface(), fn)
		// val.Index(i).Interface() to reference the actual value.
	}
*/

/* We look at the first and only field, there may be
   no fields at all which would cause a panic,We then call
   `String()` which returns the underlying value as a string
   but we know it would be wrong if the field was something
   other than a string.
*/

func main() {

}
