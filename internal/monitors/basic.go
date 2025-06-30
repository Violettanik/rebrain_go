package monitors

import "log"

type BasicMonitor struct{}

func (m *BasicMonitor) Start() error {
	log.Println("Monitor started")
	return nil
}

func (m *BasicMonitor) Stop() error {
	log.Println("Monitor stopped")
	return nil
}

func (m *BasicMonitor) Log(message string) error {
	log.Println("LOG:", message)
	return nil
}

func (m *BasicMonitor) Metric(name string, value float64, tags map[string]string) error {
	log.Printf("METRIC: %s=%.2f tags=%v", name, value, tags)
	return nil
}
