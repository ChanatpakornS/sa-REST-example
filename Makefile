run:
	go run cmd/main.go

compose:
	docker compose -f dockers/docker-compose.yml up --build