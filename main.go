package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const life_expectancy = 85

func GetInput(x string) int {
	var val int

	fmt.Println("Enter", x, ":")
	fmt.Scanln(&val)

	return val
}

func GetFloat(x string) float32 {
	var val float32

	fmt.Println("Enter", x, ":")
	fmt.Scanln(&val)

	return val
}

func main() {
	/*var age int
	var assets int
	var retirement_age int
	var burn_rate float32

	age = GetInput("age")
	fmt.Println(age)

	assets = GetInput("assets")
	fmt.Println(assets)

	retirement_age = GetInput("retirement age")
	fmt.Println(retirement_age)

	burn_rate = GetFloat("burn rate")
	fmt.Println(burn_rate)*/

	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		mean := 0.0
		stddev := 20.0
		randomFloat := rand.NormFloat64()*stddev + mean
		randomFloat = math.Max(-50, math.Min(50, randomFloat))
		fmt.Println(randomFloat)
	}

}
