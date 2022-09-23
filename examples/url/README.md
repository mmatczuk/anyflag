# URL

This example shows how `anytype` can be used for *url.URL (or pointer to any other type).

1. Run `go run . --site https://www.google.com?q=foo --realm http://foo.com --realm http://bar.com`, and see it prints:
  ```
  https://www.google.com?q=foo
  [http://foo.com http://bar.com]
  ```