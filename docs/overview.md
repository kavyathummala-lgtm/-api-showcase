# Project Overview — What You Built and Why

## The Big Idea

You built a **showcase project** that implements the exact same functionality — a Product catalog with Create, Read, Update, and Delete (CRUD) operations — in **9 different ways**.

The goal is to demonstrate how the same business logic looks and behaves across:
- **3 programming languages** (Python, Go, Java)
- **3 API communication styles** (REST, GraphQL, gRPC)

This gives you (and anyone reading your code) a direct side-by-side comparison of all three API styles in all three languages. It is a hands-on learning reference and a portfolio project.

---

## The Data Model

Every single service in this project works with one shared concept: a **Product**.

```
Product {
  id       — Unique identifier (UUID, auto-generated)
  name     — Product name (e.g. "Laptop")
  category — Product category (e.g. "Electronics")
  price    — Price as a decimal number (e.g. 999.99)
  stock    — Quantity available (e.g. 50)
}
```

This is defined once in `/proto/product.proto` as the canonical source of truth, then reimplemented independently in each language.

---

## The Three API Styles Explained

### 1. REST (Representational State Transfer)

REST is the most common style for web APIs. It maps operations to HTTP verbs and URLs.

| Operation | HTTP Method | URL |
|-----------|------------|-----|
| Get all products | GET | `/products` |
| Get one product | GET | `/products/{id}` |
| Create a product | POST | `/products` |
| Update a product | PUT | `/products/{id}` |
| Delete a product | DELETE | `/products/{id}` |

**How it works:** The client makes an HTTP request to a specific URL. The server returns JSON.

**Why use REST?** Simple, widely understood, works with any HTTP client (browsers, curl, Postman), and is stateless.

**Analogy:** Like a menu at a restaurant — each item has a fixed name (URL) and action (verb).

---

### 2. GraphQL

GraphQL is a query language for APIs developed by Facebook. Instead of many endpoints, it has **one endpoint** (`/graphql`) and the client specifies exactly what data it wants.

| Operation | Type | Example |
|-----------|------|---------|
| Get all products | Query | `{ products { id name price } }` |
| Get one product | Query | `{ product(id: "abc") { name } }` |
| Create a product | Mutation | `mutation { createProduct(input: {...}) { id } }` |
| Update a product | Mutation | `mutation { updateProduct(id: "abc", input: {...}) { id } }` |
| Delete a product | Mutation | `mutation { deleteProduct(id: "abc") }` |

**How it works:** The client sends a JSON body to `/graphql` containing a query string that describes exactly which fields it wants. The server returns only those fields.

**Why use GraphQL?** The client controls the shape of the response. No over-fetching (getting too much data) or under-fetching (needing multiple requests). Great for complex UIs.

**Analogy:** Like ordering a custom sandwich — you tell the kitchen exactly what you want, instead of choosing from a fixed menu.

---

### 3. gRPC (Google Remote Procedure Call)

gRPC is a high-performance framework developed by Google. Instead of HTTP verbs and URLs, the client calls **functions** (called RPCs) directly on the server.

| Operation | RPC Method |
|-----------|-----------|
| Get all products | `GetProducts()` |
| Get one product | `GetProduct(id)` |
| Create a product | `CreateProduct(input)` |
| Update a product | `UpdateProduct(id, input)` |
| Delete a product | `DeleteProduct(id)` |

**How it works:** The API is defined in a `.proto` file using Protocol Buffers. Code is generated in each language from this definition. Data is sent as compact binary (not text JSON).

**Why use gRPC?** Much faster than REST (binary protocol, HTTP/2). Strongly typed (the contract is enforced at compile time). Great for microservices communicating internally.

**Analogy:** Like calling a function in code — you call `GetProduct(id)` and get back a `Product` object, rather than thinking in terms of HTTP.

---

## The Three Languages Compared

### Python (with FastAPI and grpcio)

- **REST:** Uses **FastAPI** — a modern, fast Python web framework. You define routes with decorators (`@app.get("/products")`). Very little boilerplate.
- **GraphQL:** Uses **Strawberry** — a Python-first GraphQL library that uses Python type hints and dataclasses to define the schema.
- **gRPC:** Uses **grpcio** — Google's official gRPC library for Python. Stubs are auto-generated from the `.proto` file.

**Best for:** Rapid development, data science, scripting. Easiest to read and write.

---

### Go (with net/http and graphql-go)

- **REST:** Uses Go's built-in **net/http** standard library — no external web framework needed. Routes are registered manually.
- **GraphQL:** Uses **graphql-go** — a Go port of the reference GraphQL implementation. Schema is defined programmatically in code.
- **gRPC:** Uses **google.golang.org/grpc** — the official Go gRPC library. Stubs are auto-generated from the `.proto` file.

**Best for:** High performance, low memory usage, concurrent systems. Statically compiled to a single binary.

---

### Java (with Spring Boot)

- **REST:** Uses **Spring Boot** with `@RestController` and `@RequestMapping`. The most verbose of the three but very feature-rich and production-proven.
- **GraphQL:** Uses **Spring for GraphQL** — the official Spring integration. Schema is defined in a `.graphqls` file, and Java methods are annotated with `@QueryMapping` / `@MutationMapping`.
- **gRPC:** Uses **grpc-spring-boot-starter** — integrates gRPC into Spring Boot. The service implementation extends an auto-generated base class from the `.proto`.

