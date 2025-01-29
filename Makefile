migrate: build
	@cd $(dir $(lastword $(MAKEFILE_LIST))) && ./bin/migrater.exe --migrate

rollback: build
	@cd $(dir $(lastword $(MAKEFILE_LIST))) && ./bin/migrater.exe --rollback

build:
	@cd $(dir $(lastword $(MAKEFILE_LIST))) && go build -o ./bin/migrater.exe ./cmd/main/main.go

clean:
	@cd $(dir $(lastword $(MAKEFILE_LIST))) && rm -f ./bin/migrater.exe
