package generator

import (
	"bytes"
	"os"
	"text/template"
)

// generate - основная функция генерации конфига из шаблона
func generate(tmpl string, outFilePath string, fields interface{}) error {
	// Создаем шаблон и парсим его
	t, err := template.New("config").Parse(tmpl)
	if err != nil {
		return err
	}

	// Применяем шаблон с переданными данными
	var buf bytes.Buffer
	if err := t.Execute(&buf, fields); err != nil {
		return err
	}

	// Записываем результат в файл
	if err := os.WriteFile(outFilePath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

// ConfigGenerate - функция подготовки данных и вызова генератора
func ConfigGenerate(tmpl string, outFilePath string) error {
	// Подготавливаем данные для шаблона
	data := struct {
		AppName    string
		Version    string
		Database   DatabaseConfig
		Monitoring MonitoringConfig
	}{
		AppName: "MyApp",
		Version: "1.0.0",
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "admin",
			Password: "secret",
		},
		Monitoring: MonitoringConfig{
			Enabled: true,
			Port:    9090,
		},
	}

	// Вызываем генератор с подготовленными данными
	return generate(tmpl, outFilePath, data)
}

// DatabaseConfig - структура для конфига БД
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// MonitoringConfig - структура для конфига мониторинга
type MonitoringConfig struct {
	Enabled bool
	Port    int
}
