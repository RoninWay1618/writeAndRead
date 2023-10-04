package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
	age  int
}

func main() {
	people := []person{
		{"Tom", 25},
		{"Mot", 37},
		{"Pop", 22},
	}

	fileName := "file1940.txt"
	writeData(fileName, people)
	readData(fileName, &people)
}

func writeData(fileName string, people []person) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, p := range people {
		fmt.Fprintf(file, "%s - %d\n", p.name, p.age)
	}

}

func readData(fileName string, people *[]person) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		var p person
		_, err = fmt.Fscanf(file, "%s - %d\n", &p.name, &p.age)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)

		}
	}
	for _, p := range *people {
		fmt.Printf("name: %s, age: %d\n", p.name, p.age)
	}
}
