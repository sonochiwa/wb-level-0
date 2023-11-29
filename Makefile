.SILENT: run

run:
	go build -o build/main cmd/main/main.go
	./build/main

migrate:
	migrate -path ./migrations -database 'postgresql://root:root@localhost:5432/wb_db?sslmode=disable' up

migrate_down:
	migrate -path ./migrations -database 'postgresql://root:root@localhost:5432/wb_db?sslmode=disable' down