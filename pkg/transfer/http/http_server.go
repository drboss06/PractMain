package http

import (
	"authPract/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strings"
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
	err = api.RegisterProjectsHandlerFromEndpoint(ctx, gw, "localhost:8083", opts)

	mux := http.NewServeMux()
	//mux.Handle("/", gw)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			tokenHeader := r.Header.Get("Authorization")
			token := strings.Split(tokenHeader, " ")[1]

			//userId, err := authPract.ParseToken(token)
			if err != nil {
				http.Error(w, "Token is bead", 400)
				return
			}

			if token != "" {
				//r.WithContext(context.WithValue(r.Context(), "userId", cast.ToString(userId)))
				gw.ServeHTTP(w, r)
				return
			} else if token == "" {
				http.Error(w, "You are not authorise", 400)
				return
			} else {
				http.Error(w, "You are not authorise", 400)
				return
			}
		}

		gw.ServeHTTP(w, r)
	})
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
