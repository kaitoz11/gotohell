FROM golang:1.21-alpine3.18 AS build

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY ./ ./

RUN go build -o /build-bin

FROM scratch

COPY --from=build /build-bin /app

ENTRYPOINT [ "/app" ]