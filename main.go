package main

import (
	"fmt"
	"os"
	"strings"
)
func main()  {
inputfile := "sample.txt"
outputfile := "result.txt"
input, err := os.ReadFile(inputfile)
	if err != nil{
		fmt.Println("err")
		os.Exit(1)
	}
err = os.WriteFile(outputfile, input, 0644)
	fmt.Println("Successful.........")
	word := strings.Fields("hex")
	for i := 0; i < len(word); i++ {	
	}
}