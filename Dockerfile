# Import go image
FROM golang:1.14-alpine
# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to currennt directory
COPY . .
# Run command to nstall the dependencies
RUN go mod download
RUN go build -o main .
EXPOSE 8080