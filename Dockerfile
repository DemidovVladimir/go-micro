FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

ARG project

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o server ./cmd/web

FROM alpine:latest AS production

COPY --from=builder /app .

EXPOSE 4000

ENTRYPOINT ["./server"]

