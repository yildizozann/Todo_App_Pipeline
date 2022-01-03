# BUILD
FROM golang:latest 
ENV GO111MODULE=on \
    CGO_ENABLED=on \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /app
COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o todoapi

EXPOSE 5050

CMD ["./todoapi"]

