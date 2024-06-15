run: 
	@go run cmd/main.go 

push:
	@git add .
	@git commit -m $(message)
	@git push -u origin main

docker-build:
	@docker build -t docker-golang-auth .

docker-run:
	@docker run -p 8082:8081 docker-golang-auth
