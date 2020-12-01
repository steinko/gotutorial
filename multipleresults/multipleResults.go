package main

import "fmt"

func Swap (a, b string) (string, string) {
     return b , a
}

func swaping () {
	fmt.Println(Swap("Hello", "World"))
   
}