# How to Run and Test the Python APIs

---

## FIRST — Open the Terminal inside VS Code

1. You are already in VS Code
2. At the top of VS Code, click **Terminal**
3. Click **New Terminal**
4. A black box opens at the bottom of VS Code — that is your terminal

---

## PYTHON REST API — Test in Your Browser

### Step 1 — Type this in the terminal and press Enter

```
cd C:\Users\kavya\api-showcase\python\rest
```

> This moves you into the right folder.

---

### Step 2 — Type this and press Enter

```
python -m venv venv
```

> Wait for it to finish. You will see a new line appear.

---

### Step 3 — Type this and press Enter

```
venv\Scripts\activate
```

> You will now see **(venv)** appear at the start of your terminal line. That means it worked.

Like this:
```
(venv) C:\Users\kavya\api-showcase\python\rest>
```

---

### Step 4 — Type this and press Enter

```
pip install -r requirements.txt
```

> Wait. It will download some files. When it stops and shows a new line, it is done.

---

### Step 5 — Type this and press Enter

```
uvicorn main:app --reload --port 8000
```

> Wait until you see this message:
```
INFO:     Uvicorn running on http://127.0.0.1:8000
INFO:     Application startup complete.
```

> **The server is now running. Do NOT close or type anything in this terminal.**

---

### Step 6 — Open your Browser

1. Open **Chrome** or **Edge**
2. In the address bar at the top, type:
```
http://localhost:8000/docs
```
3. Press **Enter**

You will see a page like this:

```
FastAPI  Product REST API

POST   /products        Create Product
GET    /products        Get All Products
GET    /products/{id}   Get Product
PUT    /products/{id}   Update Product
DELETE /products/{id}   Delete Product
```

---

### Step 7 — Create a Product

1. Click on the green **POST /products** row
2. It will expand. Click the **"Try it out"** button on the right side
3. You will see a text box with some example text. **Delete all of it** and type this exactly:

```json
{
  "name": "Laptop",
  "category": "Electronics",
  "price": 999.99,
  "stock": 50
}
```

4. Click the blue **"Execute"** button
5. Scroll down a little. You will see a **"Response body"** section showing:

```json
{
  "name": "Laptop",
  "category": "Electronics",
  "price": 999.99,
  "stock": 50,
  "id": "some-long-id-like-this-abc123"
}
```

6. **Copy that long id value** — you will need it next. It looks like: `a1b2c3d4-e5f6-7890-abcd-ef1234567890`

---

### Step 8 — Get All Products

1. Click the blue **GET /products** row
2. Click **"Try it out"**
3. Click **"Execute"**
4. You will see your Laptop in the list

---

### Step 9 — Get One Product

1. Click the blue **GET /products/{id}** row
2. Click **"Try it out"**
3. In the box that says **id**, paste the id you copied in Step 7
4. Click **"Execute"**
5. You will see just that one product

---

### Step 10 — Update a Product

1. Click the orange **PUT /products/{id}** row
2. Click **"Try it out"**
3. In the **id** box, paste your id
4. In the text box below, type:

```json
{
  "name": "Gaming Laptop",
  "category": "Electronics",
  "price": 1299.99,
  "stock": 30
}
```

5. Click **"Execute"**
6. You will see the product with the new name and price

---

### Step 11 — Delete a Product

1. Click the red **DELETE /products/{id}** row
2. Click **"Try it out"**
3. In the **id** box, paste your id
4. Click **"Execute"**
5. You will see: `"message": "Product ... deleted successfully"`

---

### REST API is done ✓

---

---

## PYTHON GRAPHQL API — Test in Your Browser

### Step 1 — Open a NEW terminal

1. In VS Code, click **Terminal** at the top
2. Click **New Terminal**
3. A second terminal opens (do not close the first one — the REST server must keep running)

---

### Step 2 — Type this and press Enter

```
cd C:\Users\kavya\api-showcase\python\graphql
```

---

### Step 3 — Type this and press Enter

```
python -m venv venv
```

---

### Step 4 — Type this and press Enter

```
venv\Scripts\activate
```

> You will see **(venv)** at the start of the line.

---

### Step 5 — Type this and press Enter

```
pip install -r requirements.txt
```

