FROM golang:1.23.5-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN apk update && apk add make protobuf-dev g++ bash

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

ENV PATH=$PATH:/go/bin

RUN make proto
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM alpine:3.20 AS production

WORKDIR /app

RUN mkdir "assets"

COPY --from=builder /app/main .
RUN chmod +x main

EXPOSE 8082

CMD ["./main"]