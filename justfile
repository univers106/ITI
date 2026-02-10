# Запустить контейнеры (разработка)
[group('docker')]
docker-dev:
    docker-compose up

# Выключить все рабочие контейнеры
[group('docker')]
docker-down:
    docker-compose down