**Best for:** Enterprise applications, large teams, strong ecosystem, long-term maintainability.

---

## How the Project Is Structured

```
api-showcase/
│
├── proto/product.proto          ← THE CANONICAL DATA MODEL (read this first)
│
├── python/
│   ├── rest/main.py             ← FastAPI app with 5 route handlers
│   ├── graphql/main.py          ← Strawberry schema + FastAPI app
│   └── grpc/
│       ├── server.py            ← gRPC server implementation
│       └── client.py            ← Test client (runs all CRUD operations)
│
├── go/
│   ├── rest/main.go             ← net/http server with route handlers
│   ├── graphql/main.go          ← graphql-go schema + HTTP handler
│   └── grpc/
│       ├── server/main.go       ← gRPC server implementation
│       └── client/main.go       ← Test client (runs all CRUD operations)
│
├── java/
│   ├── rest/                    ← Spring Boot REST application
│   ├── graphql/                 ← Spring for GraphQL application
│   └── grpc/                    ← Spring Boot gRPC application
│
├── docker-compose.yml           ← Runs all 9 services in containers
├── k8s/                         ← Kubernetes manifests (Minikube deployment)
└── docs/                        ← This documentation
```

---

## Storage: How Data Is Stored

**All services use in-memory storage** — a simple dictionary or map that lives in RAM while the server is running.

```
Python: products = {}          (a Python dict)
Go:     products map[string]Product  (with sync.RWMutex for thread safety)
Java:   ConcurrentHashMap<String, Product>  (thread-safe map)
```

This means:
- **No database is needed** — great for a demo/learning project
- **Data is lost when the server restarts** — this is intentional; the focus is on API design, not persistence
- **IDs are UUID v4** — randomly generated, universally unique

---

## The Three Ways to Run Everything

### Option 1: Run Locally (Direct)

Each service runs as a plain process on your machine. Best for development and debugging individual services.

```bash
# Example: run Python REST
cd python/rest
python -m venv venv && source venv/bin/activate
pip install -r requirements.txt
uvicorn main:app --reload --port 8000
```

See `docs/local.md` for all commands.

### Option 2: Docker Compose

All 9 services start with one command. Each runs in its own container. Best for running the full stack at once.

```bash
docker-compose up --build
```

See `docs/docker.md` for details.

### Option 3: Minikube (Kubernetes)

All 9 services are deployed as pods in a Kubernetes cluster running on your laptop. Best for learning Kubernetes and production-like deployment.

```bash
minikube start --driver=docker
# build images and apply k8s manifests
kubectl apply -f k8s/rest/ -f k8s/graphql/ -f k8s/grpc/
```

See `docs/minikube.md` for the full walkthrough.

---

## Key Design Decisions

### Why the same domain (Products) everywhere?

Keeping the business logic identical across all 9 services lets you focus on comparing the **API style and language differences** without confusion about the actual data or operations.

### Why no database?

Adding a database would require connection strings, schemas, migrations, and setup steps that distract from the API comparison. In-memory storage keeps each service self-contained and easy to start.

### Why a `.proto` file as the source of truth?

Protocol Buffers (`.proto`) provide a language-neutral way to define a data schema and service contract. Even the REST and GraphQL services follow the same model — using the `.proto` as the canonical definition ensures everything stays consistent.

### Why no authentication or validation?

This is a showcase and learning project. Auth and input validation add complexity that obscures the API patterns. The focus is on the structure and communication style of each API.

---

## What You Can Learn from This Project

| You want to understand... | Look at... |
|--------------------------|-----------|
| REST endpoint structure | Any `rest/main.py`, `rest/main.go`, or `ProductController.java` |
| GraphQL schema definition | `python/graphql/main.py` (Strawberry), `java/graphql/src/.../schema.graphqls` |
| gRPC service definition | `proto/product.proto` |
| gRPC server implementation | Any `grpc/server.py` or `grpc/server/main.go` |
| How to call a gRPC server | `python/grpc/client.py` or `go/grpc/client/main.go` |
| Docker multi-stage builds | Any `Dockerfile` in `go/` or `java/` |
| Kubernetes deployment | Any YAML in `k8s/` |
| Spring Boot REST controller | `java/rest/src/.../ProductController.java` |
| FastAPI route definition | `python/rest/main.py` |

---

## Summary Table

| | REST | GraphQL | gRPC |
|---|------|---------|------|
| **Protocol** | HTTP 1.1 | HTTP 1.1 | HTTP/2 |
| **Data format** | JSON (text) | JSON (text) | Protocol Buffers (binary) |
| **Endpoints** | Many (`/products`, `/products/{id}`) | One (`/graphql`) | Method calls (`GetProduct`, etc.) |
| **Schema definition** | Implicit (from code) | Explicit (GraphQL SDL) | Explicit (`.proto` file) |
| **Tooling to test** | curl, Postman, browser | curl, GraphQL Playground | grpcurl, generated client |
| **Best for** | Public APIs, browsers, simplicity | Flexible UIs, mobile apps | Internal microservices, performance |
| **Python impl** | FastAPI | Strawberry + FastAPI | grpcio |
| **Go impl** | net/http | graphql-go | google.golang.org/grpc |
| **Java impl** | Spring Boot | Spring for GraphQL | grpc-spring-boot-starter |
