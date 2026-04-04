FROM golang:1.26.1-alpine as build

WORKDIR /server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/nexus_server .

FROM golang:1.26.1-alpine

COPY --from=build bin/nexus_server bin/nexus_server

ENTRYPOINT ["./bin/nexus_server"]