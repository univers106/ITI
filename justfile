# Запустить контейнеры (разработка)
[group('docker')]
docker-dev:
    docker-compose up

# Выключить все рабочие контейнеры
[group('docker')]
docker-down:
    docker-compose down

# Запустить линтер Go
[group('go')]
go-lint:
    cd backend && golangci-lint run ./...
