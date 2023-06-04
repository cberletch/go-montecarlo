package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const life_expectancy = 85

func GetInt(x string) int {
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

func GetRisk() (float32, float32) {
	var input int
	var mean float32
	var stddev float32

	fmt.Println("Select risk level from the below options:")
	fmt.Println("1. Aggresive - can accept losses upwards of 40%")
	fmt.Println("2. Semi-aggresive - can accept losses upwards of 25%")
	fmt.Println("3. Neutral - can accept losses upwards of 15%")
	fmt.Println("4. Semi-averse - can accept losses upwards of 10%")
	fmt.Println("5. Risk-averse - can accept losses upwards of 5% ")

	for {
		fmt.Print("Enter a number between 1 and 5: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input. Please Enter your selection as an integer between 1 and 5.")
			continue
		}
		if input < 1 || input > 5 {
			fmt.Println("Invalid input. Please Enter your selection as an integer between 1 and 5.")
			continue
		}
		break
	}

	if input == 5 {
		mean = 10.0
		stddev = 17.9
	}

	fmt.Println(input)

	return mean, stddev
}

func main() {
	var age int
	var assets float32
	var retirement_age int
	var burn_rate float32
	var mean float32
	var stddev float32

	age = 65 //GetInt("age")
	fmt.Println(age)

	mean, stddev = GetRisk()
	fmt.Println(mean, stddev)

	assets = 100000 //GetFloat("assets")
	fmt.Println(assets)

	retirement_age = 65 //GetInt("retirement age")
	fmt.Println(retirement_age)

	burn_rate = .04 //GetFloat("burn rate")
	fmt.Println(burn_rate)

	for i := 0; i < life_expectancy-age; i++ {

		// add logic to edit mean and std. dev based on risk tolerance later

		if age >= retirement_age {

			rand.Seed(time.Now().UnixNano() + int64(i))
			randomFloat := rand.NormFloat64() * float64(stddev+mean)
			randomFloat = math.Max(-38.5, math.Min(48.9, randomFloat))

			assets = assets * (1.0 + float32(randomFloat/100.0))
			fmt.Println(int(assets))
		}

		rand.Seed(time.Now().UnixNano() + int64(i))

		randomFloat := rand.NormFloat64() * float64(stddev+mean)
		randomFloat = math.Max(-38.5, math.Min(48.9, randomFloat))
		assets = assets * (1.0 + float32(randomFloat/100.0))
		age++
	}

}
