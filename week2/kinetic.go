package main

import (
	"fmt"
	"math"
	"strconv"
)

func GenDisplaceFn(a, vo, so *float64) func(float64) float64 {
	return func(t float64) float64 {
		return *a/2*math.Pow(t, 2) + *vo*t + *so
	}
}

func parseInput(f *string) *float64 {
	parsedValue, err := strconv.ParseFloat(*f, 64);
	if  err != nil {
		panic(err)
	}
	return &parsedValue
}

func main () {
	var a, vo, so, t string

	fmt.Printf("Pls provide value for the acceleration: ")
	fmt.Scan(&a)

	fmt.Printf("Pls provide value for the initial velocity: ")
	fmt.Scan(&vo)

	fmt.Printf("Pls provide value for initial displacement: ")
	fmt.Scan(&so)

	fmt.Printf("Pls provide value for the time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(parseInput(&a), parseInput(&vo), parseInput(&so))

	fmt.Printf("after your provided time (in seconds) %f \n", fn(*parseInput(&t)))
	fmt.Printf("after 3 seconds %f \n", fn(3))
	fmt.Printf("after 5 seconds %f \n", fn(5))
}