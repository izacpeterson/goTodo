package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to your Todo List")
	fmt.Printf("Commands: X + <line number>: delete, I + <todo item>: insert, L: List tasks Q: exit\n")
	var input string

	todoList := loadList()

	printList(todoList)

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Action (i/x/l/q): ")
		scanner.Scan()
		input = scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}

		action := input[0]
		content := strings.TrimSpace(input[1:])

		switch action {
		case 'i', 'I':
			todoList = append(todoList, content)
			printList(todoList)
			saveList(todoList)

		case 'x', 'X':
			index, err := strconv.Atoi(content)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			if index > len(todoList) {
				fmt.Println("Invalid selection")
				break
			}
			fmt.Printf("Delete: #%v, %v\n", content, todoList[index-1])

		case 'l', 'L':
			printList(todoList)
		default:
			fmt.Println("Invalid Entry")
		}

	}

	fmt.Println("Goodbye")

}

func loadList() []string {
	file, err := os.Open("tasks.tsk")
	if err != nil {
		fmt.Println("ERROR Opening File: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var todoList []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		todoList = append(todoList, scanner.Text())
	}

	return todoList
}

func printList(todoList []string) {
	println()
	for i := 0; i < len(todoList); i++ {
		fmt.Printf("%v: %v\n", i+1, todoList[i])
	}
	println()
}

func saveList(todoList []string) {
	file, err := os.Create("tasks.tsk")
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, item := range todoList {
		fmt.Fprintln(writer, item)
	}

	writer.Flush()
}
