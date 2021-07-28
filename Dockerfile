FROM golang:latest
WORKDIR /goProjects/src/flatsApp/users
COPY ./ ./
RUN go build -o bin/main cmd/main.go
CMD ["bin/main"]