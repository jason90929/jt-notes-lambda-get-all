FROM golang:1.22.1

# Set destination for COPY
WORKDIR /jt-notes-lambda-get-all

COPY . .

RUN go mod download
RUN go build -o main ./cmd/lambda

EXPOSE 8080

# Run
CMD ["./main"]