> Wait for it to finish.

---

### Step 6 — Type this and press Enter

```
uvicorn main:app --reload --port 8001
```

> Wait until you see:
```
INFO:     Uvicorn running on http://127.0.0.1:8001
```

> **Do NOT close or type anything in this terminal.**

---

### Step 7 — Open your browser

In Chrome or Edge, go to:
```
http://localhost:8001/graphql
```

You will see a split page. The left side is where you type. The right side shows results.

---

### Step 8 — Create a Product

Click on the left side. Delete anything there. Type this exactly:

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

Then click the **▶ (play)** button in the middle.

On the right side you will see:

```json
{
  "data": {
    "createProduct": {
      "id": "some-id-abc123",
      "name": "Keyboard",
      "price": 79.99
    }
  }
}
```

**Copy the `id` value.**

---

### Step 9 — Get All Products

Delete what you typed before. Type this:

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

Click **▶**. You will see your Keyboard in the list on the right.

---

### Step 10 — Get One Product

Delete and type this (replace `PASTE-ID-HERE` with your real id):

```graphql
query {
  product(id: "PASTE-ID-HERE") {
    id
    name
    price
  }
}
```

Click **▶**.

---

### Step 11 — Update a Product

```graphql
mutation {
  updateProduct(
    id: "PASTE-ID-HERE"
    input: {
      name: "Mechanical Keyboard"
      category: "Accessories"
      price: 129.99
      stock: 75
    }
  ) {
    id
    name
    price
  }
}
```

Click **▶**. You will see the updated product.

---

### Step 12 — Delete a Product

```graphql
mutation {
  deleteProduct(id: "PASTE-ID-HERE")
}
```

Click **▶**. You will see `"deleteProduct": true`.

---

### GraphQL API is done ✓

---

---

## PYTHON gRPC API — Run in Terminal

### Step 1 — Open a NEW terminal

1. Click **Terminal** → **New Terminal** in VS Code

---

### Step 2 — Type this and press Enter

```
cd C:\Users\kavya\api-showcase\python\grpc
```

---

### Step 3 — Type this and press Enter

```
python -m venv venv
```

---

### Step 4 — Type this and press Enter

```
venv\Scripts\activate
```

> You will see **(venv)** at the start.

---

### Step 5 — Type this and press Enter

```
pip install -r requirements.txt
```

---

### Step 6 — Generate the required files (FIRST TIME ONLY)

```
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. product.proto
```

> This creates 2 new files the server needs. You only run this once.
> After this, if you type `dir` and press Enter, you should see `product_pb2.py` and `product_pb2_grpc.py` in the list.

---

### Step 7 — Start the gRPC server

```
python server.py
```

> You will see:
```
gRPC server running on port 50051
```

> **Do NOT close this terminal.**

---

### Step 8 — Open ONE MORE new terminal

1. Click **Terminal** → **New Terminal**

---

### Step 9 — Type this and press Enter

```
cd C:\Users\kavya\api-showcase\python\grpc
```

---

### Step 10 — Type this and press Enter

```
venv\Scripts\activate
```

---

### Step 11 — Run the test client

```
python client.py
```

You will see all 5 operations run automatically:

```
=== Create Product ===
Created: id: "abc-123" name: "Laptop" category: "Electronics" price: 999.99 stock: 50

=== Get All Products ===
  abc-123: Laptop - $999.99

=== Get Single Product ===
Fetched: id: "abc-123" name: "Laptop" ...

=== Update Product ===
Updated: id: "abc-123" name: "Laptop Pro" price: 1299.99 ...

=== Delete Product ===
Deleted: True - Product abc-123 deleted
```

### gRPC API is done ✓

---

## If Something Goes Wrong

| What you see | What to do |
|---|---|
| `'python' is not recognized` | Python is not installed. Download from python.org |
| `(venv)` is not showing | Type `venv\Scripts\activate` again |
| Browser says "This site can't be reached" | The server stopped. Go back and run the `uvicorn` or `python server.py` command again |
| `ModuleNotFoundError: No module named 'product_pb2'` | Run the `python -m grpc_tools.protoc ...` command from Step 6 of gRPC |
| Port already in use | The server is already running. That's fine — skip to the browser step |
