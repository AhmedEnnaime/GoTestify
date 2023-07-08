run:
	@docker-compose up -d

test: 
	@go test -v ./...

stop: 
	@docker-compose down