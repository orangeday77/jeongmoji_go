# jeongmoji_go

## run
```bash
$ go run main.go
```
## installed library
web/was server
```bash
$ go get github.com/gin-gonic/gin
```
configuration
```bash
$ go get github.com/joho/godotenv
```
live reload
```bash
$ go install github.com/cosmtrek/air@latest
$ air init
$ go mod init
$ go mod tidy
$ vim .air.toml
[log]
  time = true

[misc]
  clean_on_exit = true

[screen]
  clear_on_rebuild = true
$ air -c .air.toml
```