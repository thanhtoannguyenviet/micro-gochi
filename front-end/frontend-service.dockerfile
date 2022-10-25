FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o frontend ./

RUN chmod +x /app/frontend

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/frontend /app
CMD ["/app/frontend"]