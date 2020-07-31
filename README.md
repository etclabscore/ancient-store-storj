# ancient-store-storj
> Store your core-geth block data on Storj.

This program is a persistence manager for core-geth "ancient" block data using the distributed data provider Storj,
and can be used as a replacement for the core-geth (and go-ethereum) default of persisting ancient data in local flat, compressed files.

## Requirements
[core-geth](https://github.com/etclabscore/core-geth)
## Setup Environment Variables
```sh
export STORJ_API_KEY=xxx
export STORJ_SECRET=xxx
# satellie ex. us-central-1:7777 
export STORJ_SATELLITE=satellite-name:port
```

## Install

```sh
> git clone https://github.com/etclabscore/ancient-store-storj.git
> cd ancient-store-storj
> go build .
> ./ancient-store-storj --help
```

## Run ancient server
```
./ancient-store-storj --ipcpath ~./storj/ancient.ipc
```
### Run core-geth with ancient RPC path
```sh
geth --ancient.rpc ~/.storj/ancient.ipc
```
### Local testing 
For local testing launch your own localhost based storj network and satellite. Follow the instructions listed here
https://github.com/storj/storj/wiki/Test-network

