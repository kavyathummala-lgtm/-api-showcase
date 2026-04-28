# Learning Roadmap — What, Why, and How to Learn

---

## Why This Project Exists

You are helping build an AI and LLM platform for your family.
That platform needs APIs, databases, authentication, cloud deployment, and AI integration.
Every topic in this document is a real skill used in that platform.
You are not learning for exams — you are learning to build something real.

---

## Topics You Already Completed

---

### Python, Go, Java

**What:** Three programming languages. Each one can build APIs, handle logic, and talk to databases.

**Why you learned it:**
- Python — simplest syntax, fastest to write, most popular for AI and data work
- Go — extremely fast, low memory, used at Google, Uber, Netflix for high-traffic services
- Java — standard in large enterprises and banks, most job listings require it

**How you learned it:** Built 3 API servers in each language doing the same job — managing a product list.

---

### REST, GraphQL, gRPC

**What:** Three ways an API can receive and respond to requests.

**Why you learned it:**
- REST — used in every web app, mobile app, and public API in the world
- GraphQL — used by Facebook, GitHub, Shopify — avoids sending unnecessary data
- gRPC — used between internal services at Google, Netflix — binary, very fast

**How you learned it:** Built one of each in all 3 languages — 9 servers total.

---

### Docker

**What:** Packages your code and all its dependencies into a container — a portable box that runs the same everywhere.

**Why you learned it:**
- Solves "works on my machine" problem
- Required for Kubernetes deployment
- Every company uses Docker to ship software

**How you learned it:** Wrote a Dockerfile for each of the 9 APIs and built images.

---

### Kubernetes and Minikube

**What:** Kubernetes manages and runs containers automatically. Minikube runs Kubernetes on your laptop.

**Why you learned it:**
- Restarts crashed containers automatically
- Scales up when traffic increases
- Standard deployment platform at every tech company

**How you learned it:** Wrote k8s YAML files for all 9 APIs, deployed to Minikube, tested health checks.

---

### Health Checks

**What:** A `/health` endpoint that tells Kubernetes if the server is alive and ready.

**Why you learned it:**
- Without health checks, Kubernetes sends traffic to servers that are still starting up or have crashed
- Every production API must have this

**How you learned it:** Added `/health` to all REST and GraphQL APIs, TCP socket check for gRPC.

---

### Search, Filter, and Pagination

**What:** Query parameters that let clients ask for exactly what they need — filter by name, category, price, and get results in pages.

**Why you learned it:**
- Without filtering, every request returns all data — slow and wasteful
- Without pagination, returning millions of records crashes the app
- Used by every e-commerce site, search engine, and data API

**How you learned it:** Added `?name=`, `?category=`, `?min_price=`, `?page=`, `?limit=` to all 9 APIs.

---

### GitHub and GitHub Actions (CI/CD)

**What:** GitHub stores your code. GitHub Actions automatically builds Docker images and pushes them to Docker Hub every time you push code.

**Why you learned it:**
- No manual deployment — push code and it deploys itself
- Every tech company uses CI/CD
- Standard skill required in every software job

**How you learned it:** Created a GitHub repo, pushed all 9 APIs, wrote a workflow that builds and pushes all 9 Docker images automatically.

---

## Topics Coming Next

---

### 1. Helm Charts

**What:** Right now you have 9 separate Kubernetes YAML files. Helm bundles them into one reusable package with settings you can change easily — like a recipe book instead of 9 loose papers.

**Why learn it:**
- Raw YAML files are hard to manage when you have dozens of services
- Helm is the standard way to package and deploy Kubernetes applications
- ArgoCD (next topic) uses Helm to deploy

**How to learn it easily:**
1. Install Helm on your laptop
2. Convert one of your existing k8s YAML files into a Helm chart
3. Deploy it with `helm install` instead of `kubectl apply`
4. See how one chart can deploy to dev, staging, and production with different settings

---

### 2. PostgreSQL

**What:** A real database that stores data on disk permanently. Right now your APIs store data in RAM — it disappears when the server restarts. PostgreSQL keeps data forever.

**Why learn it:**
- In-memory storage is only for learning — no real app uses it
- Every application needs a database
- PostgreSQL is free, open source, and used at Instagram, Spotify, and most startups

