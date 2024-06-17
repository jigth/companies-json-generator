package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	jsongenerator "github.com/jigth/companies-json-generator/src/json-generator"
)

func main() {
	outputPath := os.Args[1]

	companiesAmount := readUserInput()

	companies, err := jsongenerator.CreateFakeCompanies(companiesAmount)
	if err != nil {
		fmt.Println(err)
		panic("error unmarshaling companies")
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
		panic("could not create the file in path " + outputPath)
	}

	_, err = fmt.Fprintf(file, "%v", companies)	
	if err != nil {
		fmt.Println(err)
		panic("Could not write companies data in path " + outputPath)
	}

	fmt.Printf("%v Companies have been generated succesfully at %v\n", companiesAmount, outputPath)
}

func readUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number of fake companies data you want to generate")

	companiesAmountStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic("error reading the input of companies to generate")
	}

	// Trim sufixes for Windows and Unix based systems
	companiesAmountStr = strings.TrimSuffix(companiesAmountStr, "\n")
	companiesAmountStr = strings.TrimSuffix(companiesAmountStr, "\r")

	companiesAmount, err := strconv.Atoi(companiesAmountStr)
	if err != nil {
		fmt.Println(err)
		panic("error reading the input of companies to generate")
	}

	return companiesAmount
}
