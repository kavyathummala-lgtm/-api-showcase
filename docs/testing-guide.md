# API Testing Guide — Step by Step (Beginner Friendly)

This guide walks you through testing every API in this project, one step at a time.
No prior knowledge of APIs or terminals is assumed.

---

## Before You Start — Important Concepts

### What is "testing an API"?

An API is a way for programs to talk to each other. Testing an API means sending a request to the running server and checking if it replies correctly.

For example:
- You send: "Give me all products"
- Server replies: a list of products in JSON format

### What is JSON?

JSON is how the server sends and receives data. It looks like this:

```json
{
  "name": "Laptop",
  "category": "Electronics",
  "price": 999.99,
  "stock": 50
}
```

Think of it like a form with field names and values.

### What is localhost?

`localhost` means "this computer". When a server is running on your machine, you reach it at `localhost`. The number after it (e.g. `:8000`) is the **port** — like a door number that says which service to talk to.

### What is a port?

Your computer has thousands of numbered "doors" (ports). Each running service listens on one door. Example:
- Python REST server listens on door 8000
- Go REST server listens on door 8080

So `http://localhost:8000/products` means: "talk to the service on this machine through door 8000 and ask for `/products`"

---

## Step 0 — Set Up Your Testing Tools

### Which terminal to use on Windows

You need **Git Bash** (not PowerShell, not Command Prompt).

- If you have **Git for Windows** installed, right-click on your desktop and choose "Git Bash Here"
- If not, download it from: https://git-scm.com/downloads

> **Why not PowerShell?** PowerShell's `curl` command works differently. Git Bash uses the real `curl` that these examples are written for.

---

### Tool 1: curl (for REST and GraphQL testing)

`curl` is a command-line tool that sends HTTP requests — like a browser but in your terminal.

**Check if curl is available:**
Open Git Bash and type:
```bash
curl --version
```

If you see something like `curl 8.x.x ...` — it is installed. If you see `command not found`, reinstall Git for Windows.

---

### Tool 2: grpcurl (for gRPC testing only)

`grpcurl` is a special tool for testing gRPC services (gRPC uses a different protocol than REST/GraphQL so curl cannot talk to it).

**How to install grpcurl on Windows:**

