package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
}

func NewServer() *Server {
	return &Server{
		srv: grpc.NewServer(),
	}
}

func (s *Server) GrpcServer() *grpc.Server {
	return s.srv
}

func (s *Server) Listen(port int) (net.Listener, error) {
	addres := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", addres)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func (s *Server) Serve(l net.Listener) error {
	if err := s.srv.Serve(l); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
