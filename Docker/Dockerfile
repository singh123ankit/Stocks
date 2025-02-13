# Stage 1: Builder
FROM golang:1.22.5 AS builder
WORKDIR /Stocks
ADD . .
RUN go mod download
ENV GOOS=linux \
    CGO_ENABLED=0
RUN go build -o app .

# Stage 2: Final image
FROM scratch
WORKDIR /
ADD --from=builder /Stocks/app .
ENTRYPOINT ["./app"] 
