## File Transfer Using Go, gRPC, and Kubernetes

---

- Clone Repository

```shell
git clone git@github.com:mindwingx/go-k8s-file-transfer.git
```
- Start Minikube

```shell
minikube start --driver=docker
```

- Build Docker images

gRPC Service:

```shell
make grpc
```

HTTP Service:

```shell
make http
```

- Load Docker images into Minikube

```shell
make load
```

- Apply Kubernetes Deployments

```shell
make apply
```
- Port-forward Kubernetes

```shell
kubectl port-forward -n backend svc/http-k8s-srv 8080:8080
```

---

#### HTTP 

- handshake

```curl
curl http://localhost:8080/handshake
```


- upload image file

```curl
curl --location 'localhost:8080/upload' --form 'file=@"/path/to/image/file.png"'
```

