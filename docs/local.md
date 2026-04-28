# Running Locally

All services expose a **Product CRUD** API. No Docker needed — just the language runtime.

## Prerequisites

| Language | Required |
|----------|----------|
| Python   | Python 3.11+, pip |
| Go       | Go 1.22+ |
| Java     | Java 21+, Maven 3.9+ |
| gRPC (any) | `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc` |

---

## Python

### REST — port 8000
```bash
cd python/rest
python -m venv venv && source venv/bin/activate   # Windows: venv\Scripts\activate
pip install -r requirements.txt
uvicorn main:app --reload --port 8000
```
Swagger UI → http://localhost:8000/docs

### GraphQL — port 8001
```bash
cd python/graphql
python -m venv venv && source venv/bin/activate
pip install -r requirements.txt
uvicorn main:app --reload --port 8001
```
Playground → http://localhost:8001/graphql

### gRPC — port 50051
```bash
cd python/grpc
python -m venv venv && source venv/bin/activate
pip install -r requirements.txt

# Generate stubs once (creates product_pb2.py and product_pb2_grpc.py)
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. product.proto

# Terminal 1 — server
python server.py

# Terminal 2 — test with client
python client.py
```

---

## Go

### REST — port 8080
```bash
cd go/rest
go mod tidy
go run main.go
```
API → http://localhost:8080/products

### GraphQL — port 8081
```bash
cd go/graphql
go mod tidy
go run main.go
```
Endpoint → http://localhost:8081/graphql (POST)

### gRPC — port 50052
```bash
cd go/grpc

# Generate stubs once
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=. proto/product.proto
go mod tidy

# Terminal 1 — server
go run ./server

# Terminal 2 — test with client
go run ./client
```

---

## Java

### REST — port 9000
```bash
cd java/rest
mvn spring-boot:run
```
API → http://localhost:9000/products

### GraphQL — port 9001
```bash
cd java/graphql
mvn spring-boot:run
```
GraphiQL UI → http://localhost:9001/graphiql

### gRPC — port 9090
```bash
cd java/grpc
mvn package -DskipTests
java -jar target/*.jar
```
Test with grpcurl:
```bash
grpcurl -plaintext localhost:9090 list
grpcurl -plaintext -d '{}' localhost:9090 product.ProductService/GetProducts
```

---

## Quick Test with curl

### REST
```bash
# Create
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","category":"Electronics","price":999.99,"stock":50}'

# Get all
curl http://localhost:8000/products
```

### GraphQL
```bash
curl -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ products { id name price } }"}'

curl -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"mutation { createProduct(input:{name:\"Laptop\",category:\"Electronics\",price:999.99,stock:50}){ id name } }"}'
```
