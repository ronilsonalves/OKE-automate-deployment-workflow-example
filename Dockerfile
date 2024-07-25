FROM golang:1.22.5-alpine3.19 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello main.go

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=build /app/hello /hello
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/hello"]