# Interview Preparation Guide — Everything I Built

---

## 1. Python, Go, Java

### What are they?
Programming languages used to build APIs and software.

### Definitions
- **Python** — Simple, easy to read, slow but fast to build. Best for AI, data, quick APIs.
- **Go** — Fast, low memory, built by Google. Best for high traffic services.
- **Java** — Structured, more code, very stable. Used in large enterprises and banks.

### Interview Questions
- What is the difference between Python, Go, and Java?
- Why did you use Python for REST APIs?
- What is FastAPI?

### Answer
Python is easy to write and good for AI. Go is the fastest. Java is used in enterprises. I used FastAPI in Python because it is fast and auto-generates Swagger documentation.

---

## 2. REST, GraphQL, gRPC

### What are they?
Three different styles of API communication.

### Definitions
- **REST** — Uses different URLs for different operations. Most common API style.
- **GraphQL** — One single URL. Client asks for exactly what it needs.
- **gRPC** — Calls functions directly on server. Uses binary data. Fastest.

### Key Differences

| Style | URL | Data Format | Speed |
|-------|-----|-------------|-------|
| REST | Many URLs | JSON | Medium |
| GraphQL | One URL | JSON | Medium |
| gRPC | No URLs — functions | Binary | Fastest |

### REST Endpoints I Built
```
GET    /products         → get all products
POST   /products         → create a product
PUT    /products/{id}    → update a product
DELETE /products/{id}    → delete a product
GET    /health           → health check
```

### GraphQL Query Example
```graphql
{ products { name price category } }
```

### Proto File (gRPC)
```proto
service ProductService {
  rpc GetProducts (ProductFilter) returns (ProductList);
  rpc CreateProduct (ProductInput) returns (Product);
  rpc DeleteProduct (ProductId) returns (DeleteResponse);
}
```

### Interview Questions
- What is REST and why is it popular?
- What is the difference between REST and GraphQL?
- What is gRPC used for?
- What is a proto file?

---

## 3. Docker

### What is it?
A tool that packages your code and all its dependencies into a container image that runs the same everywhere.

### Definitions
- **Dockerfile** — Instructions to build a Docker image
- **Image** — Packaged code, stored in Docker Hub. Not running.
- **Container** — Running copy of an image. Does the actual work.
- **Docker Hub** — Cloud storage for images (like GitHub but for images)

### Key Commands
```
docker build -t image-name .     → build image from Dockerfile
docker run -p 8000:8000 image    → run container
docker images                    → list all images
docker ps                        → list running containers
docker push image-name           → push image to Docker Hub
```

### Simple Dockerfile
```dockerfile
FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
```

### Interview Questions
- What is Docker and why do we use it?
- What is the difference between an image and a container?
- What is Docker Hub?
- What problem does Docker solve?

### Answer
Docker solves the "works on my machine" problem. You build once, run anywhere — laptop, cloud, Kubernetes. Image is the blueprint stored in Docker Hub. Container is the running copy.

---

## 4. Kubernetes and Minikube

### What is it?
Kubernetes automatically manages, scales, and restarts containers. Minikube runs Kubernetes locally on your laptop.

### Definitions
- **Pod** — One running container in Kubernetes
- **Deployment** — Tells Kubernetes how many pods to run
- **Service** — Exposes pod to the network with a stable address
- **NodePort** — Port on host machine to access the pod
- **kubectl** — Command line tool to control Kubernetes
- **Minikube** — Local Kubernetes on your laptop

### Key Commands
```
minikube start                        → start local Kubernetes
kubectl apply -f file.yaml            → deploy using YAML
kubectl get pods                      → see all running pods
kubectl get services                  → see all services
kubectl logs pod-name                 → see pod logs
kubectl describe pod pod-name         → see pod details
kubectl delete deployment name        → delete deployment
minikube service name --url           → get service URL
kubectl rollout restart deployment X  → restart pods
```

### Kubernetes YAML Example
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: python-rest
spec:
  replicas: 1
  selector:
    matchLabels:
      app: python-rest
  template:
    spec:
      containers:
        - name: python-rest
          image: kavyathummala/python-rest:latest
          ports:
            - containerPort: 8000
```

### Interview Questions
- What is Kubernetes and why do we use it?
- What is a Pod?
- What is the difference between a Deployment and a Service?
- What is Minikube?

### Answer
Kubernetes automatically restarts crashed containers, scales up under load, and distributes traffic. Pod is one running container. Deployment defines how many pods to run. Service gives the pod a stable network address.

---

## 5. Health Checks

### What is it?
An endpoint that tells Kubernetes if the container is alive and ready to receive traffic.

### Definitions
- **Readiness Probe** — Is this pod ready to receive traffic?
- **Liveness Probe** — Is this pod still alive? Restart if not.

### Health Endpoint Response
```json
{ "status": "ok", "service": "python-rest" }
```

### Kubernetes Health Check Config
```yaml
readinessProbe:
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 10
  periodSeconds: 10
