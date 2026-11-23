# Copyright 2021 The Kubernetes Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG GOLANG_VERSION

FROM --platform=$BUILDPLATFORM golang:1.27.4@sha256:205c7636d74e0ffbe3d007162df7d74f859f5a8c8322dfb9f907977686c126cf AS builder

ARG BUILDPLATFORM
ARG TARGETARCH


WORKDIR /workspace
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -a -o kube-webhook-certgen main.go

FROM --platform=$BUILDPLATFORM gcr.io/distroless/static:nonroot@sha256:8ce6dd87e9823c5b48dcfb2c8604714ccb351ce59607327bcc4957448da279b5
ARG BUILDPLATFORM
ARG TARGETARCH

LABEL org.opencontainers.image.source=https://github.com/kubernetes/ingress-nginx

WORKDIR /
COPY --from=builder /workspace/kube-webhook-certgen /kube-webhook-certgen
USER 65532:65532
ENTRYPOINT ["/kube-webhook-certgen"]
