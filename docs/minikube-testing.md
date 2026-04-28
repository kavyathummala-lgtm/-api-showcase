# How to Deploy and Test All 9 APIs on Minikube — Step by Step

Minikube lets you run Kubernetes on your own laptop. This guide deploys all 9 APIs as containers and tests each one.

---

## Before You Start — Check These Are Installed

Open a terminal and run each command:

```
docker --version        → Docker version 27.x or higher
minikube version        → minikube version v1.x
kubectl version --client → Client Version: v1.x
```

Also make sure **Docker Desktop is open and running** (whale icon in taskbar is not spinning).

---

## STEP 1 — Start Minikube

```powershell
minikube start --driver=docker
```

Wait until you see:
```
Done! kubectl is now configured to use "minikube" cluster
```

If it says "Restarting existing docker container" — that is fine too, it means Minikube was already set up.

---

## STEP 2 — Point Docker to Minikube

This is very important. Run this so Docker builds images **inside Minikube** instead of on your regular computer:

```powershell
& minikube -p minikube docker-env --shell powershell | Invoke-Expression
```

You will see some lines starting with `$Env:DOCKER_...` — that means it worked.

> **Keep this terminal open. Run ALL docker build commands in this same terminal.**

---

## STEP 3 — Build All 9 Docker Images

Run these one at a time. Wait for each to finish before running the next.
Each one ends with: `Successfully tagged xxx:latest`

```powershell
docker build -t python-rest:latest C:\Users\kavya\api-showcase\python\rest
```
```powershell
docker build -t python-graphql:latest C:\Users\kavya\api-showcase\python\graphql
```
```powershell
docker build -t python-grpc:latest C:\Users\kavya\api-showcase\python\grpc
```
```powershell
docker build -t go-rest:latest C:\Users\kavya\api-showcase\go\rest
```
```powershell
docker build -t go-graphql:latest C:\Users\kavya\api-showcase\go\graphql
```
```powershell
docker build -t go-grpc:latest C:\Users\kavya\api-showcase\go\grpc
```
```powershell
docker build -t java-rest:latest C:\Users\kavya\api-showcase\java\rest
```
```powershell
docker build -t java-graphql:latest C:\Users\kavya\api-showcase\java\graphql
```
```powershell
docker build -t java-grpc:latest C:\Users\kavya\api-showcase\java\grpc
```

> Java builds take longer (3–5 minutes each) — just wait.

---

## STEP 4 — Deploy All 9 Services to Kubernetes

```powershell
kubectl apply -f C:\Users\kavya\api-showcase\k8s\rest\ -f C:\Users\kavya\api-showcase\k8s\graphql\ -f C:\Users\kavya\api-showcase\k8s\grpc\
```

You will see 18 lines like:
```
deployment.apps/python-rest created
service/python-rest created
...
```

---

## STEP 5 — Check All Pods Are Running

```powershell
kubectl get pods
```

Wait about 30 seconds, then run it. You should see all 9 pods with STATUS = `Running`:

```
NAME                              READY   STATUS    RESTARTS   AGE
go-graphql-xxx                    1/1     Running   0          30s
go-grpc-xxx                       1/1     Running   0          30s
go-rest-xxx                       1/1     Running   0          30s
java-graphql-xxx                  1/1     Running   0          30s
java-grpc-xxx                     1/1     Running   0          30s
java-rest-xxx                     1/1     Running   0          30s
python-graphql-xxx                1/1     Running   0          30s
python-grpc-xxx                   1/1     Running   0          30s
python-rest-xxx                   1/1     Running   0          30s
```

If any say `ContainerCreating` — wait 30 more seconds and run again.

---

## STEP 6 — Get URLs and Test Each API

> **Important for Windows:** The Minikube IP (192.168.49.2) does not work directly in your browser on Windows.
> You must use `minikube service <name> --url` which gives you a working localhost URL.
> Each command keeps that terminal open — **do not close it**.

---

### Python REST (has browser Swagger UI)

Open a **new terminal** and run:
```powershell
minikube service python-rest --url
```

You will get a URL like `http://127.0.0.1:XXXXX`

Open your browser and go to that URL + `/docs`:
```
http://127.0.0.1:XXXXX/docs
```

You will see the Swagger UI page with all 5 operations.

---

### Python GraphQL (has browser playground)

Open a **new terminal** and run:
```powershell
minikube service python-graphql --url
```

Open your browser and go to that URL + `/graphql`:
```
http://127.0.0.1:XXXXX/graphql
```

You will see the GraphQL playground where you can type queries.

