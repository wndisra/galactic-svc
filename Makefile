test:
	@echo 'Executing all unit tests ...'
	go test ./... -count=1 -race -cover -covermode=atomic

run-server:
	@echo 'Make sure Air installed (https://github.com/cosmtrek/air#installation) ...'
	air -d

api-doc:
	swag init -d ./cmd/server/,./internal/spaceship/