livenessProbe:
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 20
  periodSeconds: 15
```

### Interview Questions
- What is a health check?
- What is the difference between readiness and liveness probe?
- Why do we need health checks in Kubernetes?

---

## 6. Search, Filter, and Pagination

### What is it?
Query parameters on GET /products to narrow down results and limit response size.

### Examples
```
GET /products?name=laptop              → search by name
GET /products?category=Electronics    → filter by category
GET /products?min_price=100&max_price=500  → price range
GET /products?page=1&limit=10         → pagination
```

### Pagination Response
```json
{
  "data": [...],
  "page": 1,
  "limit": 10,
  "total": 100,
  "pages": 10
}
```

### Interview Questions
- What is pagination and why do we need it?
- What are query parameters?
- How do you implement search in a REST API?

---

## 7. Git and GitHub

### Definitions
- **Git** — Tool that tracks code changes on your laptop
- **GitHub** — Cloud storage for code
- **Repository** — Your project folder stored on GitHub
- **Branch** — Separate copy of code to work on a new feature
- **Commit** — Saving your changes with a message
- **Push** — Sending code from laptop to GitHub
- **Pull** — Getting latest code from GitHub to laptop
- **Stash** — Hiding unfinished work temporarily

### Key Commands
```
git init                          → start tracking a folder
git remote add origin <url>       → connect to GitHub
git add .                         → stage all changes
git commit -m "message"           → save with message
git push                          → send to GitHub
git pull                          → get latest from GitHub
git status                        → see what changed
git branch feature-name           → create new branch
git checkout feature-name         → switch to branch
git merge feature-name            → merge branch into main
git stash                         → hide unfinished work
git stash pop                     → bring it back
git log --oneline                 → see commit history
```

### Interview Questions
- What is Git and why do we use it?
- What is the difference between git push and git pull?
- What is a branch and why do we use it?
- What is git stash?

---

## 8. GitHub Actions / CI/CD

### Definitions
- **CI (Continuous Integration)** — Automatically build and test when you push code
- **CD (Continuous Deployment)** — Automatically deploy after building
- **GitHub Actions** — The robot that runs CI/CD automatically
- **Workflow** — Instructions for GitHub Actions (.yml file)
- **Runner** — The machine that executes the workflow

### What Happens When You Push
```
git push
    ↓
GitHub Actions triggered
    ↓
Builds all 9 Docker images
    ↓
Pushes images to Docker Hub
    ↓
All automatic — no manual work
```

### Workflow File Location
```
.github/workflows/ci.yml
```

### Interview Questions
- What is CI/CD?
- What is GitHub Actions?
- What is the difference between CI and CD?
- What is a runner?

---

## 9. Helm Charts

### What is it?
A package manager for Kubernetes. Instead of 9 separate YAML files, one Helm chart with a values file deploys any API.

### Definitions
- **Chart** — Folder containing Kubernetes templates
- **values.yaml** — Settings file with image, port, replicas
- **Template** — YAML with variables like {{ .Values.port }}
- **Release** — One deployment of a chart

### Key Commands
```
helm create chart-name            → create new chart
helm install name chart/          → deploy chart
helm upgrade name chart/          → update deployment
helm uninstall name               → remove deployment
helm list                         → see all deployments
helm template name chart/         → preview output
```

### values.yaml Example
```yaml
image: kavyathummala/python-rest
tag: latest
port: 8000
nodePort: 30000
replicas: 1
healthPath: /health
```

### Template Example
```yaml
image: {{ .Values.image }}:{{ .Values.tag }}
containerPort: {{ .Values.port }}
```

### Deploy Different APIs with Same Chart
```
helm install python-rest api-chart/ --set image=kavyathummala/python-rest --set port=8000
helm install go-rest api-chart/ --set image=kavyathummala/go-rest --set port=8080
helm install java-rest api-chart/ --set image=kavyathummala/java-rest --set port=9000
```

### Interview Questions
- What is Helm and why do we use it?
- What is a Helm chart?
- What is values.yaml?
- What is the difference between helm install and helm upgrade?

---

## 10. ArgoCD

### What is it?
A GitOps tool that watches your GitHub repo and automatically deploys to Kubernetes when code changes.

### Definitions
- **GitOps** — GitHub repo is the single source of truth for deployments
- **Sync** — ArgoCD applying latest changes from GitHub to Kubernetes
- **Healthy** — All pods running correctly
- **ArgoCD** — Tool that automates deployment from GitHub to Kubernetes

### How It Works
```
git push
    ↓
ArgoCD detects change in GitHub
    ↓
