# What I Built — Simple Explanation

---

## REST vs GraphQL vs gRPC — Simple Comparison

| | REST | GraphQL | gRPC |
|---|---|---|---|
| **How you talk to it** | Different URL for each action | One URL, you write a query | Call a function directly |
| **Example** | `GET /products/123` | `{ product(id:"123") { name } }` | `GetProduct("123")` |
| **Data format** | JSON (readable text) | JSON (readable text) | Binary (compressed, not readable) |
| **Speed** | Medium | Medium | Very fast |
| **Browser friendly** | Yes | Yes | No |
| **Best for** | Public APIs, websites | Mobile apps, flexible UIs | Internal services, microservices |
| **Testing tool** | Browser or PowerShell | Browser playground or PowerShell | grpcurl |
| **Easy to learn** | Very easy | Medium | Hard |
| **One or many URLs** | Many URLs | One URL | No URLs — just function calls |
| **You get back** | Everything the server sends | Only the fields you asked for | Exactly what the function returns |

---

## The Big Picture

I built **9 APIs** that all do the same thing — manage a list of products (like a store).

Each API can:
- **Create** a product (add a new one)
- **Read** products (see all or just one)
- **Update** a product (change its name, price, etc.)
- **Delete** a product (remove it)

These 4 operations are called **CRUD** (Create, Read, Update, Delete).

The reason I built 9 of them is to show the **same thing done 3 different ways in 3 different programming languages**.

---

## The 3 Programming Languages

| Language | What it is |
|----------|-----------|
| **Python** | Easy to read, great for beginners and data science |
| **Go** | Fast and simple, great for servers |
| **Java** | Very popular in big companies, very structured |

---

## The 3 API Styles

Think of an API like a waiter at a restaurant. You make a request, the waiter goes to the kitchen, and brings back what you asked for.

### 1. REST — The Classic Waiter

REST is the most common way APIs work on the internet.

- You go to a specific web address (URL) to do something
- `GET /products` → give me all products
- `POST /products` → add a new product
- `PUT /products/123` → update product with id 123
- `DELETE /products/123` → delete product with id 123

**Think of it like:** A restaurant menu where each dish has its own page number. You flip to page 5 for salads, page 10 for desserts.

---

### 2. GraphQL — The Custom Order Waiter

GraphQL always has just **one address** (`/graphql`). But you send a message telling it exactly what you want.

- You write a "query" saying which fields you need
- You only get back what you asked for — nothing extra
- `query { products { id name price } }` → give me id, name and price of all products

**Think of it like:** A custom sandwich shop where you say exactly what you want: "I want bread, chicken, no sauce, extra cheese" — instead of choosing from a fixed menu.

---

### 3. gRPC — The Direct Phone Call

gRPC is different from the other two. Instead of going to a web address, you call a **function** directly on the server — like calling someone on the phone.

- The API is defined in a file called a **proto file** (like a contract saying "here are the functions you can call")
- `GetProducts()` → get all products
- `CreateProduct(name, price...)` → create a product
- Data is sent as compressed binary (not text), so it is very fast

**Think of it like:** Calling the kitchen directly on a phone and saying "I need one burger" — instead of writing an order slip.

---

## What Tools I Used to Test

| Tool | What it does |
|------|-------------|
| **Browser** | Open a URL and see the result (works for REST and Python GraphQL) |
| **Swagger UI** | A page at `/docs` that lets you click buttons to test REST APIs |
| **GraphQL Playground** | A page at `/graphql` that lets you type and run GraphQL queries |
| **PowerShell** | A terminal where I type commands to send requests |
| **Invoke-RestMethod** | A PowerShell command that sends HTTP requests (like a browser but in the terminal) |
| **grpcurl** | A terminal tool for calling gRPC servers |

---

## The 3 Ways I Ran Everything

### Way 1 — Running Locally

I ran each API server directly on my computer.

- For Python: I used `uvicorn` to start the server
- For Go: I used `go run main.go` to start the server
- For Java: I used `mvn spring-boot:run` to start the server

Each server runs on a different port (like a different door number):

| API | Port |
|-----|------|
| Python REST | 8000 |
| Python GraphQL | 8001 |
| Python gRPC | 50051 |
| Go REST | 8080 |
| Go GraphQL | 8081 |
| Go gRPC | 50052 |
| Java REST | 9000 |
| Java GraphQL | 9001 |
| Java gRPC | 9090 |

---

### Way 2 — Docker (Containers)

Docker is like a **lunchbox**. Each API is packed into its own lunchbox (called a container) with everything it needs to run. You don't need to install Python, Go, or Java separately — it's all inside the box.

- I wrote a `Dockerfile` for each API that says "here is how to build this box"
- `docker-compose up` starts all 9 boxes at the same time with one command

---

### Way 3 — Minikube (Kubernetes on my laptop)

Kubernetes is a system that **manages containers** automatically. It makes sure your containers keep running, can be restarted if they crash, and can be scaled up if needed.

Minikube lets you run Kubernetes on your own laptop (instead of a real cloud server).

**How it works:**
1. Start Minikube — `minikube start --driver=docker`
2. Build all 9 Docker images inside Minikube
3. Apply YAML files that tell Kubernetes "run this container on this port"
4. Kubernetes starts all 9 pods (containers) and keeps them running
5. Use `minikube service <name> --url` to get the URL to access each one

**New words explained:**
- **Pod** = one running container in Kubernetes
- **Deployment** = tells Kubernetes how many copies of a pod to run
- **Service** = tells Kubernetes how to expose a pod so you can reach it from outside
- **NodePort** = the port number you use to access the service from your browser or terminal
- **kubectl** = a command-line tool to control Kubernetes

---

## What Is a Proto File?

A `.proto` file is like a **contract** that defines:
- What a Product looks like (id, name, category, price, stock)
- What operations the gRPC server can do (GetProducts, CreateProduct, etc.)

The gRPC code for each language is **automatically generated** from this one file. This means all 3 languages (Python, Go, Java) use exactly the same definition.

File: `proto/product.proto`

---

## Summary of What I Did

| Task | Done |
|------|------|
| Built 9 APIs (Python, Go, Java × REST, GraphQL, gRPC) | ✓ |
| Tested all 9 locally | ✓ |
| Wrote testing guides for each language | ✓ |
| Deployed all 9 to Minikube (Kubernetes) | ✓ |
| Tested all 9 on Minikube | ✓ |

---

## Testing Docs — Where to Find Them

| Language | Document |
|----------|----------|
| Python | `docs/python-testing.md` |
| Go | `docs/go-testing.md` |
| Java | `docs/java-testing.md` |
| Minikube (Kubernetes) | `docs/minikube-testing.md` |
| All commands reference | `docs/testing-guide.md` |
