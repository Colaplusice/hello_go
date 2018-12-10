package algorithm


const MaxArraySize = 100

func preMp(x string) [MaxArraySize] int {
	var i, j int
	length := len(x) - 1
	var mpNext [MaxArraySize]int
	i = 0
	j = -1
	mpNext[0] = -1
	for i < length {
		// 往前移动
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}
