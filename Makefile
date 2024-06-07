run: 
	@go run cmd/main.go 

push:
	@git add .
	@git commit -m $(message)
	@git push -u origin main