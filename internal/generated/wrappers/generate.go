//go:generate gowrap gen -p /root/go/module08/05_task/internal/monitors -i Monitor -t log -o monitor_with_log.go
//go:generate gowrap gen -p /root/go/module08/05_task/internal/monitors -i Monitor -t prometheus -o monitor_with_metrics.go

package wrappers

// Импортируем пакеты, чтобы go mod мог их отслеживать
import (
	_ "github.com/hexdigest/gowrap"
	_ "05_task/internal/monitors"
)
