package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to your Todo List")
	fmt.Printf("Commands: X + <line number>: delete, I + <todo item>: insert, Q: exit\n")
	var input string

	todoList := loadList()

	printList(todoList)

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Action: ")
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
		fmt.Printf("%v: %v\n", i, todoList[i])
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
