#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(git rev-parse --show-toplevel)
cd "$ROOT"
export GOTOOLCHAIN="${GOTOOLCHAIN:-go$(tr -d '\n' < GOLANG_VERSION)}"

command -v govulncheck >/dev/null 2>&1 || {
  echo "govulncheck is required (go install golang.org/x/vuln/cmd/govulncheck@<version>)" >&2
  exit 1
}
command -v gosec >/dev/null 2>&1 || {
  echo "gosec is required (go install github.com/securego/gosec/v2/cmd/gosec@<version>)" >&2
  exit 1
}
command -v osv-scanner >/dev/null 2>&1 || {
  echo "osv-scanner is required (go install github.com/google/osv-scanner/v2/cmd/osv-scanner@<version>)" >&2
  exit 1
}

mapfile -t MODULES < <(find . -name go.mod -not -path './vendor/*' -print | sort)

for manifest in "${MODULES[@]}"; do
  module_dir=$(dirname "$manifest")
  echo "Scanning Go module: $module_dir"
  (
    cd "$module_dir"
    # Module-level scanning avoids Go-version-specific SSA failures while still
    # checking every resolved dependency against the Go vulnerability database.
    scan_dir=$(go list -f '{{.Dir}}' ./... | head -n 1)
    govulncheck -scan=module -C "$scan_dir"
    # Keep gosec findings visible without blocking dependency/CVE gates; the
    # Trivy, govulncheck, and OSV-Scanner checks are the release blockers.
    gosec -no-fail -exclude=G109,G601,G104,G204,G304,G306,G307 \
      -tests=false -exclude-dir=test -exclude-dir=docs ./...
  )
done

osv-scanner scan source -r .
