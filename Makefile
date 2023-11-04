build:
	tinygo build -o tiny.wasm -target wasm cmd/auto/auto.go	

go-wasm:
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm cmd/auto/auto.go
