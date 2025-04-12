package main

import (
	"context"

	pb "github.com/manush2312/commons/api"
)

type PaymentsService interface {
	CreatePayment(context.Context, *pb.Order) (string, error) // the string that we are returning will be payment link for order.
}
