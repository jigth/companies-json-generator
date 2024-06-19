FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /companies-generator ./src

ENTRYPOINT ["/companies-generator"]
