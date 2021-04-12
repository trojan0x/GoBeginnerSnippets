package main

import "fmt"

func main() {
	fmt.Printf("What's your name? : ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("How old are you? : ")
	var age int
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println(name, "You are too young to continue! ")

	} else {
		fmt.Println("Nice to see you", name)
	}
}
