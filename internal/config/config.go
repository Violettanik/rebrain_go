package config

import (
    "time"
)

// Config - основная конфигурационная структура приложения
type Config struct {
    // Общие настройки приложения
    AppName    string        `yaml:"name" json:"app_name"`
    Version    string        `yaml:"version" json:"version"`
    Timeout    time.Duration `yaml:"timeout" json:"timeout"`
    
    // Настройки базы данных
    Database struct {
        Host     string `yaml:"host" json:"host"`
        Port     int    `yaml:"port" json:"port"`
        Username string `yaml:"username" json:"username"`
        Password string `yaml:"password" json:"password"`
    } `yaml:"database" json:"database"`
    
    // Настройки мониторинга
    Monitoring struct {
        Enabled bool `yaml:"enabled" json:"enabled"`
        Port    int  `yaml:"port" json:"port"`
    } `yaml:"monitoring" json:"monitoring"`
    
    // Дополнительные настройки (пример)
    Cache struct {
        Enabled  bool          `yaml:"enabled" json:"enabled"`
        TTL      time.Duration `yaml:"ttl" json:"ttl"`
        MaxItems int           `yaml:"max_items" json:"max_items"`
    } `yaml:"cache" json:"cache"`
}
