package main

import ( "fmt"

)

func main () {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
	fmt.Println(a[0], a[1])
	
	primes := [6]int{1,3,5,7,11,13}
	
	fmt.Println(primes)
}

