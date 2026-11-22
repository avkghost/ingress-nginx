# Repository Guidelines

## Project Structure & Module Organization

This repository contains the Ingress NGINX Controller, Helm chart, deployment manifests, and supporting images. Go application code is organized under `cmd/`, `internal/`, and `pkg/`; unit tests live beside Go packages as `*_test.go`. End-to-end tests and fixtures are under `test/e2e/` and `test/`, while Helm chart templates and chart tests are in `charts/ingress-nginx/`. Documentation is in `docs/`, historical release notes in `changelog/`, and build/verification helpers in `build/` and `hack/`.

## Build, Test, and Development Commands

- `make test` — run Go unit tests through the repository’s containerized test script.
- `make static-check` — run boilerplate, code generation, `gofmt`, Go lint, Lua lint, and chart checks.
- `make helm-test` — run Helm unit tests for `charts/ingress-nginx`.
- `make kind-e2e-test` — run end-to-end tests on a kind cluster.
- `make dev-env` / `make dev-env-stop` — create or delete the local kind development cluster.
- `make verify-docs` — verify generated annotation documentation.

Use `make help` to see available targets. Read `docs/developer-guide/getting-started.md` before changing controller behavior.

## Coding Style & Naming Conventions

Follow standard Go conventions and run `gofmt`; keep package names lowercase and test files named `*_test.go`. Preserve existing Kubernetes manifest, Helm, Lua, and documentation formatting. Generated files should be updated through the project’s generation or verification targets, not edited manually when avoidable.

## Testing Guidelines

Add focused Go unit tests beside changed packages. Update relevant Helm tests under `charts/ingress-nginx/tests/` and e2e coverage under `test/e2e/` when behavior spans Kubernetes resources. Run the narrowest applicable test first, then `make static-check` and broader suites before submitting.

## Commit & Pull Request Guidelines

Use concise, area-prefixed commit subjects consistent with project history, such as `Go: ...`, `Tests: ...`, `Images: ...`, or `Template: ...`. Before opening a PR, sign the applicable CLA, describe the change and validation performed, and link related issues. All changes require review, at least one collaborator `/lgtm`, and an owner-provided `approve` label. This project is being retired and accepts only a limited set of patches; new feature work is not accepted.

## Security & Configuration Tips

Do not introduce credentials or cluster-specific configuration into the repository. Review deployment and admission-controller changes carefully, and consult the project’s security guidance and retirement notice before proposing operational changes.
