FROM golang:1.18.5-alpine3.16 AS build

WORKDIR /tmp/go-build

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o ./out/echo-server .


FROM alpine:3.16

COPY --from=build /tmp/go-build/out/echo-server /app/echo-server

EXPOSE 8080

CMD ["/app/echo-server"]