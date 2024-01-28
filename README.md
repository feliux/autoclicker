# Autoclicker

Automate your mouse pointer (only for Linux).

```bash
$ go run main.go -h
$ env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/autoClicker main.go
```
