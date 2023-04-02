package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distance = 62100000
	const secondsPerDay = 86400
	var spaceLines = [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}

	fmt.Println("SpaceLine        Days Trip type Price")
	fmt.Println("=====================================")

	count := 0
	for count < 10 {
		roundTrip := rand.Intn(2) == 0
		speed := rand.Intn(15) + 16

		cost := speed + 20
		duration := distance / speed / secondsPerDay

		var tripType string
		if roundTrip {
			tripType = "Round-trip"
			cost = cost * 2
		} else {
			tripType = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v \n", spaceLines[rand.Intn(3)], duration, tripType, cost)
		count++
	}
}
