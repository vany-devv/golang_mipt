package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}
	fmt.Println("Имя пользователя:", user)

	args := os.Args[1:]
	fmt.Println("Аргументы CLI:", args)

	fmt.Println("Версия Go:", runtime.Version())
}
