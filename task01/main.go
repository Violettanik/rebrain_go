package main

import "fmt"

func main() {
	daysOfWeek := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}
	workingDays := []string{}

	for _, day := range daysOfWeek {
		if day != "Суббота" && day != "Воскресеье" {
			workingDays = append(workingDays, day)
		}
	}

	weekendDays := []string{"Суббота", "Воскресенье"}

	fmt.Println("Рабочие дни:", workingDays)
	fmt.Println("Выходные дни:", weekendDays)
}
