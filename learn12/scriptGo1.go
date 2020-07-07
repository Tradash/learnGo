package main

import "os"

func main() {
	println("Запущен... ", os.Args[0])
	env := os.Environ()
	for i, v := range env {
		println(i, v)
	}
}
