FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY go.mod go.sum ./

COPY .env.example ./ 

RUN go mod download

COPY . .

RUN go build -o api ./cmd/api/main.go

EXPOSE 3001

CMD ["./api"]