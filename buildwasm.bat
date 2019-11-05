set GOARCH=wasm
set GOOS=js
go build -o  .\docs\main.wasm .\wasm\main.go .\wasm\game.go .\wasm\events.go .\wasm\assets.go  .\wasm\debug.go .\wasm\common.go





