package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	// Decode arbitrary JSON into en empty interface
	var input interface{}
	err = json.Unmarshal(bytes, &input)
	if err != nil {
		panic(err)
	}
	log.Printf("Input: %v", input)
	tmpl, err := template.New("Test").Parse("My name is {{.name}}\nContent: {{index .array 1}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, input)
	if err != nil {
		panic(err)
	}
}
