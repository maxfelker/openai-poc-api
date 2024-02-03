FROM golang:alpine as build 
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY src/ src/
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o bin/openai-poc-api main.go

FROM alpine:latest as release  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/bin ./bin
ENTRYPOINT ./bin/openai-poc-api 