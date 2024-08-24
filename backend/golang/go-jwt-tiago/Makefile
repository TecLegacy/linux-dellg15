build:
	@go build -o ./bin/ecom ./cmd/main.go

run: build
	@./bin/ecom

test:	
	@go test -v ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down


fex:
	@migrate -database "mysql://root:mypass@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations force 20240614063742