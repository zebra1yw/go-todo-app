# go-todo-app

Go-Todo-App is a simple and efficient task management application built with Go. This application allows you to create, complete, delete, and view tasks in a user-friendly manner. 

# Preview
<img width="744" alt="Screenshot 2023-11-27 at 0 31 01" src="https://github.com/zebra1yw/go-todo-app/assets/151030350/f01d5f3c-61e3-48d6-ae67-fdcb6522ea87">

## Getting Started
To get started with Go-Todo-App, clone the repository and follow the instructions in the README.md file.

## Installation
To install it, simply run go get (you have to install go first):
```
git clone this repo
cd dir
go build ./cmd/todo
./todo your command flag
```

## Features
- **Create Tasks**: Add new tasks with a name.
  ```
  ./todo -add SampleTask
  ```
- **Complete Tasks**: Complete existing tasks.
  ```
  ./todo -complete=1
  ```
- **Delete Tasks**: Remove tasks that are no longer needed.
   ```
  ./todo -delete=1
  ```
- **View Tasks**: Display all tasks in an organized list.
  ```
  ./todo -list
  ```
## External Libraries
| Library Name | Link | License | 
|---|---|---|
| simpletable | https://github.com/alexeyco/simpletable| MIT license|

## Contributions
Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## License
Distributed under the MIT License. See `LICENSE` for more information.
