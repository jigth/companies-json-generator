package main

import (
	"fmt"
	jsongenerator "github.com/jigth/companies-json-generator/src/json-generator"
)

func main() {
	ONE_MILLION := 1000*1000
	companies, err := jsongenerator.CreateFakeCompanies(ONE_MILLION)
	if err != nil {
		panic("error unmarshaling companies")
	}

	fmt.Println(companies)
}
