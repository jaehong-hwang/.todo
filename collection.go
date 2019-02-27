package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// TodoCollection is manage .todo filesystem
type TodoCollection struct {
	file string
}

// Init todo collection directory
func (t *TodoCollection) Init() error {
	if t.file != "/" {
		return errors.New("todo collection already exists")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	f, err := os.Create(dir + "/" + TodoFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

// Add todo item
func (t *TodoCollection) Add() error {
	input, err := ioutil.ReadFile(t.file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(t.file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
