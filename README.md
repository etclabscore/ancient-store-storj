# ancient-store-storj
> Store your core-geth block data on Storj.

This program is a persistence manager for core-geth "ancient" block data using the distributed data provider Storj,
and can be used as a replacement for the core-geth (and go-ethereum) default of persisting ancient data in local flat, compressed files.

## Install

```sh
> git clone https://github.com/etclabscore/ancient-store-storj.git
> cd ancient-store-storj
> go build .
> ./ancient-store-storj --help
```



