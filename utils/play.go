package utils

import (
	"errors"
	"fmt"
)

func main() {
	//a := []int{1, 2, 3}
	_, err := Id(2)
	if err != nil {
		fmt.Println("not ok")
	}
}

func Id(nums int) (float64, error) {
	if nums > 1 {
		return 23, errors.New("math length < nums")
	}
	return 1, nil
}
