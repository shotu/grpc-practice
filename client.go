package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shotu/grpc-practice/grpc-practice/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := hellopb.NewUserServiceClient(cc)
	request := &hellopb.GraphqlapiRequest{Id: "2"}
	resp, err := client.GetUser(context.Background(), request)
	if err != nil {
		fmt.Println("Error accessing the graphql server:", err)
	}
	fmt.Printf("Receive response ", resp)
}
