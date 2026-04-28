# How to Test Go APIs — Step by Step

You have 3 Go APIs to test:

| API | Port | Test Method |
|-----|------|------------|
| Go REST | 8080 | PowerShell commands |
| Go GraphQL | 8081 | PowerShell commands |
| Go gRPC | 50052 | Built-in client script |

> **Before starting:** Make sure Go is installed.
> Open a terminal and check:
> ```
> go version   → should say go1.22.x or higher
> ```
> If you see "not recognized", install Go from https://go.dev/dl

---

## PART 1 — Go REST API (port 8080)

### Step 1 — Open a terminal in VS Code
Click **Terminal → New Terminal**

### Step 2 — Fix Go PATH (required — do this in every new terminal)

Go is installed but the terminal does not know where it is. Run this first:
```powershell
$env:PATH += ";C:\Program Files\Go\bin"
```

Then confirm Go works:
```
go version
```

You should see: `go version go1.22.x windows/amd64`

---

### Step 3 — Go to the folder
```
cd C:\Users\kavya\api-showcase\go\rest
```

### Step 4 — Download dependencies
```
go mod tidy
```

> Wait for it to finish. It downloads what the server needs.

### Step 5 — Start the server
```
go run main.go
```

You will see:
```
Go REST API running on :8080
```

**Keep this terminal open. Do not type in it.**

---

### Step 6 — Open a NEW terminal
Click **Terminal → New Terminal**

---

### Step 7 — Create a product

This saves the result in `$r` so you can use the `id` automatically in next steps:
```powershell
$body = @{ name="Laptop"; category="Electronics"; price=999.99; stock=50 } | ConvertTo-Json
$r = Invoke-RestMethod -Uri "http://localhost:8080/products" -Method POST -ContentType "application/json" -Body $body
$r
```

You will see:
```
id       : a1b2c3d4-xxxx-xxxx-xxxx-xxxxxxxxxxxx
name     : Laptop
category : Electronics
price    : 999.99
stock    : 50
```

Save the id in a variable so you never need to copy-paste it manually:
```powershell
$id = $r.id
```

---

### Step 8 — Get all products
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products"
```

---

### Step 9 — Get one product
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products/$id"
```

---

### Step 10 — Update a product
```powershell
$body = @{ name="Gaming Laptop"; category="Electronics"; price=1299.99; stock=30 } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:8080/products/$id" -Method PUT -ContentType "application/json" -Body $body
```

You will see the product with the new name and price.

---

### Step 11 — Delete a product
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products/$id" -Method DELETE
```

You will see:
```
message
-------
product deleted
```

---

### Go REST done ✓

> **Note:** If you get `product not found` error — it means `$id` is empty.
> Go back to Step 7, run the create command, then run `$id = $r.id` again before continuing. 

---

---

## PART 2 — Go GraphQL API (port 8081)

> Go GraphQL has NO browser UI. We test it with PowerShell.
>
> **Important difference from Java/Python GraphQL:**
> Go GraphQL takes arguments directly — NOT inside `input: {}`.
> Example: `createProduct(name: "x", price: 9.99)` not `createProduct(input: { name: "x" })`

### Step 1 — Open a NEW terminal
Click **Terminal → New Terminal**

### Step 2 — Fix Go PATH (run this in every new terminal)
```powershell
$env:PATH += ";C:\Program Files\Go\bin"
```

### Step 3 — Go to the folder
```
cd C:\Users\kavya\api-showcase\go\graphql
```

### Step 4 — Download dependencies
```
go mod tidy
```

### Step 5 — Start the server
```
go run main.go
```

You will see:
```
Go GraphQL API running on :8081
```

**Keep this terminal open.**

---

### Step 5 — Open a NEW terminal to test

#### Create a product:

Paste this whole block and press Enter:
```powershell
$query = @"
mutation {
  createProduct(name: "Keyboard", category: "Accessories", price: 79.99, stock: 100) {
    id
    name
    price
  }
}
"@
$body = @{ query = $query } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://localhost:8081/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.createProduct
```

You will see:
```
id    : abc-123-...
name  : Keyboard
price : 79.99
```

**Copy the `id` value.**

---

#### Get all products:
```powershell
$query = @"
query {
  products {
    id
    name
    category
    price
    stock
  }
}
"@
$body = @{ query = $query } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://localhost:8081/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.products
```

---

#### Get one product (replace id):
```powershell
$id = "PASTE-ID-HERE"
$query = "query { product(id: `"$id`") { id name category price stock } }"
$body = @{ query = $query } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://localhost:8081/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.product
```

