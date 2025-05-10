# Step 1: Build the Go binary
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Step 2: Use a minimal image to run the binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8081
CMD ["./main"]

