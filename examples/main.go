package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	client "github.com/danimagb/api-client/pkg"
	"github.com/danimagb/api-client/pkg/models"
	"github.com/google/uuid"
)

func main(){
	ctx := context.Background()
	u, err := url.Parse("http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}

	client, err := client.NewClient(
		client.WithBaseUrl(*u),
		client.WithTimeoutInMilliseconds(30),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountClassification := "Personal"
	country := "GB"
	name := "Daniel"
	jointAccount := false
	accountMatchingOutput := false
	status := "confirmed"
	newAccount := &models.AccountRequest{
		Data: &models.AccountData{
			Attributes: &models.AccountAttributes{
				AccountClassification: &accountClassification,
				AccountMatchingOptOut: &accountMatchingOutput,
				AccountNumber: "41426815",
				AlternativeNames: []string{name},
				BankID: "400300",
				BankIDCode: "GBDSC",
				BaseCurrency: "GBP",
				Bic: "NWBKGB22",
				Country: &country,
				Iban: "GB11NWBK40030041426819",
				JointAccount: &jointAccount,
				Name: []string{name},
				SecondaryIdentification: "A1B2C3D4",
				Status: &status,
				Switched: &accountMatchingOutput,
			},
			ID: uuid.NewString(),
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type: "accounts",
		},
	}


	accountCreationResponse, err := client.Accounts.Create(ctx, newAccount)

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	s, _:=json.MarshalIndent(accountCreationResponse, "", "\t")
	fmt.Printf("Created with success: \n%s\n", s)

	accountFetchResponse, err := client.Accounts.Fetch(ctx, uuid.MustParse(accountCreationResponse.Data.ID))

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	s, _=json.MarshalIndent(accountFetchResponse, "", "\t")
	fmt.Printf("Fetched with success: \n%s\n", s)

	err = client.Accounts.Delete(ctx, uuid.MustParse(accountFetchResponse.Data.ID), *accountFetchResponse.Data.Version)

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	fmt.Println("Deleted with success")
}