#!/usr/bin/env bash

# Copyright 2019 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

export KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..

# Build the version pinned in go.mod/go.sum. Keep the binary temporary so the
# script does not modify the developer's GOPATH.
cd ${KUBE_ROOT}
TOOL_DIR=$(mktemp -d)
trap 'rm -rf "${TOOL_DIR}"' EXIT
go build -o "${TOOL_DIR}/mdtoc" sigs.k8s.io/mdtoc

# Update tables of contents if necessary.
grep --include='*.md' -rl docs/enhancements/* -e '<!-- toc -->' | xargs "${TOOL_DIR}/mdtoc" --inplace
