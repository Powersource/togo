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
	todos, err := yaml.Marshal("%} \n-\ntest")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshalled:", todos)
	var text string
	err = yaml.Unmarshal(todos, &text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshalled:", text)

}
