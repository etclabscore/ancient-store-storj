module github.com/etclabscore/ancient-store-storj

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.18
	github.com/hashicorp/golang-lru v0.5.4
	github.com/mattn/go-colorable v0.1.7
	github.com/mattn/go-isatty v0.0.12
	github.com/tendermint/tendermint v0.33.6
	gopkg.in/urfave/cli.v1 v1.20.0
	storj.io/uplink v1.1.2
)

replace github.com/ethereum/go-ethereum => github.com/etclabscore/core-geth v1.11.10-0.20200731114149-9b4b7ebb3e4e
