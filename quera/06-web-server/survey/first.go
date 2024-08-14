package main

import (
	"fmt"
	"sort"
)

func main() {
	var stdInLines int
	fmt.Scanln(&stdInLines)
	flightRateSums := make(map[string]int)
	flightRateCount := make(map[string]int)
	for i := 0; i < stdInLines; i++ {
		var flight string
		fmt.Scanln(&flight)
		flightRateSums[flight] = 0
		flightRateCount[flight] = 0
	}

	fmt.Scanln(&stdInLines)
	ticketFeedbacks := make(map[[2]string]bool)
	for i := 0; i < stdInLines; i++ {
		var person, flight string
		fmt.Scanf("%s %s\n", &person, &flight)
		if _, found := flightRateSums[flight]; !found {
			fmt.Printf("Invalid flight %s\n", flight)
		} else if _, found = ticketFeedbacks[[2]string{flight, person}]; found {
			fmt.Printf("Duplicate ticket for %s %s\n", flight, person)
		} else {
			ticketFeedbacks[[2]string{flight, person}] = false
		}
	}

	fmt.Scanln(&stdInLines)
	for i := 0; i < stdInLines; i++ {
		var person, flight, comment string
		var rate int
		fmt.Scanf("%s %s %d %s\n", &person, &flight, &rate, &comment)
		if _, found := flightRateSums[flight]; !found {
			fmt.Printf("Invalid flight %s\n", flight)
		} else if hasFeedback, found := ticketFeedbacks[[2]string{flight, person}]; !found {
			fmt.Printf("Invalid passenger for %s %s\n", flight, person)
		} else if hasFeedback {
			fmt.Printf("Duplicate comment for %s by %s\n", flight, person)
		} else {
			ticketFeedbacks[[2]string{flight, person}] = true
			flightRateSums[flight] += rate
			flightRateCount[flight]++
			fmt.Printf("Accepted comment for %s by %s\n", flight, person)
		}
	}

	flights := make([]string, len(flightRateCount))
	i := 0
	for flight := range flightRateCount {
		flights[i] = flight
		i++
	}
	sort.Strings(flights)
	for _, flight := range flights {
		if flightRateCount[flight] == 0 {
			continue
		}
		score := float64(flightRateSums[flight]) / float64(flightRateCount[flight])
		fmt.Printf("Average score for %s is %.2f\n", flight, score)
	}
}
