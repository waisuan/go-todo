package main

func main() {
	todoList := NewTodoList()
	r := SetupRouter(todoList)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
