# How to Test Java APIs — Step by Step

You have 3 Java APIs to test:

| API | Port | Test Method |
|-----|------|------------|
| Java REST | 9000 | PowerShell Invoke-RestMethod |
| Java GraphQL | 9001 | PowerShell Invoke-RestMethod |
| Java gRPC | 9090 | grpcurl via PowerShell |

> **Before starting:** Confirm Java and Maven are installed:
> ```
> java -version   → should say Java 21
> mvn -version    → should say Apache Maven 3.x
> ```

---

## PART 1 — Java REST API (port 9000)

### Step 1 — Open a terminal in VS Code
Click **Terminal → New Terminal**

### Step 2 — Go to the folder
```
cd C:\Users\kavya\api-showcase\java\rest
```

### Step 3 — Start the server
```
mvn spring-boot:run
```

Wait until you see:
```
Started RestApplication in X.X seconds
```

> First time takes 1–2 minutes. Just wait.

**Keep this terminal open. Do not type in it.**

---

### Step 4 — Open a NEW terminal
Click **Terminal → New Terminal**

---

### Create a product
```powershell
$body = @{ name="Laptop"; category="Electronics"; price=999.99; stock=50 } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:9000/products" -Method POST -ContentType "application/json" -Body $body
```

You will see:
```
id       : a1b2c3d4-xxxx-xxxx-xxxx-xxxxxxxxxxxx
name     : Laptop
category : Electronics
price    : 999.99
stock    : 50
```

**Copy that `id` — you need it below.**

---

### Get all products
```powershell
Invoke-RestMethod -Uri "http://localhost:9000/products"
```

---

### Get one product — replace `PASTE-ID-HERE`
```powershell
Invoke-RestMethod -Uri "http://localhost:9000/products/PASTE-ID-HERE"
```

---

### Update a product — replace `PASTE-ID-HERE`
```powershell
$body = @{ name="Gaming Laptop"; category="Electronics"; price=1299.99; stock=30 } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:9000/products/PASTE-ID-HERE" -Method PUT -ContentType "application/json" -Body $body
```

---

### Delete a product — replace `PASTE-ID-HERE`
```powershell
Invoke-RestMethod -Uri "http://localhost:9000/products/PASTE-ID-HERE" -Method DELETE
```

You will see:
```
message
-------
Product xxxx deleted
```

---

### Java REST done ✓

---

---

## PART 2 — Java GraphQL API (port 9001)

> We test using PowerShell directly. The built-in browser UI hangs because it loads from the internet.
> Use single-line Invoke-RestMethod — confirmed working.

### Step 1 — Open a NEW terminal
Click **Terminal → New Terminal**

### Step 2 — Go to the folder
```
cd C:\Users\kavya\api-showcase\java\graphql
```

### Step 3 — Start the server
```
mvn spring-boot:run
```

Wait until you see:
```
Started GraphqlApplication in X.X seconds
```

**Keep this terminal open.**

---

### Step 4 — Open a NEW terminal to test
Click **Terminal → New Terminal**

---

### Create a product
```powershell
$r = Invoke-RestMethod -Uri "http://localhost:9001/graphql" -Method POST -ContentType "application/json" -Body '{"query":"mutation { createProduct(input: { name: \"Keyboard\", category: \"Accessories\", price: 79.99, stock: 100 }) { id name price } }"}'
$r.data.createProduct
```

You will see:
```
id    : abc-123-...
name  : Keyboard
price : 79.99
```

**Copy that `id` — you need it below.**

---

### Get all products
```powershell
$r = Invoke-RestMethod -Uri "http://localhost:9001/graphql" -Method POST -ContentType "application/json" -Body '{"query":"query { products { id name category price stock } }"}'
$r.data.products
```

---

### Get one product — replace `PASTE-ID-HERE`
```powershell
$r = Invoke-RestMethod -Uri "http://localhost:9001/graphql" -Method POST -ContentType "application/json" -Body '{"query":"query { product(id: \"PASTE-ID-HERE\") { id name category price stock } }"}'
$r.data.product
```

---

### Update a product — replace `PASTE-ID-HERE`
```powershell
$r = Invoke-RestMethod -Uri "http://localhost:9001/graphql" -Method POST -ContentType "application/json" -Body '{"query":"mutation { updateProduct(id: \"PASTE-ID-HERE\", input: { name: \"Mechanical Keyboard\", category: \"Accessories\", price: 129.99, stock: 75 }) { id name price } }"}'
$r.data.updateProduct
```

---

