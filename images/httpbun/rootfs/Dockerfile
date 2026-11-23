# Copyright 2019 The Kubernetes Authors. All rights reserved.
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

FROM golang:1.27.4@sha256:205c7636d74e0ffbe3d007162df7d74f859f5a8c8322dfb9f907977686c126cf AS builder

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV HTTPBUN_COMMIT=a6b387c438b664322734250e426b6966966c2fe4

RUN git init /go/httpbun \
    && git -C /go/httpbun fetch --depth=1 https://github.com/sharat87/httpbun.git $HTTPBUN_COMMIT \
    && git -C /go/httpbun checkout --detach $HTTPBUN_COMMIT
WORKDIR /go/httpbun

RUN CGO_ENABLED=0 GOOS=linux go build \
    -installsuffix cgo \
    -trimpath \
    -o bin/httpbun .

FROM scratch
COPY --from=builder /go/httpbun/bin/httpbun httpbun
ENV HTTPBUN_BIND=0.0.0.0:80
EXPOSE 80
ENTRYPOINT ["/httpbun"]
