# Stage 1: Builder
FROM golang:1.22.5 AS builder
WORKDIR /Stocks
ADD . .
RUN go mod download
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
RUN go build -o app -ldflags="-s -w" . 

# Stage 2: Final image
FROM scratch
WORKDIR /
COPY --from=builder /Stocks/app .
ENTRYPOINT ["./app"] 
