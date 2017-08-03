package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var fileName = "togo.yaml"

type taskFile struct {
	Version string
	Tasks   []string
}

func main() {
	tasks := []string{
		"water the plants",
		"buy groceries",
		"run bsd",
	}
	defaultTogo := taskFile{
		Version: "0.1",
		Tasks:   tasks,
	}

	todos, err := yaml.Marshal(defaultTogo)
	check(err)
	fmt.Println("Marshalled:", string(todos))

	err = ioutil.WriteFile(fileName, todos, 0644)
	check(err)

	fileContents, err := ioutil.ReadFile(fileName)
	check(err)

	var readTodos taskFile
	err = yaml.Unmarshal(fileContents, &readTodos)
	check(err)
	fmt.Println("Unmarshalled:", readTodos.Tasks[0])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
