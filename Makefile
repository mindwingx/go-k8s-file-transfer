grpc:
	docker build -t grpc-service:latest grpc-service

http:
	docker build -t http-service:latest http-service

apply-all:
	kubectl apply -f .

delete-all:
	kubectl delete -f .

apply:
	kubectl apply -f grpc-deployment.yaml -f http-deployment.yaml

delete:
	kubectl delete -f grpc-deployment.yaml -f http-deployment.yaml