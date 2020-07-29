.PHONY: build

build:
	./deps.sh
	mkdir -p build/bin
	go build -o build/bin/ancient-store-storj ./main.go