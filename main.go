package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nwmandel/go-microservice/service"
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")

	flag.Parse()
	ctx := context.Background()

	serv := service.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := service.Endpoints{
		GetEndpoint:      service.MakeGetEndpoint(serv),
		StatusEndpoint:   service.MakeStatusEndpoint(serv),
		ValidateEndpoint: service.MakeValidateEndpoint(serv),
	}

	go func() {
		log.Println("service is listening on port:", *httpAddr)
		handler := service.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
