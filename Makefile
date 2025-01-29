migrate: build
	@./bin/migrater.exe --migrate

rollback: build
	@./bin/migrater.exe --rollback

build:
	@go build -o ./bin/migrater.exe ./cmd/main/main.go

clean:
	@rm ./bin/migrater.exe
