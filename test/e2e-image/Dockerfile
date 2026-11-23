ARG E2E_BASE_IMAGE=registry.k8s.io/ingress-nginx/e2e-test-runner:v2.2.9@sha256:6eda6a8d17ff65c5af647abb0714b882047b77d18161712d77daf5f610fd4020
FROM ${E2E_BASE_IMAGE} AS BASE

FROM alpine:3.24.4@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

RUN apk update \
	&& apk upgrade && apk add -U --no-cache \
  ca-certificates \
  bash \
  tzdata \
  libc6-compat \
  openssl \
  sqlite-dev

COPY --from=BASE /go/bin/ginkgo /usr/local/bin/
COPY --from=BASE /usr/local/bin/helm /usr/local/bin/
COPY --from=BASE /usr/local/bin/kubectl /usr/local/bin/
COPY --from=BASE /usr/bin/cfssl /usr/local/bin/
COPY --from=BASE /usr/bin/cfssljson /usr/local/bin/

COPY . /

CMD [ "/e2e.sh" ]
