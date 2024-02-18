FROM golang:1.20 as builder

WORKDIR /app
COPY . . 
RUN GOOS=linux GIARCH=amd64 CGO_ENABLED=0 go build -o server ./cmd

FROM scratch 
COPY --from=builder /app/server .
CMD ["./server"]