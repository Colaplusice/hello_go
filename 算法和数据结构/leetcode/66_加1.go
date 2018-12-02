package main

func plusOne(digits []int) []int {
	length := len(digits)
	add_next := 1
	for i := length - 1; i >= 0; i-- {
		if (digits[i] == 9 && add_next == 1) {
			digits[i] = 0
			add_next = 1
			if i == 0 {
				//append(digits, 1)
			}
		} else if digits[i] != 9 {
			if add_next == 1 {
				digits[i] += 1
				add_next = 0
			}
		} else {
			digits[i] += 1
			return digits
		}
	}
	println(digits)
	return digits
}

func main() {
	a := []int{0}
	result := plusOne(a)
	print(result[0])

}
