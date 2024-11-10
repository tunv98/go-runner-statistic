package main

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
)

// GoogleSheetsRepository defines a repository for accessing Google Sheets data
type GoogleSheetsRepository struct {
	service       *sheets.Service
	spreadsheetID string
}

// NewGoogleSheetsRepository creates a new instance of GoogleSheetsRepository
func NewGoogleSheetsRepository(credentialsFile, spreadsheetID string) (*GoogleSheetsRepository, error) {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	// Create a config from the service account JSON file
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return nil, err
	}

	// Create a new Sheets client
	client := config.Client(context.Background())
	service, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return &GoogleSheetsRepository{
		service:       service,
		spreadsheetID: spreadsheetID,
	}, nil
}

// FetchData retrieves data from the specified range in the Google Sheet
func (r *GoogleSheetsRepository) FetchData(readRange string) ([][]interface{}, error) {
	resp, err := r.service.Spreadsheets.Values.Get(r.spreadsheetID, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}
