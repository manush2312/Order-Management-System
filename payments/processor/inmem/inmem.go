package inmem

import pb "github.com/manush2312/commons/api"

type Inmem struct{}

func NewInmem() *Inmem {
	return &Inmem{}
}

func (i *Inmem) CreatePaymentLink(order *pb.Order) (string, error) {
	// In a real implementation, this would create a payment link using an in-memory store
	// For this example, we'll just return a dummy link
	return "dummy/link", nil
}
