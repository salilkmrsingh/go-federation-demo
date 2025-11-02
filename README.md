# go-federation-demo
This is a POC project for graphql with apollo federation in gqlgen

To run this project.

Open 4 different terminals

Terminal 1

cd users
go mod tidy
go run users/server.go

Terminal 2

cd products
go mod tidy
go run products/server.go

Terminal 3

cd reviews
go mod tidy
go run reviews/server.go

Terminal 4

npm install --save @apollo/gateway apollo-server graphql
cd gateway
node index.js

