# Quick Start

0. `go run <*.go>`
- run go script

1. `go mod init <src path>`
- Initiate go project in `src/`

1. `go build <file>`
- create executable file in `<src path>`. Run on terminal, e.g. `./src`

1. `env GOOS=<os> GOARCH=<architecture> go build src/exploration.go <file>` ; get <os> and <arch> from `go tool dist list`
- create executable file in `<src path>`. Run on terminal, e.g. `./src`

1. `go install`
- create executable file in `GOPATH/bin`. Run on terminal, e.g. `./src`
