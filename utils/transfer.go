package utils

func toChar(i int) rune {
	//transfer int to alphabet  1:A 2:B 3C
	return rune('A' - 1 + i)
}