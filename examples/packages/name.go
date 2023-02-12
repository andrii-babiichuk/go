package main

// Рекомендую розділяти пакети пробілами в залежності від лдерела, як у прикладі.
import (
	_ "encoding/json" // Якщо потрібно підключити пакет, але він безпосередньо не викликається, використовується "_".
	"fmt"             // Системні пакети.
	// . "fmt" // Імпорт пакету в поточну область бачення. НЕ ВАРТО ТАК РОБИТИ.

	v4 "github.com/google/uuid" // Зовнішній пакет, автоматична ковертація в URL для завантаження. v4 - aліас.

	"github.com/andrii-babiichuk/go/examples/packages/messager" // Пакет з власного додатку.
)

func main() {
	fmt.Println(v4.NewString())
	fmt.Println(messager.Hello())
}
