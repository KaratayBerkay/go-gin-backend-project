# Use the latest Go image
FROM golang:latest

# Set working directory inside the container
WORKDIR /app

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go mod init github.com/KaratayBerkay/go-gin-backend-project || echo "Already exists"
RUN go mod tidy
RUN go build -o app main.go

# Set the correct CMD path
CMD ["./app"]
