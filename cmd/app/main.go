package main

import (
	"log"
	"04_task/internal/generator"
)

func main() {
	// Раскомментируйте нужную задачу при работе
	
	Task04() // Генерация маршаллера и запуск бенчмарков
}

func Task04() {
	// Генерируем код маршаллера
	if err := generator.GenerateMarshaller(); err != nil {
		log.Fatalf("Error generating marshaller: %v", err)
	}

	log.Println("Marshaller code generated successfully")

}
