// import "fmt"

// type order struct {
// 	orderId    int
// 	customerId int
// }

// func main() {
// 	o := order{
// 		orderId:    1234,
// 		customerId: 567,
// 	}
// 	fmt.Println(o)
// }

/*
We need to write a function which will take the struct 'o' in the
program above as an argument and return the following SQL insert query,
*/

// import "fmt"

// type order struct {
// 	orderId    int
// 	customerId int
// }

// func createQuery(o order) string {
// 	i := fmt.Sprintf("insert in to order values (%d, %d)", o.orderId, o.customerId)
// 	return i
// }

// func main() {
// 	o := order{
// 		orderId:    1234,
// 		customerId: 567,
// 	}
// 	fmt.Println(createQuery(o))
// }

/*
The createQuery function creates the insert query by using
the orderId and customerId fields of o. This program will output,

insert into order values(1234, 567)
*/

// import (
// 	"fmt"
// 	"reflect"
// )

// type order struct {
// 	orderId    int
// 	customerId int
// }

// type employee struct {
// 	name    string
// 	id      int
// 	address string
// 	salary  int
// 	country string
// }

// func createQuery(q interface{}) string {
// 	v := reflect.ValueOf(q)
// 	i := fmt.Sprintln("insert in to order values", v)
// 	return i

// }

// func main() {
// 	o := order{
// 		orderId:    1234,
// 		customerId: 567,
// 	}

// 	fmt.Println(createQuery(o))

// 	e := employee{
// 		name:    "Naveen",
// 		id:      565,
// 		address: "Science Park Road, Singapore",
// 		salary:  90000,
// 		country: "Singapore",
// 	}

// 	fmt.Println(createQuery(e))
// }

// import (
// 	"fmt"
// 	"reflect"
// )

// type order struct {
// 	orderId    int
// 	customerId int
// }

// type employee struct {
// 	name    string
// 	id      int
// 	address string
// 	salary  int
// 	country string
// }

// func createQuery(q interface{}) {
// 	if reflect.ValueOf(q).Kind() == reflect.Struct {
// 		v := reflect.ValueOf(q)
// 		fmt.Println("Number of fields", v.NumField())
// 		for i := 0; i < v.NumField(); i++ {
// 			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
// 		}
// 	}

// }

/*
Type represents the actual type of the interface{},
in this case main.Order and Kind represents the specific kind of the type.
In this case, it's a struct.
*/

// func main() {
// 	o := order{
// 		orderId:    1234,
// 		customerId: 567,
// 	}

// 	createQuery(o)

// 	e := employee{
// 		name:    "Naveen",
// 		id:      565,
// 		address: "Science Park Road, Singapore",
// 		salary:  90000,
// 		country: "Singapore",
// 	}

// 	createQuery(e)
// }

/* Outputs:
Number of fields 2
Field:0 type:reflect.Value value:1234
Field:1 type:reflect.Value value:567
Number of fields 5
Field:0 type:reflect.Value value:Naveen
Field:1 type:reflect.Value value:565
Field:2 type:reflect.Value value:Science Park Road, Singapore
Field:3 type:reflect.Value value:90000
Field:4 type:reflect.Value value:Singapore

*/

package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("Insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
			case reflect.String:
				query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}

}

func main() {
	o := order{
		orderId:    1234,
		customerId: 567,
	}

	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Science Park Road, Singapore",
		salary:  90000,
		country: "Singapore",
	}

	createQuery(e)

}

/*
we first check whether the passed argument is a struct.
In line no. 181 we get the name of the struct from its `reflect.Type`
using the `Name()`` method. In the next line, we use `t` and start creating the query.
The case statement in line. 186 checks whether the current field is reflect.Int,
if that's the case we extract the value of that field as `int64` using the `Int()`` method.
Similar logic is used to extract the string in line no. 188.
*/
