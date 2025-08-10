FROM golang:alpine AS builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download && go mod verify

COPY ../../cmd/ ./cmd/
COPY ../../internal/ ./internal
COPY ../../docs/ ./docs

RUN mkdir -p ./build/output

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/output/binary ./cmd/app

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /root/

COPY --from=builder /app/build/output/binary .
COPY --from=builder /app/docs/ ./docs/

RUN mkdir -p logs

EXPOSE 8000 8080

CMD ["./binary"]
