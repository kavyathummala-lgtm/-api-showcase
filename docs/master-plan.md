# Master Plan — Everything To Build and Learn

---

## Status Legend
- ✅ Done
- 🔄 In Progress
- ❌ Not Started

---

## PHASE 1 — REST APIs (✅ Complete)

### What we built
9 APIs in Python, Go, Java across REST, GraphQL, gRPC styles.

### APIs Built
| API | Language | Style | Port | Status |
|-----|----------|-------|------|--------|
| python-rest | Python | REST | 8000 | ✅ |
| go-rest | Go | REST | 8080 | ✅ |
| java-rest | Java | REST | 9000 | ✅ |
| graphql-python | Python | GraphQL | 8001 | ✅ |
| grpc-go | Go | gRPC | 50051 | ✅ |
| + 4 more | | | | ✅ |

### Endpoints built (python-rest example)
```
GET    /products         → get all products
POST   /products         → create a product
PUT    /products/{id}    → update a product
DELETE /products/{id}    → delete a product
GET    /health           → health check
```

### Output when running
```
Uvicorn running on http://127.0.0.1:8000
Application startup complete.
```

---

## PHASE 2 — Docker (✅ Complete)

### What we did
- Wrote Dockerfile for each API
- Built Docker images
- Pushed all 9 images to Docker Hub

### Key commands
```
docker build -t kavyathummala/python-rest .
docker push kavyathummala/python-rest
docker images
docker ps
```

### Output
```
[+] Building 12.3s FINISHED
kavyathummala/python-rest   latest   pushed ✅
kavyathummala/go-rest       latest   pushed ✅
(9 images total on Docker Hub)
```

---

## PHASE 3 — Kubernetes + Minikube (✅ Complete)

### What we did
- Started Minikube locally
- Deployed APIs using kubectl YAML files
- Added health checks (readiness + liveness probes)
- Added search, filter, pagination

### Key commands
```
minikube start
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl get pods
kubectl get services
kubectl logs pod-name
```

### Output
```
NAME                           READY   STATUS    AGE
python-rest-7d9f8c6b4-xk2pq   1/1     Running   2m

NAME          TYPE       PORT(S)
python-rest   NodePort   8000:30000/TCP
```

---

## PHASE 4 — GitHub Actions CI/CD (✅ Complete)

### What we did
- Created .github/workflows/ci.yml
- Every git push automatically builds all 9 Docker images and pushes to Docker Hub

### Flow
```
git push → GitHub Actions triggered → builds 9 images → pushes to Docker Hub
```

### Output (GitHub Actions tab)
```
✅ Build and Push Docker Images — 4m 32s
  ✅ Build python-rest
  ✅ Push python-rest
  (all 9 images)
```

---

## PHASE 5 — Helm Charts (✅ Complete)

### What we did
- Created api-chart with templates
- One chart deploys any of the 9 APIs by changing values.yaml

### Key commands
```
helm install python-rest api-chart/ --set image=kavyathummala/python-rest --set port=8000
helm upgrade python-rest api-chart/
helm list
helm uninstall python-rest
```

### Output
```
NAME: python-rest
STATUS: deployed
REVISION: 1
```

---

## PHASE 6 — ArgoCD (✅ Complete)

### What we did
- Installed ArgoCD on Kubernetes
- Connected to GitHub repo
- ArgoCD watches GitHub and auto-deploys using Helm when code changes

### Flow
```
git push → ArgoCD detects change → runs helm upgrade → new version deployed
```

### Output (ArgoCD dashboard)
```
Status:  ✅ Healthy
Sync:    ✅ Synced
```

---

## PHASE 7 — LangChain + LangGraph (✅ Complete)

### What we built
| File | What it does | Status |
|------|-------------|--------|
| practice1_connect.py | Connect to Llama AI | ✅ |
| practice2_tools.py | Give AI tools to call API | ✅ |
| practice3_agent.py | AI decides which tool to call | ✅ |
| practice4_pipeline.py | Fixed steps: fetch→summarize→recommend | ✅ |

### Key concepts
- **Tool** = Python function AI can call
- **Agent** = AI decides what to do
- **Pipeline** = you decide fixed steps

### Output (agent)
```
You asked: Show me all products
Agent called: get_products tool
Agent answer: Here are all products: Gaming Keyboard $89.99, Smart Watch $199.99
```

### Output (pipeline)
```
Step 1: Fetching products... Got 2 products
Step 2: AI summarizing... Done
Step 3: AI recommending... Done
Summary: Two electronics products available...
Recommendation: Gaming Keyboard is best value at $89.99
```

---

## PHASE 8 — CrewAI (✅ Complete)

### What we built
| File | What it does | Status |
|------|-------------|--------|
| practice5_crew.py | 2 agents — Researcher + Analyst | ✅ |

