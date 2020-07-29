.PHONY: build

build:
	./deps.sh
	mkdir -p build/bin
	go build -o build/bin/ancient-storj ./main.go