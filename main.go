package main

import "fmt"

func main() {
	fmt.Println("Welcome to your Todo List")
	fmt.Printf("Commands\nX + <line number>: delete\nI + <todo item>: insert\nQ: exit\n")
	var input string

	for input != "q" {
		fmt.Print("Action: ")
		fmt.Scan(&input)
		fmt.Println("You entered: ", input)
	}

	fmt.Println("Goodbye")

}
