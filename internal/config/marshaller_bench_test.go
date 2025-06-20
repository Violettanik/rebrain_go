package config

import (
	"04_task/internal/convertor"
	"testing"
	"time"
)

func BenchmarkReflectionStructToMap(b *testing.B) {
	cfg := Config{
		AppName: "benchmark_app",
		Version: "1.0.0",
		Timeout: 30 * time.Second,
		Database: struct {
			Host     string `yaml:"host" json:"host"`
			Port     int    `yaml:"port" json:"port"`
			Username string `yaml:"username" json:"username"`
			Password string `yaml:"password" json:"password"`
		}{
			Host:     "localhost",
			Port:     5432,
			Username: "bench_user",
			Password: "bench_password",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = convertor.StructToMap(cfg)
	}
}

func BenchmarkGeneratedStructToMap(b *testing.B) {
	cfg := Config{
		AppName: "benchmark_app",
		Version: "1.0.0",
		Timeout: 30 * time.Second,
		Database: struct {
			Host     string `yaml:"host" json:"host"`
			Port     int    `yaml:"port" json:"port"`
			Username string `yaml:"username" json:"username"`
			Password string `yaml:"password" json:"password"`
		}{
			Host:     "localhost",
			Port:     5432,
			Username: "bench_user",
			Password: "bench_password",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.StructToMap()
	}
}
