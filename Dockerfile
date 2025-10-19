# build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o weather-service .

# run stage
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/weather-service .

# execution perm needed
RUN chmod +x weather-service

EXPOSE 8080

# set the binary as entry point
ENTRYPOINT ["./weather-service"]

# default args
CMD ["--mode", "server"]
