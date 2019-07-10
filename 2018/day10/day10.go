package main

import "fmt"
import ioutil "io/ioutil"

var inputFilePath10 = "/Users/potato/go/src/adventofcode/2018/10inputtest.txt"

func importFile10(inputFilePath string) {
	input, err := ioutil.ReadFile(inputFilePath10)
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
}

func main() {
	importFile10(inputFilePath10)
	fmt.Println("hello")
}