### Key concepts
- **Agent** = AI + role + goal + personality
- **Task** = job given to an agent
- **Crew** = team of agents working together

### Output
```
🤖 Product Researcher → organized clean product list
🤖 Product Analyst → cheapest, most expensive, recommendation
✅ Crew Execution Completed
```

---

## PHASE 9 — Wrap AI Frameworks as APIs (✅ Complete)

### What we will do
Take the AI framework scripts and wrap them in FastAPI so they can be called like a normal API.

### Plan
```
LangGraph agent code
    ↓
FastAPI endpoint: POST /agent/ask
    ↓
Docker image
    ↓
Push to Docker Hub
    ↓
Deploy to Kubernetes
```

### Files to create
- `ai-frameworks/langgraph/api.py` — FastAPI wrapper for agent
- `ai-frameworks/crewai/api.py` — FastAPI wrapper for crew

### Expected output
```
POST /agent/ask
body: { "question": "show me all products" }

response: { "answer": "Here are your products: ..." }
```

---

## PHASE 10 — Google ADK (❌ Not Started)

### What is it
Google's framework for building AI agents. Similar to LangGraph but made by Google.

### What we will build
- Agent example — AI calls your product API
- Tool example — functions the AI can use
- Pipeline example — fixed steps

### Files to create
- `ai-frameworks/google-adk/agent.py`
- `ai-frameworks/google-adk/api.py`

---

## PHASE 11 — AWS Bedrock (❌ Not Started)

### What is it
Run AI models (Claude, Llama, Titan) on Amazon cloud instead of Groq.

### What we will build
- Agent example using AWS Bedrock
- Deploy on AWS (EKS or Lambda)

### Files to create
- `ai-frameworks/aws-bedrock/agent.py`
- `ai-frameworks/aws-bedrock/api.py`

---

## PHASE 12 — LlamaIndex (❌ Not Started)

### What is it
Connect AI to your own documents and data. AI can search through PDFs, text files, databases and answer questions about them.

### What we will build
- Pipeline that reads product data and answers questions about it

### Files to create
- `ai-frameworks/llamaindex/pipeline.py`

---

## PHASE 13 — MCP Server (❌ Not Started)

### What is it
Model Context Protocol — a standard way to give tools to AI. Instead of writing tools differently for each framework, MCP is one standard all frameworks can use.

### What we will build
- MCP server that exposes product API as tools

---

## PHASE 14 — Jenkins + Git Runner (❌ Not Started)

### What is it
- **Jenkins** = CI/CD tool like GitHub Actions but self-hosted on your own server
- **Git Runner** = self-hosted machine that runs the CI/CD jobs

### What we will build
- Jenkins pipeline that builds and deploys Docker images
- Self-hosted runner connected to GitHub

---

## PHASE 15 — Deploy Everything to Kubernetes (❌ Not Started)

### What we will do
Deploy all AI framework APIs (wrapped in FastAPI) to Kubernetes using Helm and ArgoCD.

### Expected final state
```
kubectl get pods

python-rest          Running ✅
go-rest              Running ✅
java-rest            Running ✅
langgraph-agent-api  Running ✅
crewai-api           Running ✅
google-adk-api       Running ✅
```

---

## PHASE 16 — Deploy on AWS and Google Cloud (❌ Not Started)

### AWS
- EKS (Elastic Kubernetes Service) — Kubernetes on AWS
- Or Lambda — serverless functions on AWS

### Google Cloud
- GKE (Google Kubernetes Engine) — Kubernetes on Google Cloud

---

## Quick Interview Reference

### The Full Pipeline (everything connected)
```
You write code
    ↓
git push → GitHub stores code
    ↓
GitHub Actions → builds Docker images → pushes to Docker Hub
    ↓
ArgoCD detects change → runs Helm upgrade
    ↓
Kubernetes pulls image → runs container
    ↓
Container runs your API on a port
    ↓
LangGraph / CrewAI agents call your API
    ↓
AI takes actions based on plain English goals
    ↓
FastAPI wraps AI as an API endpoint
    ↓
Deployed on Kubernetes / AWS / Google Cloud
```

### Key Definitions (one line each)
- **Docker** — packages code into a container that runs anywhere
- **Kubernetes** — manages, scales, restarts containers automatically
- **Helm** — package manager for Kubernetes
- **ArgoCD** — watches GitHub and auto-deploys to Kubernetes
- **GitHub Actions** — robot that builds and pushes images on git push
- **LangChain** — connects Python code to AI models
- **LangGraph** — builds agents and pipelines on top of LangChain
- **CrewAI** — builds teams of AI agents each with one role
- **Tool** — Python function the AI can call
- **Agent** — AI decides which tool to call
- **Pipeline** — fixed steps you define, always runs in order
- **Crew** — team of agents working together
