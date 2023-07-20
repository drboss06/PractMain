package http

import (
	"authPract/pkg/api"
	"authPract/pkg/service"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type SomeStruct struct {
}

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

		if strings.HasPrefix(r.URL.Path, "/upload") {
			file, handler, err := r.FormFile("file")
			if err != nil {
				log.Fatalf("Failed to open file: %v", err)
			}
			defer file.Close()
			f, err := os.OpenFile(viper.GetString("filedir")+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("Failed to open file: %v", err)
			}
			defer f.Close()

			io.Copy(f, file)
			////here we save our file to our path
			CSVServ := service.NewCsvProp()
			table_name, err := CSVServ.ParseCSV(handler.Filename)
			if err != nil {
				log.Fatalf("Failed to open file: %v", err)
			}

			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["table name"] = table_name
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)

			//data := SomeStruct{}
			//w.Header().Set("Content-Type", "application/json")
			//w.WriteHeader(http.StatusCreated)
			//json.NewEncoder(w).Encode(data)

			return
		}

		gw.ServeHTTP(w, r)
	})
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
