migrateup:
	goose -dir ./app/infra/persistence/postgres/migrations postgres "user=postgres dbname=desafiogo password=postgres sslmode=disable" up	
migratedown:
	goose -dir ./app/infra/persistence/postgres/migrations postgres "user=postgres dbname=desafiogo password=postgres sslmode=disable" down	
