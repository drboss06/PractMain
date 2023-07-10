package http

import (
	"authPract/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//opts2 := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := api.RegisterAdderHandlerFromEndpoint(ctx, gw, "localhost:8080", opts2)
	err := api.RegisterAdderHandlerFromEndpoint(ctx, gw, "localhost:8080", opts)
	if err != nil {
		panic(err)
	}
	err = api.RegisterAuthHandlerFromEndpoint(ctx, gw, "localhost:8082", opts)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
