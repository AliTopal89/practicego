/* 
Range lets you iterate over an array. 
 Every time it is called it returns two values, 
 the index and the value. We are choosing to 
 ignore the index value by using _ blank identifier
 It's a bit like writing to the Unix /dev/null file:
*/

package arrays

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

/*
Variadic functions can be called with 
any number of trailing arguments. 
For example, fmt.Println is a common variadic function.
variadic function is a function which accepts a variable number of arguments

https://gobyexample.com/variadic-functions
*/

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
//     sums := make([]int, lengthOfNumbers)

//     for i, numbers := range numbersToSum {
//         sums[i] = Sum(numbers)
//     }

//     return sums
// }
// take from 1 to the end" with numbers[1:]
func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        if len (numbers) == 0 {
            sums = append(sums, 0)
        } else {
            tail := numbers[1:]
            sums = append(sums, Sum(tail))
        }
    }

    return sums
}
