package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	k1            = "key1"
	step          = 7
	maxGoroutines = 4
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]int),
	}
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func main() {
	cache := NewCache()
	sem := make(chan struct{}, maxGoroutines) // Семафор для ограничения горутин
//	done := make(chan struct{})              // Канал для ожидания завершения

	// Запускаем 10 горутин для увеличения значения
	for i := 0; i < 10; i++ {
		sem <- struct{}{} // Захватываем слот
		go func() {
			defer func() { <-sem }() // Освобождаем слот
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
//			done <- struct{}{}
		}()
	}

	// Запускаем 10 горутин для установки значения
	for i := 0; i < 10; i++ {
		sem <- struct{}{} // Захватываем слот
		go func(i int) {
			defer func() { <-sem }() // Освобождаем слот
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
//			done <- struct{}{}
		}(i)
	}

	// Ожидаем завершения всех 20 горутин
//	for i := 0; i < 20; i++ {
//		<-done
//	}
        for len(sem) > 0{
                time.Sleep(time.Millisecond * 100)
        }
	// Выводим конечное значение
	fmt.Println(cache.Get(k1))
}
