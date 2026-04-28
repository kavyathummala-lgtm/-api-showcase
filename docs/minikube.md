# Running on Minikube

## Prerequisites
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) installed
- [kubectl](https://kubernetes.io/docs/tasks/tools/) installed
- Docker Desktop running (used as the Minikube driver)

---

## Step 1 — Start Minikube

```bash
minikube start --driver=docker
```

---

## Step 2 — Point Your Shell at Minikube's Docker Daemon

Images must be built **inside** Minikube, not your local Docker.

```bash
# Linux / Mac
eval $(minikube docker-env)

# Windows PowerShell
& minikube -p minikube docker-env --shell powershell | Invoke-Expression
```

> Run all `docker build` commands below **in this same terminal session**.

---

## Step 3 — Build All Images Inside Minikube

Run from the `api-showcase` root:

```bash
docker build -t python-rest:latest  ./python/rest
docker build -t python-graphql:latest ./python/graphql
docker build -t python-grpc:latest  ./python/grpc
docker build -t go-rest:latest      ./go/rest
docker build -t go-graphql:latest   ./go/graphql
docker build -t go-grpc:latest      ./go/grpc
docker build -t java-rest:latest    ./java/rest
docker build -t java-graphql:latest ./java/graphql
docker build -t java-grpc:latest    ./java/grpc
```

---

## Step 4 — Deploy All Services

Apply manifests from the central `k8s/` folder:

```bash
# REST services
kubectl apply -f k8s/rest/

# GraphQL services
kubectl apply -f k8s/graphql/

# gRPC services
kubectl apply -f k8s/grpc/
```

Or deploy all at once:
```bash
kubectl apply -f k8s/rest/ -f k8s/graphql/ -f k8s/grpc/
```

---

## Step 5 — Check Status

```bash
kubectl get pods        # all should be Running
kubectl get services    # confirm NodePorts
```

---

## Step 6 — Access Services

Get Minikube's IP:
```bash
minikube ip
# e.g. 192.168.49.2
```

| Service        | URL / Address                          | NodePort |
|----------------|----------------------------------------|----------|
| python-rest    | http://\<minikube-ip\>:30000/products  | 30000    |
| python-graphql | http://\<minikube-ip\>:30001/graphql   | 30001    |
| python-grpc    | \<minikube-ip\>:30002                  | 30002    |
| go-rest        | http://\<minikube-ip\>:30003/products  | 30003    |
| go-graphql     | http://\<minikube-ip\>:30004/graphql   | 30004    |
| go-grpc        | \<minikube-ip\>:30005                  | 30005    |
| java-rest      | http://\<minikube-ip\>:30006/products  | 30006    |
| java-graphql   | http://\<minikube-ip\>:30007/graphql   | 30007    |
| java-grpc      | \<minikube-ip\>:30008                  | 30008    |

Or let Minikube open the browser for you:
```bash
minikube service python-rest
minikube service go-rest
minikube service java-rest
```

---

## Step 7 — Tear Down

```bash
# Delete specific service
kubectl delete -f k8s/rest/python-rest.yaml

# Delete all from k8s/
kubectl delete -f k8s/rest/ -f k8s/graphql/ -f k8s/grpc/

# Stop Minikube
minikube stop

# Delete the cluster entirely
minikube delete
```

---

## Troubleshooting

| Problem | Solution |
|---------|----------|
| `ImagePullBackOff` | Re-run `eval $(minikube docker-env)` then rebuild images |
| Pod stuck in `Pending` | Run `kubectl describe pod <name>` to see reason |
| Can't reach NodePort | Confirm Minikube IP with `minikube ip`; check firewall |
| `ErrImageNeverPull` | `imagePullPolicy: Never` is set — image must exist inside Minikube |