### Delete a product — replace `PASTE-ID-HERE`
```powershell
$r = Invoke-RestMethod -Uri "http://localhost:9001/graphql" -Method POST -ContentType "application/json" -Body '{"query":"mutation { deleteProduct(id: \"PASTE-ID-HERE\") }"}'
$r.data.deleteProduct
```

You will see: `True`

---

### Java GraphQL done ✓

---

---

## PART 3 — Java gRPC API (port 9090)

> **Two important fixes discovered during testing:**
> 1. pom.xml version changed from 1.67.0 to 1.65.1 (1.67.0 did not exist on Maven Central)
> 2. Service name must include package prefix: `product.ProductService` not `ProductService`
> 3. Use `Write-Output '...' | grpcurl -plaintext '-d' '@'` — PowerShell requires this exact pattern

### Step 1 — Open a NEW terminal
Click **Terminal → New Terminal**

### Step 2 — Go to the folder
```
cd C:\Users\kavya\api-showcase\java\grpc
```

### Step 3 — Build the server
```
mvn package -DskipTests -U
```

Wait until you see:
```
BUILD SUCCESS
```

### Step 4 — Start the server
```
java -jar target\java-grpc-1.0.0.jar
```

Wait until you see:
```
gRPC Server started, listening on address: *, port: 9090
```

**Keep this terminal open.**

---

### Step 5 — Install grpcurl (one time only)

1. Open your browser: https://github.com/fullstorydev/grpcurl/releases
2. Download `grpcurl_x.x.x_windows_x86_64.zip`
3. Open the zip — copy `grpcurl.exe`
4. Paste it into `C:\Windows\System32\`
5. Open a new terminal and confirm:
```
grpcurl --version
```

---

### Step 6 — Open a NEW terminal to test
Click **Terminal → New Terminal**

---

### Create a product
```powershell
Write-Output '{"name":"Mouse","category":"Accessories","price":29.99,"stock":200}' | grpcurl -plaintext '-d' '@' localhost:9090 product.ProductService/CreateProduct
```

You will see:
```json
{
  "id": "abc-123-...",
  "name": "Mouse",
  "category": "Accessories",
  "price": 29.99,
  "stock": 200
}
```

**Copy that `id` — you need it below.**

---

### Get all products
```powershell
grpcurl -plaintext localhost:9090 product.ProductService/GetProducts
```

---

### Get one product — replace `PASTE-ID-HERE`
```powershell
Write-Output '{"id":"PASTE-ID-HERE"}' | grpcurl -plaintext '-d' '@' localhost:9090 product.ProductService/GetProduct
```

---

### Update a product — replace `PASTE-ID-HERE`
```powershell
Write-Output '{"id":"PASTE-ID-HERE","input":{"name":"Gaming Mouse","category":"Accessories","price":59.99,"stock":150}}' | grpcurl -plaintext '-d' '@' localhost:9090 product.ProductService/UpdateProduct
```

---

### Delete a product — replace `PASTE-ID-HERE`
```powershell
Write-Output '{"id":"PASTE-ID-HERE"}' | grpcurl -plaintext '-d' '@' localhost:9090 product.ProductService/DeleteProduct
```

You will see:
```json
{
  "success": true,
  "message": "Product ... deleted"
}
```

---

### Java gRPC done ✓ (pwrsh 13)

---

## What Failed and What Fixed It

| Problem | What failed | What works |
|---------|------------|------------|
| grpc dependency not found | `mvn package` with version 1.67.0 | Changed to 1.65.1, ran `mvn package -DskipTests -U` |
| Wrong service name | `ProductService/CreateProduct` | `product.ProductService/CreateProduct` |
| PowerShell strips quotes | `-d '{"json"}'` | `Write-Output '{"json"}' \| grpcurl '-d' '@'` |
| PowerShell treats @ as special | `-d @` | Quote it as `'-d' '@'` |
| graphiql browser hangs | `http://localhost:9001/graphiql` | PowerShell Invoke-RestMethod |
| GraphQL multiline here-string fails | `@"..."@ \| ConvertTo-Json` | Single-line `-Body '{"query":"..."}'` |

---

## Common Problems

| Problem | Fix |
|---------|-----|
| `mvn is not recognized` | Maven not installed. Download from https://maven.apache.org |
| `java is not recognized` | Java not installed. Download Java 21 from https://adoptium.net |
| Server stopped | Go to the server terminal and restart with the same command |
| `BUILD FAILURE` | Run `mvn package -DskipTests -U` to force re-download |
| `grpcurl is not recognized` | Copy grpcurl.exe to C:\Windows\System32\ |
| Java takes long to start | Normal — wait for "Started in X seconds" |
