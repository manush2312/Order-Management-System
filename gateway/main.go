package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/manush2312/commons"
	"github.com/manush2312/commons/discovery"
	"github.com/manush2312/commons/discovery/consul"
	"github.com/manush2312/oms-gateway/gateway"
)

var (
	serviceName = "gateway"
	httpAddr    = common.EnvString("HTTP_ADDR", ":8081")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main() {
	// conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to dial server: %v", err)
	// }
	// defer conn.Close()

	// log.Println("Dialing order service at ", orderServiceAddr)

	// c := pb.NewOrderServiceClient(conn)

	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
		panic(err)
	}

	// for health check status
	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	// we are going to use go 122 http vanilla server.
	mux := http.NewServeMux() // built-in router in Go's library.

	ordersGateway := gateway.NewGRPCGateway(registry)

	handler := NewHandler(ordersGateway)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	// we need to start http server and pass in Handler as our parameter
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start http server ")
	}
}
