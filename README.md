# 🚀 go-federation-demo

This is a **Proof of Concept (POC)** project demonstrating **GraphQL with Apollo Federation** using **gqlgen**.

---

## 🧭 Project Structure

This project contains the following services:

- 🧑‍💼 `users` — User service (Go)
- 📦 `products` — Product service (Go)
- 💬 `reviews` — Review service (Go)
- 🌐 `gateway` — Apollo Gateway (Node.js)

---

## 🏃‍♂️ How to Run the Project

Open **4 different terminals**, one for each service.

---

### 🖥️ Terminal 1 – Users Service
```bash
cd users
go mod tidy
go run server.go
```

### 🖥️ Terminal 2 – Products Service
```bash

cd products
go mod tidy
go run server.go
```

### 🖥️ Terminal 3 – Reviews Service
```bash

cd reviews
go mod tidy
go run server.go
```

### 🖥️ Terminal 4 – Apollo Gateway
```bash
npm install --save @apollo/gateway apollo-server graphql
cd gateway
node index.js
```
