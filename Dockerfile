FROM golang:latest
WORKDIR /app
COPY . .
RUN go build ./cmd/main.go
RUN ls
EXPOSE 5000
CMD ["./main"]
