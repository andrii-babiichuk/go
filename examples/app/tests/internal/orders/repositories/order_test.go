// Приклад інтеграційних тестів для роботи з БД.
// Для виконання цих тестів потрібно мати тестову БД, де будуть проходити всі тести.
// На відміну від юніт тестів, тут можна тестувати не кожен метод окремо, а комбінацію,
// наприклад TestGetCreate.
// В даному прикладі міграція запускається в тесті, але при налаштуванні в CI
// краще міграцію виконувати окремо перед всіма тестами,
// щоб не мати проблем з асинхронним виконанням тестів.
// Інтеграційні тести варто писати не лише для БД, а і для будь яких інших систем
// з якими взаємодіє ваш додаток, наприклад хранилище файлів, поштовий клієнт та ін.

// Для різних видів тестів є сенс визначати теги збірки.
// Якщо у файлі визначений тег збірки, цей файл виключається зі збірки якщо він явно не вказаний.
// Запустити такий тест можна командою `go test ./... -tags=integration_test`
//go:build integration_test

package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"tests/internal/orders"
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
	m, err := migrate.New("file://../../../migrations", dbURI+"?sslmode=disable")
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

func TestGetCreate(t *testing.T) {
	ctx := context.Background()
	resetDB := initDB()
	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() {
		resetDB()
		pool.Close()
	}()

	repo := NewOrderRepository(pool)

	newOrder := orders.Order{
		UUID:        uuid.New(),
		Date:        time.Now().UTC(),
		Description: "test",
		Price:       100,
		Status:      orders.Completed,
	}

	t.Run("create and get new order", func(t *testing.T) {
		err = repo.OrderCreate(ctx, &newOrder)

		if err != nil {
			t.Errorf("failed to creare order with error: %v", err)
		}

		fmt.Println(newOrder)
		order, err := repo.OrderGet(ctx, newOrder.UUID)
		if err != nil {
			t.Errorf("failed to get order with error: %v", err)
		}

		if !reflect.DeepEqual(newOrder, *order) {
			t.Errorf("order data is corrupted; actual: %v, expected: %v", order, newOrder)
		}
	})
}
