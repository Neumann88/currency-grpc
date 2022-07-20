package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Neumann88/currency-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type data struct {
	Date string          `json:"date"`
	Data []*api.Currency `json:"data"`
}

func main() {
	conn, err := grpc.DialContext(context.Background(), ":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := api.NewCurrencyServiceClient(conn)

	resp, err := c.GetCurrency(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalln(err)
	}

	d := data{
		Date: resp.Date,
		Data: resp.Data,
	}

	b, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("test.json", b, 0644)

}
