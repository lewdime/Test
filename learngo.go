package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(val chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	// Passing the random generated value to the val channel
	val <- value
}

func main() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)
	// go routinge will generate a random value
	// and will pass that values to "values" channel
	go CalculateValue(values)
	// value of the "values" channel will be pass a copy
	// to the value variable to be printed out
	value := <-values
	fmt.Println(value)
}
