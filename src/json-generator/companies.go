package jsongenerator

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

// CreateFakeCompanies creates a lot of fake companies using JSON format
func CreateFakeCompanies(numOfCompanies int) (string, error) {
	if numOfCompanies < 1 {
		return "", fmt.Errorf("the number you provide should be a positive integer")
	}

	type Company struct {
		Id           string `json:"id,omitempty"`
		CompanyName  string `json:"company_name,omitempty"`
		NumEmployees int    `json:"num_employees,omitempty"`
		Country      string `json:"country,omitempty"`
	}

	companies := []Company{}

	for i := 0; i < numOfCompanies; i++ {
		company := Company{
			Id:           gofakeit.UUID(),
			CompanyName:  gofakeit.Company(),
			Country:      gofakeit.CountryAbr(),
			NumEmployees: gofakeit.Number(0, 100000),
		}

		companies = append(companies, company)

	}

	companiesJSON, err := json.MarshalIndent(companies, "", "    ")
	if err != nil {
		return "", fmt.Errorf("error while marshaling companies")
	}

	return string(companiesJSON), nil
}
