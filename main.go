package main

import (
	"fmt"
	"github.com/hokkung/migrator/cmd"
)

func init() {
	fmt.Println("Loading environment variables")
}

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