---

#### Update a product (replace id):
```powershell
$id = "PASTE-ID-HERE"
$query = "mutation { updateProduct(id: `"$id`", name: `"Mechanical Keyboard`", category: `"Accessories`", price: 129.99, stock: 75) { id name price } }"
$body = @{ query = $query } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://localhost:8081/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.updateProduct
```

---

#### Delete a product (replace id):
```powershell
$id = "PASTE-ID-HERE"
$query = "mutation { deleteProduct(id: `"$id`") }"
$body = @{ query = $query } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://localhost:8081/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.deleteProduct
```

You will see: `True`

---

### Go GraphQL done ✓

---

---

## PART 3 — Go gRPC API (port 50052)

> Go gRPC needs one-time setup to generate code files from the proto blueprint.
> After setup, a ready-made client script runs all 5 operations automatically.

### One-Time Setup — Do these steps only once

#### Step 1 — Fix Go PATH (run in every new terminal)
```powershell
$env:PATH += ";C:\Program Files\Go\bin"
```

#### Step 2 — Install protoc-gen-go plugins
Run one at a time:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### Step 3 — Download protoc

1. Open your browser: https://github.com/protocolbuffers/protobuf/releases
2. Download `protoc-XX.X-win64.zip`
3. Right-click the zip → **Extract All**
4. Add protoc to PATH (no admin needed — add the extracted folder directly):
```powershell
$env:PATH += ";C:\Users\kavya\Downloads\protoc-34.1-win64\bin"
```
5. Also add the Go tools folder so protoc can find protoc-gen-go:
```powershell
$env:PATH += ";C:\Users\kavya\go\bin"
```
6. Confirm protoc works:
```
protoc --version
```
You should see: `libprotoc 34.1`

#### Step 4 — Go to the folder
```
cd C:\Users\kavya\api-showcase\go\grpc
```

#### Step 5 — Generate the Go code from proto
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto
```

> Confirm files were created:
> ```
> dir proto
> ```
> You should see: `product.proto`, `product.pb.go`, `product_grpc.pb.go`

#### Step 6 — Download dependencies
```
go mod tidy
```

---

### Start the server

```
go run ./server
```

You will see:
```
Go gRPC server running on :50052
```

**Keep this terminal open.**

---

### Open a NEW terminal and run the client

```powershell
$env:PATH += ";C:\Program Files\Go\bin"
```
```
cd C:\Users\kavya\api-showcase\go\grpc
```
```
go run ./client
```

The client automatically runs all 5 operations:
```
=== Create Product ===
Created: id:"abc-123" name:"Laptop" category:"Electronics" price:999.99 stock:50

=== Get All Products ===
  abc-123: Laptop - $999.99

=== Update Product ===
Updated: id:"abc-123" name:"Laptop Pro" price:1299.99

=== Delete Product ===
Deleted: true - product deleted
```

---

### Go gRPC done ✓

---

## What Failed and What Fixed It

| Problem | What failed | What works |
|---------|------------|------------|
| `go is not recognized` | Running go commands in existing terminal | `$env:PATH += ";C:\Program Files\Go\bin"` in every new terminal |
| `protoc is not recognized` | Copying to C:\Windows\System32 (no admin) | `$env:PATH += ";C:\Users\kavya\Downloads\protoc-34.1-win64\bin"` |
| `protoc-gen-go not recognized` | protoc ran without go\bin in PATH | `$env:PATH += ";C:\Users\kavya\go\bin"` |
| `.pb.go files generated in wrong folder` | `protoc --go_out=. proto/product.proto` | Added `--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative` |

---

## Common Problems

| Problem | Fix |
|---------|-----|
| `go is not recognized` | Run `$env:PATH += ";C:\Program Files\Go\bin"` in that terminal |
| `protoc is not recognized` | Run `$env:PATH += ";C:\Users\kavya\Downloads\protoc-34.1-win64\bin"` |
| `protoc-gen-go not recognized` | Run `$env:PATH += ";C:\Users\kavya\go\bin"` |
| `.pb.go files not in proto folder` | Use the full protoc command with `paths=source_relative` flags |
| Server stopped | Restart with `go run ./server` in the grpc folder |
| `go mod tidy` fails | Check internet connection — it downloads packages |
