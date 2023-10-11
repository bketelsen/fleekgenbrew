set dotenv-load

default: build


lint:
  golangci-lint run

snapshot:
  goreleaser release --clean --snapshot

build:
  @source ./.env
  @go build -a -tags netgo -ldflags '-w -extldflags "-static"' github.com/bketelsen/fleekgenbrew

run: build
  ./fleekgenbrew

deploy: build
  flyctl deploy