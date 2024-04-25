FROM golang:latest AS build

WORKDIR /app

COPY . .
RUN go mod init
RUN go build -o src

# Stage 2: Create a minimal image to run the application
FROM scratch
COPY --from=build /app /app
CMD ["/app"]