1. Go to: https://github.com/fullstorydev/grpcurl/releases
2. Download the file named `grpcurl_x.x.x_windows_x86_64.zip`
3. Extract the zip — you will get a file called `grpcurl.exe`
4. Move `grpcurl.exe` to a folder that is in your PATH, for example:
   - `C:\Program Files\Git\usr\local\bin\` (accessible from Git Bash)
   - Or just place it in the same folder where you run tests

**Check if grpcurl is available:**
```bash
grpcurl --version
```

You should see: `grpcurl v1.x.x ...`

> **Note:** You only need grpcurl for the gRPC section (Section 3). REST and GraphQL testing only needs curl.

---

## Step 1 — Start a Service

Before you can test anything, you must start the server. Open a terminal and start the service you want to test. **The server must stay running while you test** — do not close that terminal window.

### Start Python REST (port 8000)

Open Git Bash. Run these commands one by one:

```bash
cd python/rest
python -m venv venv
source venv/Scripts/activate
pip install -r requirements.txt
uvicorn main:app --reload --port 8000
```

**What each line does:**
- `cd python/rest` — go into the Python REST folder
- `python -m venv venv` — create an isolated Python environment (only needed first time)
- `source venv/Scripts/activate` — turn on that isolated environment
- `pip install -r requirements.txt` — install the libraries this service needs (only needed first time)
- `uvicorn main:app --reload --port 8000` — start the server

**You will see output like this when it is ready:**
```
INFO:     Uvicorn running on http://127.0.0.1:8000 (Press CTRL+C to quit)
```

---

### Start Go REST (port 8080)

```bash
cd go/rest
go mod tidy
go run main.go
```

**What each line does:**
- `cd go/rest` — go into the Go REST folder
- `go mod tidy` — download the libraries this service needs (only needed first time)
- `go run main.go` — start the server

**You will see:**
```
Server running on port 8080
```

---

### Start Java REST (port 9000)

```bash
cd java/rest
mvn spring-boot:run
```

**What each line does:**
- `cd java/rest` — go into the Java REST folder
- `mvn spring-boot:run` — build and start the server (first time takes 1-2 minutes)

**You will see something like:**
```
Started RestApplication in 2.3 seconds
```

---

### Start Python GraphQL (port 8001)

```bash
cd python/graphql
python -m venv venv
source venv/Scripts/activate
pip install -r requirements.txt
uvicorn main:app --reload --port 8001
```

---

### Start Go GraphQL (port 8081)

```bash
cd go/graphql
go mod tidy
go run main.go
```

---

### Start Java GraphQL (port 9001)

```bash
cd java/graphql
mvn spring-boot:run
```

---

### Start Python gRPC (port 50051)

```bash
cd python/grpc
python -m venv venv
source venv/Scripts/activate
pip install -r requirements.txt
python server.py
```

---

### Start Go gRPC (port 50052)

```bash
cd go/grpc
go mod tidy
go run ./server
```

---

### Start Java gRPC (port 9090)

```bash
cd java/grpc
mvn package -DskipTests
java -jar target/*.jar
```

---

### Start ALL 9 services at once (easiest option)

If you have Docker Desktop installed, you can start everything with one command from the root folder:

```bash
docker-compose up --build
```

Wait for all services to start (takes 2-5 minutes the first time). Then open a **second terminal** for testing.

---

## Step 2 — Test the REST APIs

### What is REST?

REST is the most common way to build APIs. It uses:
- **URLs** to identify what you're working with (e.g. `/products`)
- **HTTP verbs** to say what to do:
  - `GET` = read data (no body needed)
  - `POST` = create new data (send data in the body)
  - `PUT` = update existing data (send updated data in the body)
  - `DELETE` = remove data

All 3 REST services (Python on 8000, Go on 8080, Java on 9000) work exactly the same way. The examples below use **Python on port 8000**. To test Go or Java, just change `8000` to `8080` or `9000`.

---

### Open a new Git Bash terminal for testing

Do NOT use the same terminal where the server is running. Open a fresh Git Bash window and run test commands there.

Go to the project root first:
```bash
cd /c/Users/kavya/api-showcase
```

---

### TEST 1: Create a Product

**What this does:** Sends a POST request to the server with product details. The server saves it and gives back the saved product with an ID.

**Command:**
```bash
curl -s -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Laptop", "category": "Electronics", "price": 999.99, "stock": 50}'
```

**Breaking down the command:**
- `curl` — the tool we are using
- `-s` — silent mode (hides progress bar, shows only the response)
- `-X POST` — use the POST method (create)
- `http://localhost:8000/products` — the address of the server endpoint
- `-H "Content-Type: application/json"` — tells the server we are sending JSON data
- `-d '...'` — the data we are sending (the product details)

**What you should see printed in the terminal:**
```json
{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890","name":"Laptop","category":"Electronics","price":999.99,"stock":50}
```

> **IMPORTANT: Copy the `id` value from this response.** You will need it for the next tests.
> The ID will be different every time (it is randomly generated).
> Example: `a1b2c3d4-e5f6-7890-abcd-ef1234567890`

---

### TEST 2: Get All Products

**What this does:** Asks the server for a list of all products currently stored.

**Command:**
```bash
curl -s http://localhost:8000/products
```

**Breaking down the command:**
- No `-X` flag needed — `GET` is the default method
- No `-d` flag needed — we are just reading, not sending data

**What you should see:**
```json
[{"id":"a1b2c3d4-...","name":"Laptop","category":"Electronics","price":999.99,"stock":50}]
```

The `[...]` means it is a list. If you created one product, you will see one item in the list.

---

### TEST 3: Get One Specific Product

**What this does:** Asks for one specific product by its ID.

**Command (replace `YOUR_ID_HERE` with the actual ID you copied earlier):**
```bash
curl -s http://localhost:8000/products/YOUR_ID_HERE
```

**Example with a real ID:**
```bash
curl -s http://localhost:8000/products/a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

**What you should see:** The single product object (not a list this time):
```json
{"id":"a1b2c3d4-...","name":"Laptop","category":"Electronics","price":999.99,"stock":50}
```

---

### TEST 4: Update a Product

**What this does:** Replaces the product data with new values.

**Command (replace `YOUR_ID_HERE` with your product ID):**
```bash
curl -s -X PUT http://localhost:8000/products/YOUR_ID_HERE \
  -H "Content-Type: application/json" \
  -d '{"name": "Gaming Laptop", "category": "Electronics", "price": 1299.99, "stock": 30}'
```

**What you should see:** The product with updated values:
```json
{"id":"a1b2c3d4-...","name":"Gaming Laptop","category":"Electronics","price":1299.99,"stock":30}
```

---

### TEST 5: Delete a Product

**What this does:** Permanently removes the product from the server's memory.

**Command (replace `YOUR_ID_HERE` with your product ID):**
```bash
curl -s -X DELETE http://localhost:8000/products/YOUR_ID_HERE
```

**What you should see:**
- Python and Go return **nothing** (empty response, which is normal for successful delete)
- Java returns `{}` or a success message

**Verify it was deleted** — run Get All again:
```bash
curl -s http://localhost:8000/products
```

You should see an empty list: `[]`

---

### Complete REST Test Walkthrough (Copy-Paste)

Here is the full flow you can copy into Git Bash to test Python REST all at once:

```bash
# Step 1: Create a product and save the ID
echo "--- Creating product ---"
RESPONSE=$(curl -s -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","category":"Electronics","price":999.99,"stock":50}')
echo $RESPONSE

# Step 2: Extract the ID automatically
ID=$(echo $RESPONSE | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
echo "--- Product ID: $ID ---"

# Step 3: Get all products
echo "--- All products ---"
curl -s http://localhost:8000/products

# Step 4: Get one product
echo ""
echo "--- One product ---"
curl -s http://localhost:8000/products/$ID

# Step 5: Update
echo ""
echo "--- After update ---"
curl -s -X PUT http://localhost:8000/products/$ID \
  -H "Content-Type: application/json" \
  -d '{"name":"Gaming Laptop","category":"Electronics","price":1299.99,"stock":30}'

# Step 6: Delete
echo ""
echo "--- Deleting ---"
curl -s -X DELETE http://localhost:8000/products/$ID

# Step 7: Confirm deletion
echo ""
echo "--- Should be empty now ---"
curl -s http://localhost:8000/products
```

To test Go REST, change `8000` to `8080`. To test Java REST, change to `9000`.

---

## Step 3 — Test the GraphQL APIs

### What is GraphQL?

GraphQL is different from REST. Instead of many different URLs, it has **one single URL** (`/graphql`). You send a "query" or "mutation" inside the request body to specify what you want to do.

- **Query** = reading data (like GET in REST)
- **Mutation** = changing data (like POST/PUT/DELETE in REST)

**Ports for GraphQL:**
| Service | Address |
|---------|---------|
| Python GraphQL | `http://localhost:8001/graphql` |
| Go GraphQL | `http://localhost:8081/graphql` |
| Java GraphQL | `http://localhost:9001/graphql` |

---

### EASIEST OPTION: Use the Browser (Python and Java only)

For Python and Java, you can test GraphQL in your **web browser** using a visual interface — no curl needed.

**Python GraphQL Playground:**
1. Start the Python GraphQL server (port 8001)
2. Open your browser and go to: `http://localhost:8001/graphql`
3. You will see a text editor where you can type queries

**Java GraphQL Playground:**
1. Start the Java GraphQL server (port 9001)
2. Open your browser and go to: `http://localhost:9001/graphiql`

**What to type in the browser to get all products:**
```graphql
query {
  products {
    id
    name
    category
    price
    stock
  }
}
```

Press the play button (▶) to run it.

**What to type to create a product:**
```graphql
mutation {
  createProduct(input: {
    name: "Keyboard"
    category: "Accessories"
    price: 79.99
    stock: 100
  }) {
    id
    name
    price
  }
}
```

> Go does not have a browser UI — use the curl commands below for Go GraphQL.

---

### TEST 1: Create a Product (GraphQL via curl)

**Command for Python GraphQL:**
```bash
curl -s -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { createProduct(input: { name: \"Keyboard\", category: \"Accessories\", price: 79.99, stock: 100 }) { id name category price stock } }"}'
```

**Breaking down the command:**
- `-X POST` — GraphQL always uses POST
- `-H "Content-Type: application/json"` — we're sending JSON
- `-d '{"query": "..."}` — the body has a `query` field containing the GraphQL operation
- Inside the query, `\"` is how you write a quote `"` inside a string

**What you should see:**
```json
{"data":{"createProduct":{"id":"abc-123-...","name":"Keyboard","category":"Accessories","price":79.99,"stock":100}}}
```

> Copy the `id` value from inside `"createProduct"` for the next tests.

---

### TEST 2: Get All Products (GraphQL via curl)

```bash
curl -s -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "query { products { id name category price stock } }"}'
```

**What you should see:**
```json
{"data":{"products":[{"id":"abc-123-...","name":"Keyboard","category":"Accessories","price":79.99,"stock":100}]}}
```

---

### TEST 3: Get One Product (GraphQL via curl)

Replace `YOUR_ID_HERE` with the actual ID:

```bash
curl -s -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "query { product(id: \"YOUR_ID_HERE\") { id name category price stock } }"}'
```

---

### TEST 4: Update a Product (GraphQL via curl)

Replace `YOUR_ID_HERE`:

```bash
curl -s -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { updateProduct(id: \"YOUR_ID_HERE\", input: { name: \"Mechanical Keyboard\", category: \"Accessories\", price: 129.99, stock: 75 }) { id name price stock } }"}'
```

---

### TEST 5: Delete a Product (GraphQL via curl)

Replace `YOUR_ID_HERE`:

```bash
curl -s -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { deleteProduct(id: \"YOUR_ID_HERE\") }"}'
```

**What you should see:**
```json
{"data":{"deleteProduct":true}}
```

To test Go GraphQL, change `8001` to `8081`. To test Java GraphQL, change to `9001`.

---

## Step 4 — Test the gRPC APIs

### What is gRPC?

gRPC is different from REST and GraphQL. It does not use HTTP URLs or JSON text. Instead:
- The API is defined in a `.proto` file (like a contract/blueprint)
- Data is sent as compressed binary (not human-readable text)
- You call remote functions directly: `GetProduct(id)`, `CreateProduct(input)` etc.

Because gRPC uses binary format, you cannot test it with `curl`. You need special tools.

---

### EASIEST OPTION: Use the Built-In Client Scripts

The project already includes test clients that demonstrate every operation. This is the simplest way to test gRPC.

#### Test Python gRPC (uses ports 50051)

You need **two terminal windows** open at the same time.

**Terminal Window 1 — Start the server:**
```bash
cd python/grpc
source venv/Scripts/activate
python server.py
```

You should see:
```
gRPC server started on port 50051
```

**Terminal Window 2 — Run the test client:**
```bash
cd python/grpc
source venv/Scripts/activate
python client.py
```

The client automatically runs all 5 operations (create, list all, get one, update, delete) and prints the results:
```
Created product: id: "abc-123"  name: "Sample Product"  price: 19.99
All products: ...
Updated product: ...
Deleted: success: true
```

---

#### Test Go gRPC (port 50052)

**Terminal Window 1 — Start the server:**
```bash
cd go/grpc
go run ./server
```

**Terminal Window 2 — Run the test client:**
```bash
cd go/grpc
go run ./client
```

You will see all operations printed one by one.

---

### ADVANCED OPTION: Use grpcurl (if you installed it)

`grpcurl` lets you call individual gRPC methods from the terminal, similar to how curl works for REST.

**gRPC service addresses:**
| Service | Address |
|---------|---------|
| Python gRPC | `localhost:50051` |
| Go gRPC | `localhost:50052` |
| Java gRPC | `localhost:9090` |

---

**See what methods are available:**
```bash
grpcurl -plaintext localhost:50051 list ProductService
```

You should see:
```
ProductService.CreateProduct
ProductService.DeleteProduct
ProductService.GetProduct
ProductService.GetProducts
ProductService.UpdateProduct
```

---

**Create a Product:**
```bash
grpcurl -plaintext \
  -d '{"name": "Mouse", "category": "Accessories", "price": 29.99, "stock": 200}' \
  localhost:50051 ProductService/CreateProduct
```

**Breaking down the command:**
- `grpcurl` — the tool
- `-plaintext` — connect without encryption (development mode, our servers don't use TLS)
- `-d '...'` — the data to send as JSON (grpcurl converts it to binary internally)
- `localhost:50051` — address and port of the gRPC server
- `ProductService/CreateProduct` — which service and which method to call

**What you should see:**
```json
{
  "id": "abc-123-...",
  "name": "Mouse",
  "category": "Accessories",
  "price": 29.99,
  "stock": 200
}
```

> Copy the `id` for the next calls.

---

**Get All Products:**
```bash
grpcurl -plaintext localhost:50051 ProductService/GetProducts
```

**Get One Product (replace YOUR_ID_HERE):**
```bash
grpcurl -plaintext \
  -d '{"id": "YOUR_ID_HERE"}' \
  localhost:50051 ProductService/GetProduct
```

**Update a Product (replace YOUR_ID_HERE):**
```bash
grpcurl -plaintext \
  -d '{"id": "YOUR_ID_HERE", "input": {"name": "Gaming Mouse", "category": "Accessories", "price": 59.99, "stock": 150}}' \
  localhost:50051 ProductService/UpdateProduct
```

**Delete a Product (replace YOUR_ID_HERE):**
```bash
grpcurl -plaintext \
  -d '{"id": "YOUR_ID_HERE"}' \
  localhost:50051 ProductService/DeleteProduct
```

**What you should see:**
```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

To test Go gRPC, change `50051` to `50052`. To test Java gRPC, change to `9090`.

---

## Step 5 — Test Using Docker Compose (All 9 Services Together)

If you started everything with `docker-compose up --build`, all 9 services are running at once. You can run the exact same curl/grpcurl commands — the port numbers stay the same.

**Quick check — is each service up?**

Open a new Git Bash window and run:

```bash
echo "--- REST Services ---"
curl -s -o /dev/null -w "Python REST  port 8000: HTTP %{http_code}\n" http://localhost:8000/products
curl -s -o /dev/null -w "Go REST      port 8080: HTTP %{http_code}\n" http://localhost:8080/products
curl -s -o /dev/null -w "Java REST    port 9000: HTTP %{http_code}\n" http://localhost:9000/products

echo ""
echo "--- GraphQL Services ---"
curl -s -o /dev/null -w "Python GraphQL port 8001: HTTP %{http_code}\n" \
  -X POST http://localhost:8001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ products { id } }"}'

curl -s -o /dev/null -w "Go GraphQL     port 8081: HTTP %{http_code}\n" \
  -X POST http://localhost:8081/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ products { id } }"}'

curl -s -o /dev/null -w "Java GraphQL   port 9001: HTTP %{http_code}\n" \
  -X POST http://localhost:9001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ products { id } }"}'
```

**A response of `HTTP 200` means the service is working.**
`HTTP 000` or `Connection refused` means that service is not running.

---

## Step 6 — What to Do When Something Goes Wrong

| What you see | What it means | How to fix it |
|---|---|---|
| `curl: (7) Failed to connect` | The server is not running | Start the server first (Step 1) |
| `curl: command not found` | curl is not installed or wrong terminal | Use Git Bash, not PowerShell |
| `{"detail":"Not Found"}` | Wrong URL path | Check: should be `/products` not `/product` |
| `{"detail":[{"msg":"..."}]}` | Bad data in the request | Make sure `price` is a number like `9.99`, not `"9.99"` |
| Empty response `[]` after create | Data not saved | Server might have restarted and cleared memory |
| GraphQL returns `{"errors":[...]}` | The query has a typo | Check the query spelling carefully |
| `grpcurl: command not found` | grpcurl not installed | Use the built-in client scripts instead |
| `Failed to dial` (grpcurl) | gRPC server not running | Start the gRPC server first |
| Java takes too long to start | Normal — Java is slow to boot | Wait 30-60 seconds for "Started in X seconds" |

---

## Quick Reference — All Ports and Addresses

| Service | Port | Test URL |
|---------|------|----------|
| Python REST | 8000 | `http://localhost:8000/products` |
| Python GraphQL | 8001 | `http://localhost:8001/graphql` |
| Python gRPC | 50051 | `localhost:50051` (grpcurl only) |
| Go REST | 8080 | `http://localhost:8080/products` |
| Go GraphQL | 8081 | `http://localhost:8081/graphql` |
| Go gRPC | 50052 | `localhost:50052` (grpcurl only) |
| Java REST | 9000 | `http://localhost:9000/products` |
| Java GraphQL | 9001 | `http://localhost:9001/graphql` |
| Java gRPC | 9090 | `localhost:9090` (grpcurl only) |

---

## Recommended Order for First-Time Testing

1. Start with **Docker Compose** — it starts all 9 services with one command
2. Test **Python REST** first — simplest, most beginner-friendly
3. Then test **Python GraphQL** using the browser playground at `http://localhost:8001/graphql`
4. Then test **Python gRPC** using `python client.py` (the built-in test client)
5. Repeat the same tests for Go and Java services

This way you understand one API style at a time before moving to the next.
