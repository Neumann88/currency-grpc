package main

import (
	"log"

	desc "github.com/Neumann88/currency-grpc/server/internal/grpc"
	"github.com/Neumann88/currency-grpc/server/internal/grpc/handler"
	"github.com/Neumann88/currency-grpc/server/pkg/grpc"
)

func main() {
	srv := grpc.NewServer()

	lstr, err := srv.Listen(8080)
	if err != nil {
		log.Fatalln(err)
	}

	desc.Register(
		srv.GrpcServer(),
		desc.Debs{
			CurrencyHander: handler.NewCurrencyHandler(),
		})

	if err := srv.Serve(lstr); err != nil {
		log.Fatalln(err)
	}
}
