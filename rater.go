package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Front end
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name: ")
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Println("Please rate from 1 to 5: ")
	userRatingString, _ := reader.ReadString('\n')
	userRatingNumber, _ := strconv.ParseFloat(strings.TrimSpace(userRatingString), 64)

	// Back end
	fmt.Printf("Hello %v, \n Your rating is %v star. \n Rating recorded at %v \n", userName, userRatingNumber, time.Now().Format(time.Stamp))

	if userRatingNumber == 5 {
		fmt.Println("Good rating")
	} else if userRatingNumber == 4 || userRatingNumber == 3 {
		fmt.Println("Average rating")
	} else {
		fmt.Println("Bad rating")
	}
}
