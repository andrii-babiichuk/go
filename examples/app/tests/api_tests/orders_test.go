// Приклад API тестів
// Використовуються для перевірки функціоналу на всіх рівнях від інтерфейсу до взаємодії з іншими системами.
// Є два варіанти реалізувати такі тести
// 1) створити http клієнт та робити запити напряму на запущену систему
// 2) підняти тестовий сервер за допомогою тестів, та робити запити за допомогою бібліотек, (наприклад github.com/steinfletcher/apitest)
// Перший варіант кращий тим що тестується реальне середовище, його не потрібно додатково налаштовувати в тестах.
// Другий варіант дозволяє більш гнучко налаштовувати сервер, при потребі частину даних можна емулювати, але
// це вимагає додаткового налаштування в тестах, та не дозволяє перевіряти працюючий сервер, що інколи може бути корисно.
// Тут розглядається другий варіант.

// Для різних видів тестів є сенс визначати теги збірки.
// Якщо у файлі визначений тег збірки, цей файл виключається зі збірки якщо він явно не вказаний.
// Запустити такий тест можна командою `go test ./... -tags=api_test`
//go:build api_test

package api_tests

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/steinfletcher/apitest"

	"tests/handlers"
	"tests/internal/orders/repositories"
	"tests/internal/orders/services"
)

var dbURI string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file, using system env variables")
	}

	dbURI = os.Getenv("DATABASE_URL")
}

func initDB() func() {
	m, err := migrate.New("file://../migrations", dbURI+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	log.Println("db scheme is up to date")

	return func() {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}

func TestAll(t *testing.T) {
	// Перед запуском тестів необхідно підключитись до БД
	// Міграцію можна виконувати за межами тестів, але в даному випадку вона запускається в тестах.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resetDB := initDB()
	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() {
		resetDB()
		pool.Close()
	}()

	// Ініціалізація всіх необхідних ресурсів (репозиторій, сервіс, хендлер)
	orderRepository := repositories.NewOrderRepository(pool)
	orderService := services.NewOrderService(orderRepository)
	h := handlers.NewHttpHandler(orderService)

	router := httprouter.New()
	h.Register(router)

	// Тестування створення ордеру неуспішне
	apitest.New().
		Handler(router).
		Post("/api/v1/orders").
		Body(`{"uuid": "548f8ed1-77ee-424a-b004-16cb25430b81", "description": "test", "price": 10, "date": "invalid date"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Тестування створення ордеру
	apitest.New().
		Handler(router).
		Post("/api/v1/orders").
		Body(`{"uuid": "548f8ed1-77ee-424a-b004-16cb25430b88", "description": "test", "price": 10, "date": "2023-04-10T13:56:38.452Z"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	// Тестування отримання ордеру
	apitest.New().
		Handler(router).
		Get("/api/v1/orders/548f8ed1-77ee-424a-b004-16cb25430b88").
		Expect(t).
		Body(`{"uuid": "548f8ed1-77ee-424a-b004-16cb25430b88", "description": "test", "price": 10, "date": "2023-04-10T13:56:38.452Z"}`).
		Status(http.StatusOK).
		End()
}
