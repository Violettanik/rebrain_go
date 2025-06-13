package main

import (
	"context"
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
	// Создаем контекст с таймаутом 100ms
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cache := NewCache()
	sem := make(chan struct{}, maxGoroutines) // Семафор для ограничения горутин

	// Запускаем 10 горутин для увеличения значения
	for i := 0; i < 10; i++ {
		select {
		case sem <- struct{}{}: // Захватываем слот если есть место
			go func() {
				defer func() { <-sem }() // Освобождаем слот
				cache.Increase(k1, step)
				
				select {
				case <-time.After(100 * time.Millisecond):
				case <-ctx.Done():
					return
				}
			}()
		case <-ctx.Done():
			break // Прерываем если контекст завершен
		}
	}

	// Запускаем 10 горутин для установки значения
	for i := 0; i < 10; i++ {
		select {
		case sem <- struct{}{}: // Захватываем слот если есть место
			go func(i int) {
				defer func() { <-sem }() // Освобождаем слот
				cache.Set(k1, step*i)
				
				select {
				case <-time.After(100 * time.Millisecond):
				case <-ctx.Done():
					return
				}
			}(i)
		case <-ctx.Done():
			break // Прерываем если контекст завершен
		}
	}

	// Ожидаем завершения контекста
	<-ctx.Done()

	// Выводим конечное значение
	fmt.Println("Final value:", cache.Get(k1))
	fmt.Println("Context error:", ctx.Err())
}
