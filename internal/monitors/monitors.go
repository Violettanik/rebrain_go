package monitors

// Monitor интерфейс для мониторинга приложения
type Monitor interface {
	// Start запускает мониторинг
	Start() error
	
	// Stop останавливает мониторинг
	Stop() error
	
	// Log отправляет лог-сообщение
	Log(message string) error
	
	// Metric отправляет метрику
	Metric(name string, value float64, tags map[string]string) error
}
