package main

func Repeate( letter string, count int) string {
	var repeated string
	for i :=0; i < count; i++ {
		repeated = repeated + letter
	}
	return repeated
}