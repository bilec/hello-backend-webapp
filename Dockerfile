FROM golang:alpine AS builder
ENV APP hello-world-webapp
ENV GOOS linux
ENV GOARCH amd64
WORKDIR /webapp
COPY . .
RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -o server .

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /webapp/server .
CMD ["./server"]
