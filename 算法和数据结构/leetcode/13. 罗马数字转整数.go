package main

import "fmt"

func romanToInt(s string) int {
	value := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	money := 0
	pre_value := 100000
	for _, each := range s {
		current_value := value[string(each)]
		if current_value > pre_value {
			money += current_value
			money -= (pre_value) * 2
		} else {
			money += current_value
		}
		pre_value = current_value
	}

	return money

}

func main() {
	money := romanToInt("MCMXCIV")
	fmt.Println(money)

}
