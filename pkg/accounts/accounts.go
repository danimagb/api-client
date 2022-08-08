package accounts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danimagb/api-client/pkg/core"
	"github.com/danimagb/api-client/pkg/models"
	"github.com/google/uuid"
)

const(
	baseAccountsPath string = "/v1/organisation/accounts"
)

type AccountsClient struct{
	baseClient core.Client
}

func New(baseClient core.Client) *AccountsClient{
	return &AccountsClient{
		baseClient: baseClient,
	}
}

func(ac *AccountsClient) Fetch(ctx context.Context, id uuid.UUID) (*models.AccountResponse, error){
	accountResponse := &models.AccountResponse{}
	apiError := &models.APIError{}

	apiReq := core.NewRequestBuilder(http.MethodGet).
		WithPath(baseAccountsPath).
		WithPath(id.String()).
		WithContext(ctx).
		WithResultWriteTo(accountResponse).
		WithErrorWriteTo(apiError).
		Build()

	response, err := ac.baseClient.Send(apiReq)

	if err != nil{
		return  nil, err
	}

	if response.StatusCode() != 200{
		clientError := core.NewApiClientError(
			"Status code does not represent success for this request",
			response.StatusCode(),
			apiError.ErrorMessage, response)
		return nil, clientError
	}

	return accountResponse, nil
}

func(ac *AccountsClient) Create(ctx context.Context, accountData *models.AccountRequest) (*models.AccountResponse, error){
	accountResponse := &models.AccountResponse{}
	apiError := &models.APIError{}

	apiReq := core.NewRequestBuilder(http.MethodPost).
		WithPath(baseAccountsPath).
		WithBody(accountData).
		WithContext(ctx).
		WithResultWriteTo(accountResponse).
		WithErrorWriteTo(apiError).
		Build()

	response, err := ac.baseClient.Send(apiReq)

	if err != nil {
		return  nil, err
	}

	if response.StatusCode() != 201 {
		clientError := core.NewApiClientError(
			"Status code does not represent success for this request",
			response.StatusCode(),
			apiError.ErrorMessage, response)
		return nil, clientError
	}

	return accountResponse, nil
}

func(ac *AccountsClient) Delete(ctx context.Context, id uuid.UUID, version int64) error{
	apiError := &models.APIError{}

	apiReq := core.NewRequestBuilder(http.MethodDelete).
		WithPath(baseAccountsPath).
		WithPath(id.String()).
		WithQueryParam("version", strconv.FormatInt(version, 10)).
		WithContext(ctx).
		WithErrorWriteTo(apiError).
		Build()

	response, err := ac.baseClient.Send(apiReq)

	if err != nil {
		return  err
	}

	if response.StatusCode() != 204 {
		clientError := core.NewApiClientError(
			"Status code does not represent success for this request",
			response.StatusCode(),
			apiError.ErrorMessage, response)
		return clientError
	}

	return nil
}