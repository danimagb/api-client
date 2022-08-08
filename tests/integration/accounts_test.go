package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/danimagb/api-client/pkg/core"
	"github.com/danimagb/api-client/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	sut := SetupNewClient(t)

	t.Run("Given an existent account id should return the account", func(t *testing.T) {
		// Arrange
		expected, err := sut.Accounts.Create(context.Background(), buildTestAccount())

		if err != nil {
			t.Errorf("Error creating test account: %v", err)
		}

		// Act
		actual, err := sut.Accounts.Fetch(context.Background(), uuid.MustParse(expected.Data.ID))

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Given a non existent account id should return an error with 404 status code", func(t *testing.T) {
		// Act
		actual, err := sut.Accounts.Fetch(context.Background(), uuid.New())

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		assert.IsType(t, http.StatusNotFound, err.(*core.ApiClientError).StatusCode)
	})
}

func TestCreate(t *testing.T) {
	sut := SetupNewClient(t)

	t.Run("Given a valid account should return the account created", func(t *testing.T) {
		// Act
		actual, err := sut.Accounts.Create(context.Background(), buildTestAccount())

		// Assert
		assert.Nil(t, err)
		assert.NotEmpty(t, actual)
	})

	t.Run("Given an invalid empty account should return an error with 400 status code", func(t *testing.T) {
		// Act
		actual, err := sut.Accounts.Create(context.Background(), &models.AccountRequest{})

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		assert.IsType(t, http.StatusBadRequest, err.(*core.ApiClientError).StatusCode)
	})
}

func TestDelete(t *testing.T) {
	sut := SetupNewClient(t)

	t.Run("Given an existent account id and version should not return any error", func(t *testing.T) {
		// Arrange
		expected, err := sut.Accounts.Create(context.Background(), buildTestAccount())

		if err != nil {
			t.Errorf("Error creating test account: %v", err)
		}

		// Act
		err = sut.Accounts.Delete(context.Background(), uuid.MustParse(expected.Data.ID), *expected.Data.Version)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("Given a non existent account id should return an error with 404 status code", func(t *testing.T) {
		// Act
		err := sut.Accounts.Delete(context.Background(), uuid.New(), 0)

		// Assert
		assert.NotNil(t, err)
		assert.IsType(t, http.StatusNotFound, err.(*core.ApiClientError).StatusCode)
	})
}


func buildTestAccount() *models.AccountRequest{

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

	return newAccount
}