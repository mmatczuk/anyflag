# URL

This example shows how `anytype` can be used for adding a custom validation to a flag.

1. Run `go run . --auth user:pass`, and see it prints:
  ```
  user:pass
  ```

1. Run `go run . --auth user:`, and see it prints:
  ```
  Error: invalid argument "user:" for "--auth" flag: password is required
  ```

1. Run `go run . --auth :pass`, and see it prints:
  ```
  Error: invalid argument "user:" for "--auth" flag: password is required
  ```
