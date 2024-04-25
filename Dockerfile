FROM golang:latest AS build
WORKDIR /opt/app/
COPY app/ .

RUN cd api/src && ls -la && go build -o ../build/api main.go

# Stage 2: Create a minimal image to run the application
FROM scratch
WORKDIR /opt/api
COPY --from=build /opt/app/api/build/api api
CMD ["./api"]
