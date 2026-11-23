ARG GOLANG_VERSION

FROM golang:1.27.4-alpine3.24@sha256:70bd7ddea8b6cbe5d6943cde279d6b67feb10cc924fb9100d1d8757c1575b280 AS build

WORKDIR /go/src/greeter-server

COPY . .
RUN go build -o /greeter-server

FROM scratch

COPY --from=build /greeter-server /

EXPOSE 50051

CMD ["/greeter-server"]
