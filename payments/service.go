package main

import (
	"context"

	pb "github.com/manush2312/commons/api"
	"github.com/manush2312/oms-payments/processor"
)

type service struct {
	// we will recieve payment processor as a dependency.
	processor processor.PaymentProcessor
}

func NewService(processor processor.PaymentProcessor) *service {
	return &service{processor: processor}
}

func (s *service) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	// we need to connect to payment processor

	link, err := s.processor.CreatePaymentLink(o)
	if err != nil {
		return "", err
	}

	// update the order with the link.

	return link, nil // here we are returning payment link for order created.
}
