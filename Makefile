build: build-linux-amd64

build-all: build-linux build-darwin

build-amd64: build-linux-amd64 build-darwin-amd64

build-arm64: build-linux-arm64 build-darwin-arm64

build-linux: build-linux-amd64 build-linux-arm64

build-darwin: build-darwin-amd64 build-darwin-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o color-picker-linux-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o color-picker-linux-arm64

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o color-picker-darwin-amd64

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o color-picker-darwin-arm64

release:
	goreleaser release --clean

clean:
	rm color-picker-*