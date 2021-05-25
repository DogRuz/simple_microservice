package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"simple_microservice/pkg/api"
)

func main() {
	flag.Parse()
	inn := flag.Arg(0)
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewGetInformationClient(conn)

	res, err = c.Add(context.Background(), &api.GetInformationRequest{Inn: inn})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
