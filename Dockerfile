# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

# RUN go build -o ./cmd/main.go /smart_wallet
RUN go build -o /smart_wallet ./cmd/main.go

EXPOSE 8080

CMD [ "/smart_wallet" ]
