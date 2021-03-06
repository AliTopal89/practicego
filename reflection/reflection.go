package main

import (
	"fmt"
	"reflect"
)

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

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
			fmt.Println("Key:", i, "Value", val)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
			fmt.Println("Key:", i, "Value", val)
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
			fmt.Println("Key:", key, "Value:", val)
		}

	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		v := val.Call(nil)
		// A nil interface value, which is an interface value that doesn't
		// have an underlying value. This is the zero value of an interface type.
		for _, v := range v {
			walk(v.Interface(), fn)
		}
	}

	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }

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
