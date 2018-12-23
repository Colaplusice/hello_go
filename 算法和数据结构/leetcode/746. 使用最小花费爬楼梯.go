package main

import "fmt"

func min(value1 int, value2 int) int {
	if value1 < value2 {
		return value1
	}
	return value2
}
func minCostClimbingStairs(cost []int) int {
	result := make([]int, len(cost))

	result[0] = cost[0]
	result[1] = cost[1]
	for index := 2; index < len(cost); index++ {
		if index+1==len(result){
		result[index] = min(result[index-2]+cost[index], result[index-1])
		}else{
		result[index] = min(result[index-2]+cost[index], result[index-1]+cost[index])
		}
	}
	return result[len(cost)-1]
}
func main() {
	coast := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res:=minCostClimbingStairs(coast)
	fmt.Println(res)
}