**How to learn it easily:**
1. Run PostgreSQL in Docker — one command, no installation needed
2. Replace the in-memory dictionary in Python REST with a PostgreSQL table
3. Learn 4 SQL commands: `INSERT`, `SELECT`, `UPDATE`, `DELETE` — they match CRUD exactly
4. Add PostgreSQL to Go and Java REST next

---

### 3. JWT Authentication

**What:** Right now anyone can call your APIs with no login. JWT adds a lock. User logs in with username and password, gets a token (a long string). Every future request must send that token. Server checks it — valid token means access allowed, no token means blocked.

**Why learn it:**
- You cannot ship a real API without authentication
- JWT is stateless — no session storage needed, works perfectly with Kubernetes
- Used by every app that has user accounts

**How to learn it easily:**
1. Add a `/login` endpoint that returns a JWT token
2. Add a middleware that checks the token on every request
3. Test with Postman — first login to get token, then use token to call protected endpoints
4. Do this in Python REST first, then Go and Java

---

### 4. Unit Tests

**What:** Code that tests your own code automatically. You write a function, then write a test that calls that function and checks the result is correct. Run all tests with one command.

**Why learn it:**
- Every company requires tests before merging code
- Without tests, you don't know you broke something until a user reports it
- Tests let you change code confidently

**How to learn it easily:**
- Python — use `pytest`. Write a test file, run `pytest`, see pass/fail
- Go — built-in `testing` package, run `go test ./...`
- Java — use JUnit, already included in Spring Boot
- Start simple: test that `GET /products` returns 200, test that creating a product works

---

### 5. ArgoCD

**What:** A tool that watches your GitHub repo. When you push code, ArgoCD detects the change and automatically deploys the new version to Kubernetes. No manual `kubectl apply` needed.

**Why learn it:**
- This is GitOps — your Git repo is the single source of truth for what runs in production
- Standard deployment method at tech companies
- Combines with Helm: ArgoCD deploys your Helm charts automatically

**How to learn it easily:**
1. Install ArgoCD inside Minikube
2. Point it at your GitHub repo
3. Push a change — watch ArgoCD detect it and redeploy automatically
4. Open the ArgoCD dashboard — visual map of all your deployments

---

### 6. Jenkins

**What:** A self-hosted CI/CD server. Does the same job as GitHub Actions — build, test, deploy — but runs on your own machine or server instead of GitHub's cloud.

**Why learn it:**
- Older companies (banks, enterprises) use Jenkins and will not switch
- More control and customization than GitHub Actions
- Knowing both makes you employable at any company

**How to learn it easily:**
1. Run Jenkins in Docker — one command
2. Create a Pipeline job that builds one Docker image
3. Compare it side by side with your GitHub Actions workflow — they do the same thing differently
4. You already know the concepts — Jenkins just has a different UI

---

### 7. Terraform

**What:** Infrastructure as code. Instead of clicking buttons in AWS to create servers, you write a file describing what you want — servers, databases, networks — and Terraform creates it all automatically.

**Why learn it:**
- Right now Minikube runs on your laptop. Terraform moves you to real cloud
- Infrastructure created by clicking buttons cannot be repeated or tracked. Terraform files can
- Standard skill for DevOps and cloud engineering roles

**How to learn it easily:**
1. Create a free AWS account
2. Write a simple Terraform file that creates one server
3. Run `terraform apply` — watch the server appear in AWS
4. Run `terraform destroy` — server deleted, no cost
5. Progress to creating a Kubernetes cluster on AWS (EKS)

---

### 8. Angular

**What:** A frontend framework for building websites. Right now your APIs are invisible — only developers can call them with tools like Postman. Angular builds the actual screen users see — buttons, forms, tables, pages.

**Why learn it:**
- Your APIs need a frontend so real users can interact with them
- Angular connects to your REST API and displays the data as a website
- Full stack = backend APIs + frontend — much more valuable portfolio

**How to learn it easily:**
1. Install Angular CLI — one command
2. Create a new Angular project — one command generates everything
3. Build a product list page that calls your Python REST API
4. Add a form to create a new product
5. Add delete and edit buttons
6. You now have a complete full-stack application

