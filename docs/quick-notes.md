# My Quick Notes — Simple and Short

---

## What Did I Build?
I built **9 API servers** that all manage a product list.
Each server supports: **Create, Read, Update, Delete** (CRUD) operations.

**Why:** To learn and compare 3 API styles across 3 languages side by side.

---

## 3 API Styles

**REST**
- **What:** Each operation has its own HTTP endpoint
- **Why:** Most widely used API style — every web and mobile app uses it
- **How:**
  - `GET /products` = fetch all products
  - `POST /products` = create a product
  - `PUT /products/id` = update a product
  - `DELETE /products/id` = delete a product

**GraphQL**
- **What:** Single endpoint `/graphql` — client asks for exactly what it needs
- **Why:** Avoids over-fetching — REST sends all fields, GraphQL sends only requested fields
- **How:** Client sends a query like `{ products { name price } }` and gets back only name and price

**gRPC**
- **What:** Client calls a function directly on a remote server
- **Why:** Faster than REST — uses binary Protocol Buffers instead of JSON text
- **How:** Defined by a `.proto` file — code is auto-generated from it in every language

---

## 3 Languages

| Language | Framework | Why Used |
|----------|-----------|---------|
| Python | FastAPI, Strawberry, grpcio | Simple syntax, fast to build |
| Go | net/http, graphql-go, grpc | High performance, low memory |
| Java | Spring Boot, Spring GraphQL, grpc-starter | Standard in large enterprises |

---

## Key Tools

**Docker**
- **What:** Packages the app and all its dependencies into a container image
- **Why:** Solves "works on my machine" problem — runs the same everywhere
- **How:** Write a `Dockerfile` → run `docker build` → run `docker run`

**Kubernetes**
- **What:** Orchestrates and manages containers automatically
- **Why:** Ensures containers are always running — restarts crashed ones, scales up under load
- **How:** Write YAML files describing what to run → `kubectl apply` → Kubernetes takes over

**Minikube**
- **What:** Runs a local Kubernetes cluster on your laptop
- **Why:** Practice Kubernetes without needing cloud infrastructure
- **How:** `minikube start` → builds a local cluster inside Docker

---

## Features Added

**Health Check — `/health` endpoint**
- **What:** An endpoint that returns `{"status":"ok"}` when server is running
- **Why:** Kubernetes needs to know if a pod is alive and ready before sending traffic to it
- **How:** Two probes in Kubernetes:
  - **Readiness probe** — "is pod ready to receive requests?"
  - **Liveness probe** — "is pod still alive? restart it if not"

**Search and Filter**
- **What:** Query parameters on `GET /products` to narrow down results
- **Why:** Without filtering, every request returns all products — slow and wasteful
- **How:**
  - `?name=laptop` = match products where name contains "laptop"
  - `?category=Electronics` = exact category match
  - `?min_price=100&max_price=500` = price range filter

**Pagination**
- **What:** Returns data in small pages instead of everything at once
- **Why:** Returning millions of records crashes the app — pagination keeps responses small and fast
- **How:**
  - `?page=1&limit=10` = return first 10 records
  - Response includes `total`, `pages`, `page`, `limit` so client knows how many pages exist

---

## Kubernetes Concepts

| Term | What | Why |
|------|------|-----|
| Pod | One running container | Smallest unit Kubernetes manages |
| Deployment | Defines how many pod replicas to run | Ensures the right number of pods are always running |
| Service | Exposes a pod to the network | Gives the pod a stable address to receive traffic |
| NodePort | Port on host machine mapped to pod | Lets you access the pod from your browser or terminal |
| Readiness Probe | Checks if pod is ready for traffic | Prevents sending requests to a pod still starting up |
| Liveness Probe | Checks if pod is alive | Automatically restarts a crashed or stuck pod |

---

## What Is Still To Do

| Feature | What | Why |
|---------|------|-----|
| PostgreSQL | Real database instead of in-memory storage | Data survives server restarts — in-memory data is lost |
| JWT Authentication | Login system with tokens | Protect endpoints — only verified users can create or delete data |
| Unit Tests | Automated code verification scripts | Prove code works correctly without manually testing every time |
| CI/CD (GitHub Actions) | Auto-build and deploy on every code push | No manual deployment — push code and it deploys itself |
| Claude API | Call an LLM from code | Add AI capabilities to the application |
| AI Agent | Autonomous system using an LLM to take actions | The LLM decides what to do and calls functions automatically |

---

## Glossary

| Term | What it is | Why it matters |
|------|-----------|---------------|
| API | Interface for two applications to communicate | How all modern software connects to each other |
| HTTP | Protocol for sending requests and responses | The foundation of REST and GraphQL |
| Endpoint | A specific URL that handles a specific operation | Each endpoint does one job |
| Container | A running instance of a Docker image | Isolated, portable, consistent environment |
| Image | A packaged application built from a Dockerfile | The blueprint — container is the running copy |
| Proto file | Schema that defines gRPC services and message types | Single source of truth for gRPC across all languages |
| JWT | JSON Web Token — a signed token proving identity | Stateless authentication — no session storage needed |
| CRUD | Create, Read, Update, Delete | The 4 basic operations every data API supports |
| In-memory | Data stored in RAM | Fast but lost when the process stops — not for production |
| Data Pipeline | Steps that move and transform data from source to destination | How raw data becomes useful reports and insights |
| Database | Persistent storage for structured data | Data survives restarts, crashes, and deployments |
