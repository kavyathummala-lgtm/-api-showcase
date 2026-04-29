# Important Commands Reference

---

## Git Commands

```
git init                          → start tracking a folder
git remote add origin <url>       → connect to GitHub
git add .                         → stage all changes
git commit -m "message"           → save with message
git push                          → send to GitHub
git pull                          → get latest from GitHub
git status                        → see what changed
git branch <name>                 → create new branch
git checkout <name>               → switch branch
git merge <name>                  → merge branch
git stash                         → hide unfinished work
git stash pop                     → bring it back
```

---

## Docker Commands

```
docker build -t name .            → build image
docker run -p 8000:8000 name      → run container
docker images                     → see all images
docker ps                         → see running containers
docker push name                  → push image to Docker Hub
```

---

## Kubernetes Commands

```
kubectl apply -f file.yaml        → deploy using YAML file
kubectl get pods                  → see all running pods
kubectl get services              → see all services
kubectl delete pod <name>         → delete a pod
kubectl logs <pod-name>           → see pod logs
kubectl describe pod <name>       → see pod details
minikube start                    → start Kubernetes locally
minikube service <name> --url     → get service URL
```

---

## Helm Commands

```
helm create <name>                → create new chart
helm install <name> <chart>       → deploy chart
helm upgrade <name> <chart>       → update deployment
helm uninstall <name>             → remove deployment
helm list                         → see all deployments
helm template <name> <chart>      → preview output
```

---

## The Full Flow In Commands

```
Write code
      ↓
git add . → git commit -m "msg" → git push
      ↓
GitHub Actions builds Docker image automatically
      ↓
helm upgrade python-rest api-chart/
      ↓ (with ArgoCD this happens automatically)
kubectl get pods → check everything running
```

---

## Key Commands To Memorize

| Task | Command |
|------|---------|
| Save and push code | `git add .` → `git commit -m "msg"` → `git push` |
| Check pods | `kubectl get pods` |
| Deploy with Helm | `helm install name chart/` |
| Update deployment | `helm upgrade name chart/` |
| See Helm deployments | `helm list` |
| Get service URL | `minikube service name --url` |
