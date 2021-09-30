package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// accept user input
	fmt.Print("What's the biggest number: ")
	var number int
	fmt.Scanf("%d", &number)

	// some intro
	fmt.Println("You entered:", number)
	fmt.Println("Guess a number between 1 and", number)

	// generate a random number
	var random int
	randomNumber(&random, number)

	// loop until user guesses correctly
	// and record the number of guesses
	var guesses int
	
	for {
		fmt.Print("Guess a number: ")
		var guess int
		fmt.Scan(&guess)
		guesses++
		if guess == random {
			fmt.Println("You guessed correctly!")
			fmt.Println("Number of guesses:", guesses)
			break
		} else if guess > random {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
}

func randomNumber(random *int, number int) {
	// generate new seed using current time
	rand.Seed(time.Now().UnixNano())

	// random a number between 1 and 10
	*random = rand.Intn(number) + 1
}

// UnixNano explanation:
// UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC. The result is undefined if the Unix time in nanoseconds cannot be represented by an int64 (a date before the year 1678 or after 2262). Note that this means the result of calling UnixNano on the zero Time is undefined. The result does not depend on the location associated with t.