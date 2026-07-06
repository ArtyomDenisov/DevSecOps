package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword принимает чистый пароль и возвращает его хэш
func HashPassword(password string) (string, error) {
	// 14 — это стоимость (cost) или сложность вычислений.
	// Чем выше число, тем дольше генерируется хэш и тем сложнее его взломать.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash сравнивает введенный пароль с хэшем из базы данных
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Возвращает true, если пароли совпадают
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// ... проверка логина и пароля ...

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    "уникальный_длинный_случайный_токен",
		Path:     "/",
		MaxAge:   3600,                    // 1 час
		HttpOnly: true,                    // Защита от кражи через JavaScript (XSS)
		Secure:   true,                    // Передача только по HTTPS
		SameSite: http.SameSiteStrictMode, // Защита от CSRF
	}

	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "Вход успешно выполнен!")
}
