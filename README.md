# go-protobuf-builder

`go-protobuf-builder` is a starter project for defining and generating Protocol Buffers APIs with Buf and Go. It demonstrates how to organize shared models, API surfaces, and the supporting tooling required to produce language stubs.

## Project Layout

- `proto/`
  - `buf.yaml` / `buf.gen.yaml` configure the Buf workspace and code-generation plugins.
  - `example/common/v1/example_resources.proto` centralizes shared resource definitions such as `Example`, `AuditInfo`, and the lifecycle enum.
  - `example/api/v1/example_service.proto` exposes CRUD RPCs that operate on the shared `Example` resource.
  - `example/api/v1/example_audit_service.proto` adds a read-only service for retrieving audit information.
- `protogen.sh` wraps `buf generate` to emit Go stubs into `gen/go`.
- `Makefile` provides convenience targets for installing Buf, formatting, linting, generating, and cleaning artifacts.

## Prerequisites

- Go 1.22 or later available on the `PATH` (e.g., via `goenv`).
- Buf CLI. Run `make install` to install the pinned version into your Go bin directory.

## Common Commands

```bash
# Install the Buf CLI
make install

# Format proto sources in-place
make format

# Run Buf lint checks
make lint

# Generate Go code under gen/go
make build

# Remove generated code
make cleanup

# Run the example gRPC server (requires generated code)
make serve
```

## Generated Code

All generated artifacts live under `gen/go`, which is ignored by Git. Commit the generated files only if your workflow requires it.
