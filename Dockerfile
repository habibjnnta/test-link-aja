FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

# RUN go test ./tests/...

RUN go build -o test-link-aja

CMD ["./test-link-aja"]