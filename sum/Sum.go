package array

func Sum ( array [5]int) int {
	result := 0
	for i:=0; i< 5; i++ {
		result =+ result + array[i]
	} 
	return result
}