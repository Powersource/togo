package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type TaskFile struct {
	Version string
	Tasks   []string
}

func main() {
	todos, err := yaml.Marshal("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)

}
