### Запуск тестів

Для запуску АПІ та інтеграційних тестів, необхідно підняти базу даних локально.

```shell
docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -e POSTGRES_PASSWORD=postgres -d postgres
```

Запуск тестів

- Юніт `go test ./... -tags=unit_test`
- Інтеграційні `go test ./... -tags=integration_test`
- АПІ `go test ./... -tags=api_test`