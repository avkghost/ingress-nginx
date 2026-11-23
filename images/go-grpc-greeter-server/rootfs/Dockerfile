ARG GOLANG_VERSION

FROM golang:1.27.4-alpine3.23@sha256:f70850ef6ce341e8b6427f677049145d63c0b7ce002d4d7e3d464828b1b96dcd AS build

WORKDIR /go/src/greeter-server

COPY . .
RUN go build -o /greeter-server

FROM scratch

COPY --from=build /greeter-server /

EXPOSE 50051

CMD ["/greeter-server"]
