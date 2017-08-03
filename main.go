package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var fileName = "togo.yaml"

type taskFile struct {
	Version string
	Tasks   taskList
}

type taskList []string

func main() {
	tasks := taskList{
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

	for _, el := range readTogoFile() {
		fmt.Println(el)
	}
}

func readTogoFile() taskList {
	fileContents, err := ioutil.ReadFile(fileName)
	check(err)

	var readTodos taskFile
	err = yaml.Unmarshal(fileContents, &readTodos)
	check(err)

	return readTodos.Tasks
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
