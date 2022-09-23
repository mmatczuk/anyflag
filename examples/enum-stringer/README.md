# Enum Stringer

This example shows how `anytype` can be used for enums. It requires Go [stringer tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) to be installed.

1. Install stringer `go install golang.org/x/tools/cmd/stringer@latest`
1. Regenerate stringer files if needed `go generate`
1. Run `go run . --pill Placebo --jar Aspirin --jar Ibuprofen`, and see it prints:
  ```
  Placebo
  [Aspirin Ibuprofen]
  ```