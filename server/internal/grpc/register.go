package grpc

import (
	"github.com/Neumann88/currency-grpc/api"
	"github.com/Neumann88/currency-grpc/server/internal/grpc/handler"
	"google.golang.org/grpc"
)

type Debs struct {
	*handler.CurrencyHander
}

func Register(srv grpc.ServiceRegistrar, d Debs) {
	api.RegisterCurrencyServiceServer(srv, d.CurrencyHander)
}
