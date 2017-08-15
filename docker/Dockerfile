# Multi-stage build form
FROM golang:latest

# Copy in code...
RUN mkdir /app
ADD main.go /app/

# Build app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

# Could use something even smaller...
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app/
COPY --from=0 /app/hello .

# Expose ports for future REST Interface
EXPOSE 8080

CMD ["/app/hello"]
