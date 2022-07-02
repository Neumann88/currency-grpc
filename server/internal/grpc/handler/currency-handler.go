package handler

import (
	"context"

	"github.com/Neumann88/currency-grpc/api"
)

type CurrencyHander struct {
	api.UnimplementedCurrencyServiceServer
}

func NewCurrencyHandler() *CurrencyHander {
	return &CurrencyHander{}
}

func (a *CurrencyHander) GetCurrency(ctx context.Context, req *api.CurrencyRequest) (*api.CurrencyResponse, error) {
	return &api.CurrencyResponse{
		Name:        api.Currency_USD,
		Country:     "USA",
		Description: "description",
		Change:      1,
		RateUSD:     1,
		UpdatedAt:   "date",
	}, nil
}