---

### Go REST (browser shows raw JSON)

Open a **new terminal** and run:
```powershell
minikube service go-rest --url
```

Open your browser and go to that URL + `/products`:
```
http://127.0.0.1:XXXXX/products
```

You will see `[]` (empty list — no products yet). That means it is working.

---

### Go GraphQL (no browser UI — use PowerShell)

Open a **new terminal** and run:
```powershell
minikube service go-graphql --url
```

Copy the URL it gives (e.g. `http://127.0.0.1:15549`). Keep that terminal open.

Open another **new terminal** and run (replace the URL with yours):
```powershell
$body = @{ query = "query { products { id name } }" } | ConvertTo-Json
$result = Invoke-RestMethod -Uri "http://127.0.0.1:15549/graphql" -Method POST -ContentType "application/json" -Body $body
$result.data.products
```

Empty output = working (no products yet).

---

### Java REST (browser shows raw JSON)

Open a **new terminal** and run:
```powershell
minikube service java-rest --url
```

Open your browser and go to that URL + `/products`:
```
http://127.0.0.1:XXXXX/products
```

You will see `[]`.

---

### Java GraphQL (no browser UI — use PowerShell)

Open a **new terminal** and run:
```powershell
minikube service java-graphql --url
```

Copy the URL. Keep that terminal open. In another **new terminal** run (replace the URL):
```powershell
$r = Invoke-RestMethod -Uri "http://127.0.0.1:XXXXX/graphql" -Method POST -ContentType "application/json" -Body '{"query":"query { products { id name } }"}'
$r.data.products
```

Empty output = working.

---

### Python gRPC (use grpcurl)

Open a **new terminal** and run:
```powershell
minikube service python-grpc --url
```

Copy the URL (e.g. `http://127.0.0.1:44565`). Keep that terminal open.

In another **new terminal** run (replace the port number with yours):
```powershell
grpcurl -plaintext -import-path C:\Users\kavya\api-showcase\python\grpc -proto product.proto 127.0.0.1:44565 product.ProductService/GetProducts
```

You will see `{}` = working (no products yet).

---

### Go gRPC (use grpcurl)

Open a **new terminal** and run:
```powershell
minikube service go-grpc --url
```

Copy the URL. Keep that terminal open. In another **new terminal** run (replace port):
```powershell
grpcurl -plaintext -import-path C:\Users\kavya\api-showcase\go\grpc\proto -proto product.proto 127.0.0.1:20170 product.ProductService/GetProducts
```

You will see `{}` = working.

---

### Java gRPC (use grpcurl)

Open a **new terminal** and run:
```powershell
minikube service java-grpc --url
```

Copy the URL. Keep that terminal open. In another **new terminal** run (replace port):
```powershell
grpcurl -plaintext -import-path C:\Users\kavya\api-showcase\java\grpc\src\main\proto -proto product.proto 127.0.0.1:2185 product.ProductService/GetProducts
```

You will see `{}` = working.

---

## All 9 APIs Confirmed Working ✓

| API | How to Test | Expected Output |
|-----|------------|-----------------|
| Python REST | Browser `/docs` | Swagger UI page |
| Python GraphQL | Browser `/graphql` | GraphQL playground |
| Python gRPC | grpcurl | `{}` |
| Go REST | Browser `/products` | `[]` |
| Go GraphQL | PowerShell | empty (no products) |
| Go gRPC | grpcurl | `{}` |
| Java REST | Browser `/products` | `[]` |
| Java GraphQL | PowerShell | empty (no products) |
| Java gRPC | grpcurl | `{}` |

---

## How to Stop Everything

```powershell
kubectl delete -f C:\Users\kavya\api-showcase\k8s\rest\ -f C:\Users\kavya\api-showcase\k8s\graphql\ -f C:\Users\kavya\api-showcase\k8s\grpc\
```

Then stop Minikube:
```powershell
minikube stop
```

---

## Common Problems

| Problem | Fix |
|---------|-----|
| `minikube start` fails | Open Docker Desktop first, wait for it to fully start |
| Pod shows `ErrImageNeverPull` | You forgot Step 2 — run the `docker-env` command and rebuild images |
| Pod shows `CrashLoopBackOff` | Run `kubectl logs <pod-name>` to see the error |
| Browser says "This site can't be reached" | Use `minikube service <name> --url` — the 192.168.49.2 IP does not work on Windows |
| grpcurl says "server does not support reflection" | Use `-import-path` and `-proto` flags as shown above |
| Port number changes every time | Normal — `minikube service --url` gives a new port each session |
