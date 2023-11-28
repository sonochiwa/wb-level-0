.SILENT: run

run:
	go build -o build/main cmd/main/main.go
	./build/main