package main

import "fmt"

func main() {
	const my_mars_weight = 125 * 0.3783
	const my_age = 2022 - 1994
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(my_mars_weight)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(my_age)
	fmt.Print(" years old.\n")
	fmt.Println("My weight on the surface of Mars is", my_mars_weight, "lbs, and I would be", my_age, "years old.")
}
