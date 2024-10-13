package redis_models

import (
	"fmt"
	"net/http"
)

// Bind - метод для проверки и валидации данных перед их использованием
func (br *Book) Bind(r *http.Request) error {
	// Например, валидация данных
	if br.Title == "" {
		return fmt.Errorf("не указан заголовок")
	}
	return nil
}
