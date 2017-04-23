package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	// var score int = 4
	// fmt.Println(score)

	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	var thanos = "I am thanos"
	fmt.Println(thanos)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3. // float64
	fmt.Printf("%v, %T", thorRating, thorRating)
	fmt.Println()

	var ironman, captAmerica string = "I am Ironman", "I am capt America"
	fmt.Println(ironman, captAmerica)

	var defaultValue int // default value for each data types.
	fmt.Println(defaultValue)

	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "I am antman"
	) // standalone variables
	fmt.Println(spiderman, age, powers, antman)

	// User Inputs

	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	var name string = "Anup"
	var len, _ = fmt.Println(name)
	fmt.Println(len)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// name, _ = reader.ReadString('\n')
	// fmt.Println(name)

	// fmt.Print("Enter your rating: ")
	// ratingString, _ := reader.ReadString('\n')
	// ratingFloat, _ := strconv.ParseFloat(strings.TrimSpace(ratingString), 64)
	// fmt.Println(ratingFloat + 2)

}
