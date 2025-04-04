
```bash
$ go run cmd/cli.go -h
$ env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/autoclicker cmd/cli.go

$ goreleaser release --snapshot --clean
$ goreleaser check
$ goreleaser build --single-target

$ export GITHUB_TOKEN="YOUR_GH_TOKEN"
$ git tag -a v0.1.0 -m "First release"
$ git push origin v0.1.0
$ goreleaser release
```
