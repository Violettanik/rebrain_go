package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
        _ "net/http/pprof"
	"06_task/internal/pkg/util"
)

func main() {
        go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// Настройка маршрутов
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)

	// Конфигурация сервера
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр длины из query string (по умолчанию 100)
	length := 100
	if l := r.URL.Query().Get("length"); l != "" {
		_, err := fmt.Sscanf(l, "%d", &length)
		if err != nil {
			http.Error(w, "Invalid length parameter", http.StatusBadRequest)
			return
		}
	}

	// Генерируем строку с помощью Pad
	result := util.Pad("Hello", length)

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, result)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "ok"}`)
}
