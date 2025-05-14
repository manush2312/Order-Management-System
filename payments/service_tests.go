package main

import (
	"context"
	"testing"

	"github.com/manush2312/commons/api"
	"github.com/manush2312/oms-payments/processor/inmem"
)

func TestService(t *testing.T) {
	// we are not directly testing the service on stripe, we are creating a inmemory processor and testing the service on it.
	processor := inmem.NewInmem()
	svc := NewService(processor)

	t.Run("shold create a payement link", func(t *testing.T) {
		link, err := svc.CreatePayment(context.Background(), &api.Order{})

		if err != nil {
			t.Errorf("CreatePayment() error = %v, want nil", err)
		}
		if link == "" {
			t.Error("CreatePayment() link is empty")
		}

	})
}
