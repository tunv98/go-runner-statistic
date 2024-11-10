package main

import (
	"fmt"
	"log"
)

const (
	spreadsheetId   = "1ZezPDSiHrKY6U1C1ARBPcS1TRHcgU4MoblRO1K0wxCU"
	credentialsFile = "./service_account.json"
	nameRange       = "15/7!A2:A16"
	amountRange     = "15/7!F2:F16"
)

func main() {
	repo, err := NewGoogleSheetsRepository(credentialsFile, spreadsheetId)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	names, err := repo.FetchData(nameRange)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	amounts, err := repo.FetchData(amountRange)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	report := report{
		fromDate:  "15/07/2024",
		toDate:    "02/12/2024",
		summaries: makeSummaries(names, amounts),
	}
	fmt.Println(report.toString())
}
