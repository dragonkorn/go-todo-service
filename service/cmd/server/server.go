package main

import (
	container "service/internal/container"
)

func main() {
	println("hello world")
	c, err := container.NewContainer()
	if err != nil {
		panic(err)
	}

	err = c.Run()
	if err != nil {
		panic(err)
	}
}
