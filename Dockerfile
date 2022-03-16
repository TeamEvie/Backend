FROM golang:1.18
WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY prisma ./prisma
RUN go run github.com/prisma/prisma-client-go generate
COPY . .
RUN go build -o main
CMD ["./main"]