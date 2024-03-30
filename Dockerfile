#Build stage
FROM golang:1.20-alpine
# Set the working directory inside the container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist ./server
CMD ./out/dist