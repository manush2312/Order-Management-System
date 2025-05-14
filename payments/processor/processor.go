package processor

import (
	pb "github.com/manush2312/commons/api"
)

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
