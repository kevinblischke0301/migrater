run: build
	@./bin/migrater.exe

build:
	@go build -o ./bin/migrater.exe ./cmd/main.go

clean:
	@rm ./bin/migrater.exe
