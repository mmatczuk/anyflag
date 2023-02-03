# Redact

This example shows how `WithRedact` can be used to redact passwords.

1. Run `go run . --site https://user:password@www.google.com --realm http://foo.com --realm http://u:p@bar.com`, and see it prints:
  ```
	realm [http://foo.com,http://u:xxxxx@bar.com]
    site https://user:xxxxx@www.google.com
  ```