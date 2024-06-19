package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	jsongenerator "github.com/jigth/companies-json-generator/src/json-generator"
)

func main() {
	outputPath, companiesAmount, err := parseArguments()
	if err != nil {
		panic(err)
	}

	companies, err := jsongenerator.CreateFakeCompanies(companiesAmount)
	if err != nil {
		panic(err)
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

	file.Close()

	fmt.Printf("%v Companies have been generated succesfully at %v\n", companiesAmount, outputPath)
}

func parseArguments() (string, int, error) {
	if len(os.Args) < 3 {
		return "", 0, errors.New(
			"\n\nerror parsing arguments.\nUse the program like this:\n" +
				"go run main.go <outputPath> <numOfOutputElements>\n" +
				"For example, to generate 30 inputs do this command:\n\t" +
				"go run main.go myFile.json 30",
		)
	}

	outputPath := os.Args[1]
	objectsAmount := os.Args[2] // Num of objects to generate

	parsedAmount, err := getParsedInt(objectsAmount)
	if err != nil {
		return "", 0, err
	}

	return outputPath, parsedAmount, nil
}

func getParsedInt(strNumber string) (int, error) {
	companiesAmount, err := strconv.Atoi(strNumber)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("could not parse objects amount, be sure you use an int")
	}

	return companiesAmount, nil
}
