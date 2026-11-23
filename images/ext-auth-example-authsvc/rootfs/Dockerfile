ARG GOLANG_VERSION

FROM golang:1.27.4-alpine3.23@sha256:f70850ef6ce341e8b6427f677049145d63c0b7ce002d4d7e3d464828b1b96dcd AS builder
RUN mkdir /authsvc
WORKDIR /authsvc
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o authsvc authsvc.go

FROM gcr.io/distroless/base-debian11@sha256:ac69aa622ea5dcbca0803ca877d47d069f51bd4282d5c96977e0390d7d256455
COPY --from=builder /authsvc/authsvc /
EXPOSE 8080
ENTRYPOINT ["/authsvc"]
