package main 


func cool() [10]int {
	result := [10]int{}
	for i := 1; i < 10; i++ {
		result[i-1] = i
	}
	return result
}