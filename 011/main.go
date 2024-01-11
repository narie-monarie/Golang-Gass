package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	content := "Hello from Golang"
	file, err := os.Create("./files/text.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters", length)
	defer file.Close()

	poodle := Dog{"poodle", 10, "Woof"}
	fmt.Printf("%+v\n ", poodle)

	poodle.Speak()

	var a string
	fmt.Scanln(&a)
	fmt.Scanln("The number is: ")
	aint, err := strconv.ParseInt(strings.TrimSpace(a), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aint)
	}

	//new() -> Allocates but does not initialize memory
	//make() -> allocates and initializes memory

	anInt := 42
	//p points to the memory adress of anInt
	p := &anInt
	*p = 12

	println(anInt)

}
