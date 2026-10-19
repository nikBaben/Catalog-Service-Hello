# Catalog-Service-Hello

Простой HTTP-сервис на Go: GET / возвращает текст `hello from catalog`.
На сервере доступен через NGINX: http://45.155.204.229/catalog/.

## Локальный запуск

```bash
docker compose up --build -d
curl http://127.0.0.1:38082/
docker compose down
```

## Тесты

```bash
go test ./...
```

Тесты выполняются в CI этого репозитория до публикации образа. В Deployment тестов и проверок готовности нет.

## Выпуск

Push в main запускает тесты, сборку Linux amd64-образа и публикацию в `ghcr.io/nikbaben/catalog-service-hello`.
Workflow отдельным SSH deploy key обновляет только `catalog.env` в [Deployment-Hello](https://github.com/nikBaben/Deployment-Hello).
Deployment применяет только изменённый сервис: `docker compose pull catalog` и `docker compose up -d --no-deps catalog`.
Контейнеры Account-Service и NGINX при обычном выпуске Catalog не пересоздаются.

`DEPLOYMENT_REPO_SSH_KEY` хранится в Actions Secrets. Ключ имеет право записи только в Deployment-Hello. GITHUB_TOKEN для публикации образа GitHub создаёт автоматически.
У серверного workflow уже есть свой SSH-доступ к Cloud.ru; передавать его в этот репозиторий не требуется.

Dockerfile компилирует статическую Go-программу и помещает её в минимальный scratch-образ. В рабочем контейнере нет компилятора, shell и тестов.
