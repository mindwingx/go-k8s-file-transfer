grpc:
	docker build -t grpc-service:latest grpc-service

http:
	docker build -t http-service:latest http-service

load:
	minikube image load grpc-service:latest && \
	minikube image load http-service:latest && \
	minikube image load alpine:latest

apply:
	kubectl apply -f grpc-deployment.yaml -f http-deployment.yaml

delete:
	kubectl delete -f grpc-deployment.yaml -f http-deployment.yaml