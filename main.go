package main

import "fmt"

func main() {
	readers := map[string]map[string][]string{
		"Читатель1": {
			"Книги":	{"Книга A", "Книга B"},
			"Периодика":	{"Журнал X"},
		},
		"Читатель2": {
			"Книги":	{"Книга C"},
			"Периодика":	{"Журнал Y", "Журнал Z"},
		},
		"Читатель3": {
			"Книги":	{},
			"Периодика":	{"Журнал W"},
		},
	}

	fmt.Println("Количество читателей с изданиями на руках:", len(readers))

	for reader, publications := range readers {
		totalPublications := len(publications["Книги"]) + len(publications["Периодика"])
		fmt.Printf("%s: %d изданий на руках\n", reader, totalPublications) 
	}
}
