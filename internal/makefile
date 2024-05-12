build:
	@go build -o bin/tentrek cmd/api/main.go
run: build
	@./bin/tentrek
dockerrun:
	@docker run 
	-d \
	-p 8080:8080 \
	--env-file .env \
	--name tentrek \
	tentrek-api:0.1.0