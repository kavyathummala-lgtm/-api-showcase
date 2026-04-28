# Running with Docker

## Prerequisites
- Docker Desktop running

---

## Run a Single Service

Pick any service below, build its image, and run it.

### Python
```bash
# REST
docker build -t python-rest ./python/rest
docker run -p 8000:8000 python-rest

# GraphQL
docker build -t python-graphql ./python/graphql
docker run -p 8001:8001 python-graphql

# gRPC
docker build -t python-grpc ./python/grpc
docker run -p 50051:50051 python-grpc
```

### Go
```bash
# REST
docker build -t go-rest ./go/rest
docker run -p 8080:8080 go-rest

# GraphQL
docker build -t go-graphql ./go/graphql
docker run -p 8081:8081 go-graphql

# gRPC
docker build -t go-grpc ./go/grpc
docker run -p 50052:50052 go-grpc
```

### Java
```bash
# REST  (first build takes ~2 min — Maven downloads deps)
docker build -t java-rest ./java/rest
docker run -p 9000:9000 java-rest

# GraphQL
docker build -t java-graphql ./java/graphql
docker run -p 9001:9001 java-graphql

# gRPC
docker build -t java-grpc ./java/grpc
docker run -p 9090:9090 java-grpc
```

---

## Run All 9 Services at Once

```bash
# From the api-showcase root
docker-compose up --build
```

To run only a subset:
```bash
docker-compose up --build python-rest go-rest java-rest
```

To stop everything:
```bash
docker-compose down
```

---

## Port Map

| Service        | Host Port |
|----------------|-----------|
| python-rest    | 8000      |
| python-graphql | 8001      |
| python-grpc    | 50051     |
| go-rest        | 8080      |
| go-graphql     | 8081      |
| go-grpc        | 50052     |
| java-rest      | 9000      |
| java-graphql   | 9001      |
| java-grpc      | 9090      |

---

## Verify a Running Container

```bash
# List running containers
docker ps

# Tail logs
docker logs -f <container-name>

# Stop a single container
docker stop <container-name>
```
