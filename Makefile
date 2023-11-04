build:
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm cmd/auto/auto.go

go-wasm:
	tinygo build -o tiny.wasm -target wasm cmd/auto/auto.go	
