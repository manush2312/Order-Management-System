package main

import (
	"context"

	pb "github.com/manush2312/commons/api"
)

type service struct {
	// we will recieve payment processor as a dependency.
}

func NewService() *service {
	return &service{}
}

func (s *service) CreatePayment(context.Context, *pb.Order) (string, error) {
	// we need to connect to payment processor
	return "", nil // here we are returning payment link for order created.
}
