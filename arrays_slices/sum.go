/* 
Range lets you iterate over an array. 
 Every time it is called it returns two values, 
 the index and the value. We are choosing to 
 ignore the index value by using _ blank identifier
 It's a bit like writing to the Unix /dev/null file:
*/

package main

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

https://gobyexample.com/variadic-functions
*/

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numbersToSum {
        sums[i] = Sum(numbers)
    }

    return sums
}
