# API Showcase

REST, GraphQL, and gRPC implemented in Python, Go, and Java.  
All 9 services expose the same **Product CRUD** domain.

## Structure

```
api-showcase/
├── docs/           Run guides — local, Docker, Minikube
├── proto/          Canonical product.proto (source of truth)
├── k8s/
│   ├── rest/       Kubernetes manifests for REST services
│   ├── graphql/    Kubernetes manifests for GraphQL services
│   └── grpc/       Kubernetes manifests for gRPC services
├── python/
│   ├── rest/       FastAPI · port 8000
│   ├── graphql/    Strawberry + FastAPI · port 8001
│   └── grpc/       grpcio · port 50051
├── go/
│   ├── rest/       net/http · port 8080
│   ├── graphql/    graphql-go · port 8081
│   └── grpc/       google.golang.org/grpc · port 50052
├── java/
│   ├── rest/       Spring Boot · port 9000
│   ├── graphql/    Spring for GraphQL · port 9001
│   └── grpc/       grpc-spring-boot-starter · port 9090
├── docker-compose.yml
└── README.md
```

## How to Run

| Guide | Link |
|-------|------|
| Run locally (no Docker) | [docs/local.md](docs/local.md) |
| Run with Docker | [docs/docker.md](docs/docker.md) |
| Run on Minikube | [docs/minikube.md](docs/minikube.md) |

## Run Everything with Docker Compose

```bash
docker-compose up --build
```

## Port Reference

| Service        | Local  | Minikube NodePort |
|----------------|--------|-------------------|
| python-rest    | 8000   | 30000             |
| python-graphql | 8001   | 30001             |
| python-grpc    | 50051  | 30002             |
| go-rest        | 8080   | 30003             |
| go-graphql     | 8081   | 30004             |
| go-grpc        | 50052  | 30005             |
| java-rest      | 9000   | 30006             |
| java-graphql   | 9001   | 30007             |
| java-grpc      | 9090   | 30008             |
