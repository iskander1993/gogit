package main

import "fmt"

func sayBye(name string) {
	fmt.Println("Goodbye", name)
}

func greeting(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greeting("Iskander")
	fmt.Println("You are AWESOME!!!")
	sayBye("Salary")
}
