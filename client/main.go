package main

import (
	"context"
	"log"

	"github.com/Neumann88/currency-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), ":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := api.NewCurrencyServiceClient(conn)

	resp, err := c.GetCurrency(context.Background(), &api.CurrencyRequest{Name: api.Currency_USD})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(
		resp.Name,
		resp.Country,
		resp.Description,
		resp.Change,
		resp.RateUSD,
		resp.UpdatedAt,
	)
}
