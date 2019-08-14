// material followed for this tutorial https://www.callicoder.com/golang-slices/

package arrays
import "fmt"

func cities() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indianCities := asianCities[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indianCities = ", indianCities)
}
