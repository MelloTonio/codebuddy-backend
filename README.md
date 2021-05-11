<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-56%25-brightgreen.svg?longCache=true&style=flat)</a> <a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://goreportcard.com/badge/github.com/MelloTonio/desafiogo)</a>

# Bank API

## Instructions to Run

### Prod
- toDo

### Local
- Database: `docker-compose up` will start the PostgreSQL DB
- `make migrateup` will migrate the database structure
- `go run main.go` will start the api

## Routes
- base_url: http://localhost:3001

### Accounts
 - POST `{{base_url}}/accounts/create` -  Create one account in the database
 - GET `{{base_url}}/accounts/{accountID}` - Get one account by ID
 - GET `{{base_url}}/accounts/all` - List all accounts
 
### Transfer
- GET `{{base_url}}/transfers/{accountID}` - List all transfer by account ID
- POST `{{base_url}}/transferTo` - Perform a transfer between two accounts (Login Required) (QueryParams: accountID_destination, amount)

### Auth
- POST `{{base_url}}/accounts/login` - Authenticate the account and returns JWT token


