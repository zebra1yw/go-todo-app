package main

import (
	"bufio"
	"flag"
	"fmt"
	go_todo_app "github.com/zebra1yw/go-todo-app"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	completed := flag.Int("complete", 0, "mark todo as completed")
	del := flag.Int("delete", 0, "delete todo")
	list := flag.Bool("list", false, "list all todos")
	flag.Parse()
	todos := &go_todo_app.Todos{}
	if err := todos.Read(todoFile); err != nil {
		fmt.Print(fmt.Errorf("can't read the file"))
	}
	switch {
	case *add:
		task, err := readInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Print(fmt.Errorf("can't read the todo name"))
		}
		todos.Add(task)
		err = todos.Save(todoFile)
		if err != nil {
			fmt.Print(fmt.Errorf("can't add the todo"))
		}
	case *completed > 0:
		err := todos.Complete(*completed)
		if err != nil {
			fmt.Print(fmt.Errorf("can't complete the todo"))
		} else {
			err := todos.Save(todoFile)
			if err != nil {
				fmt.Print(fmt.Errorf("can't save the completed todo"))
			}
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Print(fmt.Errorf("can't delete the todo"))
		} else {
			err := todos.Save(todoFile)
			if err != nil {
				fmt.Print(fmt.Errorf("can't save the deleted todo"))
			}
		}
	case *list:
		todos.Print()
	default:
		fmt.Print(fmt.Errorf("invalid command"))
	}

}

func readInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", fmt.Errorf("empty todo is not allowed")
	}
	return text, nil
}
