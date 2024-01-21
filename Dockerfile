FROM golang:1.19-buster
ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/server/main.go

ENTRYPOINT ["./app", "-config", "api-project-dev"]
