package main

import "fmt"

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		num := i
		in_num := true
		new_num := 0
		for num > 0 {
			new_num *= 10
			if a := is_in(num % 10); a != -1 {
				new_num += a
			} else {
				in_num = false
				break
			}
			num /= 10
		}
		if in_num && new_num != i {
			fmt.Println(i)
			count += 1
		}

	}
	return count
}
func is_in(num int) int {
	num_dict := map[int]int{2: 5, 6: 9, 1: 1, 8: 8, 5: 2, 9: 6, 0: 0}
	if value, ok := num_dict[num]; ok {
		return value
	}
	return -1
}

func main() {
	result := rotatedDigits(10)
	fmt.Println(result)

}
