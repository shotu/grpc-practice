# grpc-practice

This codebase has three components. Please follow the steps to run the components

## 1. Start the simple graphql server
 ```go run graphqlserver/server.go```

 This will start the server http://localhost:8080

 You can access the endpoint exposed using curl : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'"

## 2. Start the gRPC server

   ```go run main.go```

   server will start listening at 0.0.0.0:50051
## 3. Access the gRPC server using golang client

   ```go run client.go```

   This will get the response from gRPC server which internally queries the Graphql server and servers the response over gRPC
