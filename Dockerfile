FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go build -o myapp

# Stage 2: Create a minimal image to run the application
FROM scratch
COPY --from=build /app/myapp /myapp
CMD ["/myapp"]
