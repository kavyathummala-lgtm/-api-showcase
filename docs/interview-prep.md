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

### Answers
**What is REST and why is it popular?**
REST uses different URLs for different actions — GET to read, POST to create, PUT to update, DELETE to remove. It is popular because it is simple, uses JSON which everyone understands, and works with any programming language.

**What is the difference between REST and GraphQL?**
REST has many URLs — one for products, one for users, one for orders. GraphQL has ONE URL and you ask for exactly what you need in one request. REST might give you 20 fields but you only needed 3. GraphQL gives you exactly 3.

**What is gRPC used for?**
gRPC is used when speed is critical — like communication between internal services where milliseconds matter. Instead of sending readable JSON text, it sends compressed binary data which is much faster. Used by Google, Netflix internally.

**What is a proto file?**
A proto file defines the structure of your API — what functions exist, what data goes in, what data comes back. It is like a contract both sides agree to. The `.proto` file generates code automatically for any language.

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

### Answers
**What is a health check?**
A health check is a special URL `/health` in your API that returns "ok" when the service is running fine. Kubernetes calls this URL automatically to check if the container is alive and ready.

**What is the difference between readiness and liveness probe?**
Readiness probe asks "is this pod ready to receive traffic?" — if no, Kubernetes stops sending requests to it but does not kill it. Liveness probe asks "is this pod still alive?" — if no, Kubernetes kills it and starts a new one. Readiness = ready for users. Liveness = still breathing.

**Why do we need health checks in Kubernetes?**
Without health checks, Kubernetes has no way to know if your app crashed inside the container. The container could be running but your code is broken. Health checks let Kubernetes detect this and automatically restart the bad pod.

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

### Answers
**What is pagination and why do we need it?**
Pagination means returning results in small pages instead of all at once. If your database has 10,000 products and someone calls GET /products, sending all 10,000 in one response would be slow and crash the browser. Pagination returns 10 or 20 at a time — page 1, page 2, etc. — so it stays fast.

**What are query parameters?**
Query parameters are extra instructions added to a URL after a `?`. Example: `/products?category=Electronics&limit=10`. The `category=Electronics` tells the API to filter by Electronics. The `limit=10` tells it to return only 10 results. Multiple parameters are joined with `&`.

**How do you implement search in a REST API?**
You read the query parameter from the URL and filter the data. In FastAPI: `@app.get("/products")` with `name: str = None` as a parameter. If name is provided, filter the list to only items where the name contains that word. Return the filtered list.

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

### Answers
**What is Git and why do we use it?**
Git tracks every change you make to your code. If you break something, you can go back to an earlier version. If two developers edit the same file, Git helps merge both changes. Without Git, teams would constantly overwrite each other's work.

**What is the difference between git push and git pull?**
`git push` sends your changes from your laptop UP to GitHub. `git pull` brings the latest changes FROM GitHub down to your laptop. Push = upload. Pull = download.

**What is a branch and why do we use it?**
A branch is a separate copy of the code where you can work on a new feature without touching the main code. When the feature is ready and tested, you merge it back into main. This way broken code never reaches production. Example: `git branch add-login` creates a branch for building login.

**What is git stash?**
Git stash temporarily hides your unfinished changes so you can switch to another task. Your changes are saved but not committed. `git stash` hides them. `git stash pop` brings them back. Useful when someone asks you to fix an urgent bug while you are in the middle of building a feature.

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

### Answers
**What is CI/CD?**
CI/CD is the automated process of building, testing, and deploying code every time a developer pushes. Without it, developers manually build and deploy — slow, error-prone. With CI/CD, the moment you push code, a robot builds it, runs tests, and deploys it automatically.

**What is GitHub Actions?**
GitHub Actions is the built-in CI/CD robot inside GitHub. You write instructions in a `.yml` file inside `.github/workflows/`. Every time you push, GitHub Actions reads that file and runs the steps — build Docker image, push to Docker Hub, deploy, etc.

**What is the difference between CI and CD?**
CI (Continuous Integration) = automatically build and test the code when pushed. CD (Continuous Deployment) = automatically deploy the tested code to production. CI is the build and test part. CD is the release part. In our project: CI builds the 9 Docker images, CD pushes them to Docker Hub.

**What is a runner?**
A runner is the machine that actually executes the GitHub Actions steps. GitHub provides free runners (`ubuntu-latest`) in the cloud. When you push, GitHub starts a fresh Linux machine, runs all your steps, and shuts it down. You can also set up your own runner on your own server (self-hosted runner).

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

### Answers
**What is Helm and why do we use it?**
Helm is a package manager for Kubernetes — like how npm installs packages for Node.js, Helm installs applications into Kubernetes. Without Helm, you need a separate deployment.yaml and service.yaml for every API. With Helm, one chart deploys any of the 9 APIs just by changing the values.

**What is a Helm chart?**
A Helm chart is a folder with template YAML files that have variables like `{{ .Values.port }}` instead of hardcoded values. The chart is a reusable blueprint. You fill in the values and it generates the final Kubernetes YAML for you.

**What is values.yaml?**
`values.yaml` is the settings file for a Helm chart. It contains the actual values — which image to use, which port, how many replicas, which health check path. When you run `helm install`, Helm reads values.yaml and fills in all the `{{ .Values.X }}` variables in the templates.

**What is the difference between helm install and helm upgrade?**
`helm install` creates a brand new deployment — use this the first time. `helm upgrade` updates an existing deployment — use this when you want to change the image version or any setting. If you run `helm install` when it already exists, you get an error. `helm upgrade` safely applies the change.

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

### Answers
**What is LangChain?**
LangChain is a toolkit that makes it easy to connect your Python code to AI models like Claude, ChatGPT, Llama, and Gemini. Instead of writing complex API calls yourself, LangChain gives you simple functions — just call `llm.invoke("your question")` and get the AI response back.

**What is an LLM?**
LLM stands for Large Language Model. It is the AI brain — like Llama by Meta, GPT-4 by OpenAI, Claude by Anthropic, Gemini by Google. These models are trained on huge amounts of text and can understand and generate human language, answer questions, summarize, and write code.

**What is Groq?**
Groq is a cloud platform that runs open source AI models like Llama for free — no credit card needed. It is extremely fast because they built custom chips for AI. In our project we use Groq to run the Llama 3.3 70B model without paying anything.

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

### Answers
**What is LangGraph?**
LangGraph is built on top of LangChain and lets you build AI agents and automated pipelines. LangChain just connects you to an AI model. LangGraph goes further — it lets the AI use tools (call APIs, run functions) and make decisions about what to do next.

**What is the difference between an Agent and a Pipeline?**
An Agent is flexible — you give it a goal in plain English and it figures out which tools to call and in what order. A Pipeline is fixed — you define exactly which steps run and in what order, every time. Agent = AI decides. Pipeline = you decide.

**What is a Tool in LangGraph?**
A Tool is a Python function that the AI can choose to call. You mark it with `@tool` and give it a clear description. The AI reads the description, decides if it needs that function, and calls it with the right parameters. In our project: `get_products`, `create_product`, and `delete_product` are tools the AI calls when needed.

**What is the difference between LangChain and LangGraph?**
LangChain = gives you the connection to the AI model. LangGraph = adds the ability to build agents and multi-step workflows on top of LangChain. You need LangChain to use LangGraph. LangGraph is what makes the AI actually DO things, not just talk.

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
