migrateup:
	goose -dir ./app/infra/persistence/postgres/migrations postgres "user=postgres dbname=desafiogo password=postgres sslmode=disable" up	
migratedown:
	goose -dir ./app/infra/persistence/postgres/migrations postgres "user=postgres dbname=desafiogo password=postgres sslmode=disable" down	

test-coverage:
	@echo "==> Checking test coverage..."
	go get github.com/kyoh86/richgo
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html