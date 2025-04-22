package main

import "fmt"

func main() {
    workingDays := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница"}
    weekendDays := []string{"Суббота", "Воскресенье"}

    // Объединяем слайсы с выходными днями и рабочими днями
    allDays := append(workingDays, weekendDays...)

    fmt.Println("Все дни недели:", allDays)
}
