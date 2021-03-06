FROM golang:1.18.3-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o bin/server cmd/server.go



FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bin/server /server

USER nonroot:nonroot

ENTRYPOINT ["/server"]
