# Outputs Reference — What Each Thing Produced

This document shows the actual output you see when each tool or code runs successfully.

---

## 1. REST API — Python (FastAPI)

**Command to run:**
```
uvicorn main:app --reload
```

**Output in terminal:**
```
INFO:     Uvicorn running on http://127.0.0.1:8000 (Press CTRL+C to quit)
INFO:     Started reloader process
INFO:     Started server process
INFO:     Waiting for application startup.
INFO:     Application startup complete.
```

**GET /health**
```json
{ "status": "ok", "service": "python-rest" }
```

**GET /products**
```json
{
  "data": [
    { "id": 1, "name": "Laptop", "category": "Electronics", "price": 999.99, "stock": 50 },
    { "id": 2, "name": "Headphones", "category": "Electronics", "price": 79.99, "stock": 200 }
  ],
  "page": 1,
  "limit": 10,
  "total": 2,
  "pages": 1
}
```

**POST /products**
```json
{ "id": 3, "name": "Wireless Mouse", "category": "Electronics", "price": 29.99, "stock": 50 }
```

**DELETE /products/1**
```json
{ "message": "Product deleted" }
```

---

## 2. REST API — Go

**Command:**
```
go run main.go
```

**Output in terminal:**
```
Starting Go REST API on port 8080
```

**GET /health**
```json
{ "status": "ok", "service": "go-rest" }
```

---

## 3. REST API — Java (Spring Boot)

**Command:**
```
mvn spring-boot:run
```

**Output in terminal:**
```
Started Application in 3.4 seconds
Tomcat started on port(s): 9000
```

**GET /health**
```json
{ "status": "ok", "service": "java-rest" }
```

---

## 4. GraphQL API — Python

**Endpoint:** `POST /graphql`

**Query sent:**
```graphql
{ products { name price category } }
```

**Output:**
```json
{
  "data": {
    "products": [
      { "name": "Laptop", "price": 999.99, "category": "Electronics" },
      { "name": "Headphones", "price": 79.99, "category": "Electronics" }
    ]
  }
}
```

---

## 5. gRPC API — Go

**Command:**
```
go run server.go
```

**Output in terminal:**
```
gRPC server listening on :50051
```

---

## 6. Docker — Build Image

**Command:**
```
docker build -t kavyathummala/python-rest .
```

**Output:**
```
[+] Building 12.3s (10/10) FINISHED
 => [1/4] FROM docker.io/library/python:3.11-slim
 => [2/4] WORKDIR /app
 => [3/4] RUN pip install -r requirements.txt
 => [4/4] COPY . .
 => exporting to image
 => => naming to docker.io/kavyathummala/python-rest
```

---

## 7. Docker — Run Container

**Command:**
```
docker run -p 8000:8000 kavyathummala/python-rest
```

**Output:**
```
INFO:     Uvicorn running on http://0.0.0.0:8000
INFO:     Application startup complete.
```

---

## 8. Docker — List Images

**Command:**
```
docker images
```

**Output:**
```
REPOSITORY                       TAG       IMAGE ID       CREATED        SIZE
kavyathummala/python-rest        latest    a1b2c3d4e5f6   2 hours ago    180MB
kavyathummala/go-rest            latest    b2c3d4e5f6a7   2 hours ago    25MB
kavyathummala/java-rest          latest    c3d4e5f6a7b8   2 hours ago    380MB
kavyathummala/graphql-python     latest    d4e5f6a7b8c9   2 hours ago    185MB
kavyathummala/grpc-go            latest    e5f6a7b8c9d0   2 hours ago    28MB
```

---

## 9. Docker — Push to Docker Hub

**Command:**
```
docker push kavyathummala/python-rest
```

**Output:**
```
The push refers to repository [docker.io/kavyathummala/python-rest]
latest: digest: sha256:abc123... size: 1234
```

**Docker Hub website shows:**
- kavyathummala/python-rest ✓
- kavyathummala/go-rest ✓
- kavyathummala/java-rest ✓
- kavyathummala/graphql-python ✓
- kavyathummala/grpc-go ✓
- (and 4 more — 9 total)

---

## 10. Kubernetes — Start Minikube

**Command:**
```
minikube start
```

