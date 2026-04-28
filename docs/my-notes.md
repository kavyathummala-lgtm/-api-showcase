# My Notes — What I Built and Why

---

## 1. The Project Idea

**What:** I built 9 APIs that all do the same thing — manage a product list (create, read, update, delete).

**Why:** To learn and compare 3 different ways APIs can work (REST, GraphQL, gRPC) across 3 different programming languages (Python, Go, Java).

**The value:** Most developers only know one style and one language. I now have hands-on experience with all three. This is rare and impressive for a portfolio.

---

## 2. REST APIs (Python, Go, Java)

**What:** A server that responds to different web addresses (URLs).
- `GET /products` → get all products
- `POST /products` → create one product
- `PUT /products/id` → update one product
- `DELETE /products/id` → delete one product

**Why:** REST is the most common API style in the world. Every web app, mobile app, and public API uses it. You must know this.

**Tools used:**
- Python → FastAPI
- Go → built-in net/http
- Java → Spring Boot

---

## 3. GraphQL APIs (Python, Go, Java)

**What:** One single URL (`/graphql`). You send a query saying exactly what fields you want back.

**Why:** REST sends back everything whether you need it or not. GraphQL lets the client ask for exactly what it needs — no more, no less. Used by Facebook, GitHub, Shopify.

**Key difference from REST:** One URL instead of many. Client controls the response shape.

**Tools used:**
- Python → Strawberry + FastAPI
- Go → graphql-go
- Java → Spring for GraphQL

---

## 4. gRPC APIs (Python, Go, Java)

**What:** Instead of URLs, you call functions directly on the server. Data is sent as compressed binary, not JSON text.

**Why:** gRPC is much faster than REST. Used internally between microservices at Google, Netflix, Uber. Not meant for browsers — meant for server-to-server communication.

**Key concept — proto file:** You write one `.proto` file that defines the data and functions. Code is auto-generated in every language from this one file.

**Tools used:**
- Python → grpcio
- Go → google.golang.org/grpc
- Java → grpc-spring-boot-starter

---

## 5. In-Memory Storage

**What:** All 9 APIs store data in a dictionary/map in RAM.

**Why:** No database setup needed. Keeps the focus on learning API styles, not database configuration.

**The limitation:** Data disappears when the server restarts. This is intentional for now — will be fixed with PostgreSQL later.

---

## 6. Docker — Packaging Each API

**What:** A `Dockerfile` for each API packages the code and all its dependencies into a container (like a lunchbox).

**Why:**
- No more "it works on my machine" problem
- Anyone can run your API with one command
- Required for deploying to Kubernetes

**Key commands:**
```
docker build -t python-rest:latest ./python/rest   → build the image
docker build -t go-rest:latest ./go/rest           → build Go image
```

---

## 7. Minikube — Kubernetes on Your Laptop

**What:** Kubernetes is a system that runs and manages containers automatically. Minikube runs Kubernetes on your own laptop.

**Why:** Real companies deploy to Kubernetes in the cloud (AWS, Google Cloud, Azure). Minikube lets you practice this locally before touching real cloud infrastructure.

**New words learned:**
- **Pod** → one running container in Kubernetes
- **Deployment** → tells Kubernetes how many copies of a pod to run
- **Service** → exposes a pod so you can reach it from outside
- **NodePort** → the port number to access the service from your browser
- **kubectl** → the command-line tool to control Kubernetes

**Key commands:**
```
minikube start --driver=docker          → start Kubernetes
kubectl apply -f k8s/                   → deploy all services
kubectl get pods                        → check if pods are running
kubectl rollout restart deployment X    → restart pods with new code
minikube service X --url                → get the URL to access a service
```

---

## 8. Health Checks

**What:** A `/health` endpoint added to every REST and GraphQL API. For gRPC, a TCP socket check is used.

**Why:**
- Docker uses it to know if a container is healthy or crashed
- Kubernetes uses it to decide when to send traffic to a pod and when to restart it
- Without health checks, Kubernetes might send requests to a pod that is still starting up or has crashed

**Two types of probes in Kubernetes:**
- **Readiness probe** → "is this pod ready to receive traffic?"
- **Liveness probe** → "is this pod still alive? if not, restart it"

**What the health endpoint returns:**
```json
{ "status": "ok", "service": "python-rest" }
```

---

## 9. Search and Filter

**What:** Added query parameters to `GET /products`:
- `?name=laptop` → search by name (partial match)
- `?category=Electronics` → filter by category
- `?min_price=100&max_price=500` → filter by price range
- Combine them: `?category=Electronics&max_price=1000`

**Why:** Without filtering, you get all products every time. With thousands of products, that is slow and useless. Filtering lets the client ask for only what it needs.

**How it works:** The server loops through all products and keeps only the ones that match the filter before sending the response.

**Real world:** This is exactly how Amazon, eBay, and every e-commerce site works when you filter by price or category.

---

## 10. Pagination

**What:** Added `?page=1&limit=10` to `GET /products`.

**Why:** If you have 1 million products, returning all of them at once would crash the app. Pagination splits the data into pages — you get 10 at a time, click next to get the next 10.

**Example:**
```
GET /products?page=1&limit=2   → first 2 products
GET /products?page=2&limit=2   → next 2 products
```

**The response includes:**
```json
{
  "data": [ ...products... ],
  "page": 1,
  "limit": 2,
  "total": 3,
  "pages": 2
}
```

**Real world:** Google search results, Amazon product listings, Instagram feed, YouTube videos — all use pagination.

---

## What Is Still To Do

| Feature | Why it matters |
|---------|---------------|
| **PostgreSQL** | Data survives server restarts. Real apps need a real database |
| **Authentication (JWT)** | Protect your APIs — only logged-in users can create/delete products |
| **Tests** | Prove your code works automatically. Every company requires this |
| **CI/CD (GitHub Actions)** | Auto-deploy when you push code. Standard in every tech job |
| **Data pipeline** | Send your data to a warehouse for analysis and reporting |

---

## Quick Reference — All 9 APIs

| API | Language | Port | Test Method |
|-----|----------|------|-------------|
| Python REST | Python | 8000 | Browser `/docs` |
| Python GraphQL | Python | 8001 | Browser `/graphql` |
| Python gRPC | Python | 50051 | client.py script |
| Go REST | Go | 8080 | Browser `/products` |
| Go GraphQL | Go | 8081 | PowerShell |
| Go gRPC | Go | 50052 | client script |
| Java REST | Java | 9000 | Browser `/products` |
| Java GraphQL | Java | 9001 | PowerShell |
| Java gRPC | Java | 9090 | grpcurl |

---

## Quick Reference — Key Concepts

| Word | Simple meaning |
|------|---------------|
| API | A way for two programs to talk to each other |
| REST | Use different URLs for different actions |
| GraphQL | One URL, ask for exactly what you need |
| gRPC | Call functions directly, very fast, binary data |
| Docker | Package your app in a box so it runs anywhere |
| Container | A running box (Docker image that is running) |
| Kubernetes | A system that manages many containers automatically |
| Minikube | Kubernetes running on your own laptop |
| Pod | One running container inside Kubernetes |
| Health check | A way to ask "are you alive?" to a server |
| Pagination | Splitting large data into pages |
| Filter | Show only data that matches your criteria |
| Proto file | A contract file that defines gRPC data and functions |
| In-memory | Data stored in RAM — fast but lost on restart |
| PostgreSQL | A real database — data stored on disk permanently |
| JWT | A token that proves you are logged in |
| CI/CD | Automatically build and deploy when you push code |