ArgoCD runs helm upgrade automatically
    ↓
New version deployed to Kubernetes
    ↓
Zero manual work
```

### Interview Questions
- What is ArgoCD?
- What is GitOps?
- What is the difference between GitHub Actions and ArgoCD?
- What does Healthy and Synced mean in ArgoCD?

### Answer
GitHub Actions builds Docker images. ArgoCD deploys them to Kubernetes. Together they form a complete automated pipeline — push code, image built, deployed automatically.

---

## 11. LangChain

### What is it?
A toolkit that connects your code to AI models like Claude, ChatGPT, Llama, Gemini.

### Definitions
- **LLM** — Large Language Model — AI brain like Llama, Claude, GPT
- **LangChain** — Framework to connect and use LLMs in your code
- **API Key** — Password to use an AI model from the cloud

### Simple Example
```python
from langchain_groq import ChatGroq

llm = ChatGroq(model="llama-3.3-70b-versatile", api_key="your-key")
response = llm.invoke("What is an API?")
print(response.content)
```

### Interview Questions
- What is LangChain?
- What is an LLM?
- What is Groq?

---

## 12. LangGraph

### What is it?
Built on LangChain. Connects AI to tools and lets it take actions automatically.

### Definitions
- **Tool** — A function the AI can call (get products, create product)
- **Agent** — AI that decides which tool to call based on your request
- **Pipeline** — Fixed steps that run automatically in order
- **ReAct** — Reasoning + Acting — how agents think and act

### Tool Example
```python
from langchain_core.tools import tool

@tool
def get_products() -> str:
    """Get all products from the API."""
    response = requests.get("http://localhost:8000/products")
    return response.json()
```

### Agent Example
```python
from langgraph.prebuilt import create_react_agent

agent = create_react_agent(llm, tools)
result = agent.invoke({"messages": [("user", "Show me all products")]})
```

### Pipeline Example
```python
workflow = StateGraph(PipelineState)
workflow.add_node("fetch", fetch_products)
workflow.add_node("summarize", summarize_products)
workflow.add_node("recommend", recommend_product)
workflow.add_edge("fetch", "summarize")
workflow.add_edge("summarize", "recommend")
```

### Interview Questions
- What is LangGraph?
- What is the difference between an Agent and a Pipeline?
- What is a Tool in LangGraph?
- What is the difference between LangChain and LangGraph?

---

## 13. CrewAI

### What is it?
A framework to build a team of AI agents. Each agent has a specific role and works together to complete a task.

### Definitions
- **Agent** — One AI worker with a specific role
- **Task** — Job assigned to an agent
- **Crew** — The team of agents working together
- **Role** — What the agent specializes in

### Example
```python
from crewai import Agent, Task, Crew

researcher = Agent(
    role="Product Researcher",
    goal="Find all products from the API",
    backstory="Expert at finding product information"
)

analyst = Agent(
    role="Product Analyst",
    goal="Analyze products and find best value",
    backstory="Expert at analyzing products"
)

crew = Crew(agents=[researcher, analyst], tasks=[...])
result = crew.kickoff()
```

### Interview Questions
- What is CrewAI?
- What is the difference between LangGraph and CrewAI?
- When would you use CrewAI instead of LangGraph?

### Answer
LangGraph = one agent doing everything. CrewAI = team of agents each with one job. Use CrewAI for complex tasks that need multiple roles — researcher, analyst, writer.

---

## Quick Reference — All Commands

### Git
```
git add . → git commit -m "msg" → git push    (save and send)
git pull                                        (get latest)
git status                                      (check changes)
git branch name → git checkout name            (create and switch branch)
git merge name                                  (merge branch)
git stash → git stash pop                      (hide and restore)
```

### Docker
```
docker build -t name .       (build image)
docker run -p 8000:8000 name (run container)
docker push name             (push to Docker Hub)
docker images                (list images)
docker ps                    (list containers)
```

### Kubernetes
```
kubectl apply -f file.yaml          (deploy)
kubectl get pods                    (check pods)
kubectl get services                (check services)
kubectl logs pod-name               (see logs)
minikube start                      (start local k8s)
minikube service name --url         (get URL)
```

### Helm
```
helm install name chart/            (deploy)
helm upgrade name chart/            (update)
helm list                           (see deployments)
helm uninstall name                 (remove)
```

---

## The Full System In One Picture

```
You write code
    ↓
git push → GitHub stores code
    ↓
GitHub Actions → builds Docker images → pushes to Docker Hub
    ↓
ArgoCD detects change → runs Helm upgrade
    ↓
Kubernetes pulls image from Docker Hub → runs container
    ↓
Container runs your API on a port
    ↓
LangGraph / CrewAI agents call your API automatically
    ↓
AI takes actions based on plain English goals
```
