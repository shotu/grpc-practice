package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shotu/grpc-practice/graphqlapi"
	"github.com/shotu/grpc-practice/grpc-practice/hellopb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GetUser(ctx context.Context, request *hellopb.GraphqlapiRequest) (*hellopb.GraphqlapiResponse, error) {
	userId := request.Id
	fmt.Println("UserID ", userId)
	var grphqlEP string
	grphqlEP = "http://localhost:8080/graphql?query={user(id:%221%22){id,name}}" // TODO change to this use the userid from request

	gqlResponse := graphqlapi.QueryGraphQLEP(grphqlEP)

	response := &hellopb.GraphqlapiResponse{
		Id:   gqlResponse.Data.User.ID,
		Name: gqlResponse.Data.User.Name,
	}
	return response, nil
}

func main() {

	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)
	s := grpc.NewServer()
	hellopb.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)
}
