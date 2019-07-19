/* 
 Range lets you iterate over an array. 
 Every time it is called it returns two values, 
 the index and the value. We are choosing to 
 ignore the index value by using _ blank identifier
 It's a bit like writing to the Unix /dev/null file:
*/

package main

func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}