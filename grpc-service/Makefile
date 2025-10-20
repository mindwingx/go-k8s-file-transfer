proto:
	protoc -I pb \
    	pb/*.proto \
    	--go_opt paths=source_relative \
    	--go-grpc_opt paths=source_relative \
    	--go_out api \
    	--go-grpc_out api
