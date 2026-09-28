package main

import (
	"fmt"
	"os"
)

func main()  {
inputfile := "sample.txt"
outputfile := "result.txt"

input, err := os.ReadFile(inputfile)
	if err != nil{
		fmt.Println("err")
	}
err = os.WriteFile(outputfile, input, 0644)
}