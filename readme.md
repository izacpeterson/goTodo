# Simple Todo List Application

This is a simple command-line Todo List application written in Go. It allows you to manage your tasks by adding, deleting, and listing them. Tasks are saved to a file so they persist between runs.

## Features

- **Add tasks**: Insert a new task into your todo list.
- **Delete tasks**: Remove a task from your todo list by its line number.
- **List tasks**: Display all tasks in your todo list.
- **Persistent storage**: Tasks are saved to a file and loaded on startup.

## Commands

- **I + <todo item>**: Insert a new task.
- **X + <line number>**: Delete a task by its line number.
- **L**: List all tasks.
- **Q**: Exit the application.

## Getting Started

### Prerequisites

- Go (https://golang.org/doc/install)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/todo-list.git
   ```
2. Change to the project directory:
   ```sh
   cd todo-list
   ```
3. Build the application:
   ```sh
   go build
   ```

### Usage

1. Run the application:
   ```sh
   ./todo-list
   ```
2. Follow the on-screen instructions to manage your tasks.

## Code Explanation

### main.go

This is the main file containing the logic for the Todo List application.

- `main()`: The entry point of the application. It loads the todo list, handles user input, and performs actions based on commands.
- `loadList()`: Loads the todo list from a file.
- `printList()`: Prints the current todo list to the console.
- `saveList()`: Saves the current todo list to a file.

## Example

```sh
Welcome to your Todo List
Commands: X + <line number>: delete, I + <todo item>: insert, L: List tasks Q: exit

1: Buy groceries
2: Call Alice
3: Finish report

Action (i/x/l/q): I Clean the house
1: Buy groceries
2: Call Alice
3: Finish report
4: Clean the house

Action (i/x/l/q): X 2
1: Buy groceries
2: Finish report
3: Clean the house

Action (i/x/l/q): L
1: Buy groceries
2: Finish report
3: Clean the house

Action (i/x/l/q): Q
Goodbye

```