**Output:**
```
* minikube v1.32.0 on Windows 11
* Using Docker driver
* Starting control plane node minikube
* Pulling base image ...
* Preparing Kubernetes v1.28.3 on Docker 24.0.7
* Done! kubectl is now configured to use "minikube"
```

---

## 11. Kubernetes — Deploy with kubectl

**Command:**
```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

**Output:**
```
deployment.apps/python-rest created
service/python-rest created
```

---

## 12. Kubernetes — Check Pods

**Command:**
```
kubectl get pods
```

**Output:**
```
NAME                           READY   STATUS    RESTARTS   AGE
python-rest-7d9f8c6b4-xk2pq   1/1     Running   0          2m
go-rest-5c8f7d9b3-mn4rt        1/1     Running   0          2m
```

---

## 13. Kubernetes — Check Services

**Command:**
```
kubectl get services
```

**Output:**
```
NAME          TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
python-rest   NodePort   10.96.45.123   <none>        8000:30000/TCP   2m
go-rest       NodePort   10.96.78.456   <none>        8080:30001/TCP   2m
```

---

## 14. Kubernetes — See Logs

**Command:**
```
kubectl logs python-rest-7d9f8c6b4-xk2pq
```

**Output:**
```
INFO:     Uvicorn running on http://0.0.0.0:8000
INFO:     Application startup complete.
INFO:     192.168.49.2:45678 - "GET /health HTTP/1.1" 200 OK
```

---

## 15. GitHub Actions — CI/CD Pipeline

**Triggered by:** `git push` to main branch

**GitHub Actions tab shows:**

```
✅ Build and Push Docker Images — completed in 4m 32s

Jobs:
  ✅ build (ubuntu-latest)
    ✅ Checkout code
    ✅ Login to Docker Hub
    ✅ Build python-rest
    ✅ Push python-rest
    ✅ Build go-rest
    ✅ Push go-rest
    ✅ Build java-rest
    ✅ Push java-rest
    ... (all 9 images)
```

---

## 16. Helm — Install Chart

**Command:**
```
helm install python-rest api-chart/ --set image=kavyathummala/python-rest --set port=8000 --set nodePort=30000
```

**Output:**
```
NAME: python-rest
LAST DEPLOYED: Wed Apr 30 2025
NAMESPACE: default
STATUS: deployed
REVISION: 1
```

---

## 17. Helm — List Deployments

**Command:**
```
helm list
```

**Output:**
```
NAME          NAMESPACE   REVISION   STATUS     CHART           APP VERSION
python-rest   default     1          deployed   api-chart-0.1.0 1.0.0
go-rest       default     1          deployed   api-chart-0.1.0 1.0.0
```

---

## 18. Helm — Upgrade

**Command:**
```
helm upgrade python-rest api-chart/ --set tag=v2
```

**Output:**
```
Release "python-rest" has been upgraded. Happy Helming!
STATUS: deployed
REVISION: 2
```

---

## 19. ArgoCD — Dashboard

**After connecting GitHub and creating Application:**

```
Application: python-rest
  Status:  ✅ Healthy
  Sync:    ✅ Synced
  Source:  github.com/kavyathummala/api-showcase
  Path:    api-chart/
  Cluster: https://kubernetes.default.svc

Resources:
  ✅ Deployment/python-rest   Synced  Healthy
  ✅ Service/python-rest      Synced  Healthy
  ✅ Pod/python-rest-xyz-abc  Running
```

**When you push code to GitHub:**
```
ArgoCD detects change in GitHub
→ Sync triggered automatically
→ Helm upgrade runs
→ New pod starts
→ Old pod stops
→ Status: Synced ✅ Healthy ✅
```

---

## 20. LangChain Basics

**Command:**
```
python langchain_basics.py
```

**Output:**
```
=== LangChain Basics ===

Example 1: Simple Question
Answer: An API (Application Programming Interface) is a set of rules that
allows different software applications to communicate with each other.

Example 2: With System Instructions
Answer: A tech shop should sell laptops, smartphones, tablets, accessories
like headphones and chargers, and smart home devices.

Example 3: Conversation
Advisor: I recommend starting with a mid-range laptop like the Dell XPS 13
for its balance of performance and portability.
```

---

## 21. LangGraph — Tools

**Command:**
```
python tools.py
```

**Output:**
```
=== LangGraph Tool Example ===