---

### 9. Claude API

**What:** Instead of going to claude.ai and typing, you call Claude directly from your code. Your Python/Go/Java app sends a message to Claude and gets a response back — fully automated inside your application.

**Why learn it:**
- This is exactly what your family's platform does — call an LLM from code
- You can add AI features to any application — summaries, recommendations, answers
- Anthropic's API is simple — a few lines of Python is enough to get started

**How to learn it easily:**
1. Get an Anthropic API key
2. Install the Python SDK: `pip install anthropic`
3. Write 10 lines of Python that send a message and print the response
4. Add a `/ask` endpoint to your Python REST API that accepts a question and returns Claude's answer
5. Call it from Postman — your API now has AI built in

---

### 10. Simple AI Agent

**What:** An agent goes further than answering questions — it takes actions. You give it a goal, it decides what steps to take, calls your APIs, processes results, and completes the task without you directing each step.

**Example:** Tell the agent "find the most expensive product and delete it."
- Agent calls `GET /products` to list all products
- Agent finds the one with the highest price
- Agent calls `DELETE /products/id` to remove it
- Agent reports what it did

**Why learn it:**
- This is the core of what your family is building — an AI system that takes actions, not just answers questions
- Agents are the next wave of AI applications — every company is building them now
- You already have the APIs — the agent just calls them intelligently

**How to learn it easily:**
1. Learn Claude's tool use feature — you define functions, Claude decides when to call them
2. Give the agent two tools: list products and delete product
3. Ask it to find and delete the cheapest item — watch it figure out the steps
4. Expand: add create and update as tools — now the agent can manage your entire product list

---

### 11. Playwright

**What:** A robot that controls a real browser automatically. It opens your Angular website, clicks buttons, fills forms, checks that the right data appears, and reports if anything is broken. Runs thousands of checks in seconds.

**Why learn it:**
- Manual testing is slow and misses edge cases
- Every company runs automated browser tests before shipping
- Playwright works with any website and supports Python, JavaScript, and Java

**How to learn it easily:**
1. Install Playwright — one command
2. Write a test that opens your Angular product list page and checks it loaded
3. Write a test that fills the create product form and submits it
4. Write a test that checks the new product appears in the list
5. Run all tests with one command — instant feedback on every change

---

## Complete Learning Order

| # | Topic | Builds On |
|---|-------|-----------|
| Done | Python, Go, Java | — |
| Done | REST, GraphQL, gRPC | Languages |
| Done | Docker | Code |
| Done | Kubernetes + Minikube | Docker |
| Done | Health Checks | Kubernetes |
| Done | Search, Filter, Pagination | APIs |
| Done | GitHub Actions | Docker + GitHub |
| 1 | Helm Charts | Kubernetes YAML files |
| 2 | PostgreSQL | APIs + SQL basics |
| 3 | JWT Authentication | APIs + PostgreSQL |
| 4 | Unit Tests | All APIs |
| 5 | ArgoCD | Helm + GitHub |
| 6 | Jenkins | GitHub Actions concepts |
| 7 | Terraform | Docker + Kubernetes |
| 8 | Angular | REST APIs |
| 9 | Claude API | Python + APIs |
| 10 | AI Agent | Claude API + your APIs |
| 11 | Playwright | Angular frontend |

---

## The Final Picture

```
User opens Angular website
        ↓
Angular calls your REST / GraphQL / gRPC API
        ↓
JWT checks login token — blocks if not authenticated
        ↓
API reads and writes data to PostgreSQL database
        ↓
API optionally calls Claude API for AI responses
        ↓
AI Agent can call the API autonomously to take actions
        ↓
APIs run in Docker containers
        ↓
Kubernetes manages containers (config packaged with Helm)
        ↓
ArgoCD auto-deploys when you push to GitHub
        ↓
GitHub Actions builds Docker images on every push
        ↓
Jenkins does the same on your own server
        ↓
Terraform created the cloud servers everything runs on
        ↓
Unit Tests + Playwright verify everything works automatically
```

Everything connects. Every topic you learn adds one more piece to this picture.