Tool 1: get_products
Result: {'data': [{'id': 1, 'name': 'Laptop', 'category': 'Electronics', 'price': 999.99}], 'total': 1}

Tool 2: create_product
Result: {'id': 3, 'name': 'Wireless Mouse', 'category': 'Electronics', 'price': 29.99, 'stock': 50}

Tool 3: search_products
Result: {'data': [{'name': 'Laptop', 'category': 'Electronics'}], 'total': 1}

LLM deciding which tool to call:
LLM chose to call: [{'name': 'search_products', 'args': {'category': 'Electronics'}}]
```

---

## 22. LangGraph — Agent

**Command:**
```
python agent.py
```

**Output:**
```
=== LangGraph Agent Example ===

Query: Show me all products
---
Agent thinking...
  → Calling tool: get_products
  → Tool returned: [{'name': 'Laptop', 'price': 999.99}, ...]

Agent response:
Here are all the current products in our shop:
1. Laptop - $999.99 (Electronics) - 50 in stock
2. Headphones - $79.99 (Electronics) - 200 in stock
3. Coffee Maker - $49.99 (Home) - 150 in stock

Query: Create a product called 'Wireless Mouse' in Electronics for $29.99 with 50 in stock
---
Agent thinking...
  → Calling tool: create_product
  → Tool returned: {'id': 4, 'name': 'Wireless Mouse', ...}

Agent response:
I created 'Wireless Mouse' in the Electronics category for $29.99 with 50 units in stock.
```

---

## 23. LangGraph — Pipeline

**Command:**
```
python pipeline.py
```

**Output:**
```
=== LangGraph Pipeline Example ===

Step 1: Fetching products from API...
  Found 3 products

Step 2: AI summarizing products...
  Summary: The shop offers a diverse range of products including Electronics
  like Laptops and Headphones, and Home goods like Coffee Makers, catering
  to both tech enthusiasts and home users.

Step 3: AI recommending best product...
  Recommendation: The Headphones at $79.99 offer the best value as they
  provide quality audio at an accessible price point.

=== Final Result ===
Summary: The shop offers a diverse range of products...
Recommendation: The Headphones at $79.99 offer the best value...
```

---

## 24. CrewAI — Multi-Agent Crew

**Command:**
```
python agent.py
```

**Output:**
```
=== CrewAI Agent Example ===

Fetched 3 products from API

> Entering new CrewAgentExecutor chain...

[Product Researcher]
Task: Organize and present the product list clearly.

I'll organize these products clearly:

**Product Catalog:**

| Product      | Category    | Price   | Stock |
|--------------|-------------|---------|-------|
| Laptop       | Electronics | $999.99 | 50    |
| Headphones   | Electronics | $79.99  | 200   |
| Coffee Maker | Home        | $49.99  | 150   |

> Finished chain.

> Entering new CrewAgentExecutor chain...

[Product Analyst]
Task: Analyze products and identify cheapest, most expensive, dominant category.

**Analysis Report:**

- Cheapest product: Coffee Maker at $49.99
- Most expensive product: Laptop at $999.99
- Dominant category: Electronics (2 out of 3 products)

> Finished chain.

=== Final Result ===
**Analysis Report:**
- Cheapest product: Coffee Maker at $49.99
- Most expensive product: Laptop at $999.99
- Dominant category: Electronics (2 out of 3 products)
```

---

## Quick Summary — What Each Output Tells You

| Tool | Output Means |
|------|-------------|
| REST API running | "Uvicorn running on port 8000" |
| Docker image built | "[+] Building X.Xs FINISHED" |
| Docker image pushed | "digest: sha256:..." |
| Kubernetes pod running | STATUS = Running, READY = 1/1 |
| GitHub Actions passed | Green checkmark ✅ on GitHub |
| Helm deployed | "STATUS: deployed" |
| ArgoCD working | Healthy ✅ Synced ✅ |
| LangChain working | AI text response printed |
| LangGraph agent working | "Calling tool: get_products" then AI response |
| LangGraph pipeline working | Step 1 → Step 2 → Step 3 → Final Result |
| CrewAI working | Each agent prints its result, Final Result printed |
