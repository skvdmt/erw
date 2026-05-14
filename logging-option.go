package erw

import (
	"fmt"
	"strings"
)

// LoggingOption функциональные опции для логирования ошибки.
type LoggingOption func(internal *LoggingWrapper)

// LoggingWrapper информация функциональных опций
// внутреннего состояния обертки ошибок.
type LoggingWrapper struct {
	messages []string
}

// Location указания места возникновения ошибки.
func Location(locations ...string) LoggingOption {
	return func(internal *LoggingWrapper) {
		internal.messages = append(internal.messages,
			fmt.Sprintf("location: %s", strings.Join(locations, ".")))
	}
}

// Error передача внутренней ошибки возвращенной сторонним пакетом для логирования.
func Error(err error) LoggingOption {
	return func(lw *LoggingWrapper) {
		lw.messages = append(lw.messages, fmt.Sprintf("error: %s", err))
	}
}

// SQL передача sql запроса и его аргументов для логирования.
func SQL(query string, args ...any) LoggingOption {
	return func(lw *LoggingWrapper) {
		if len(args) > 0 {
			a := []string{}
			for _, arg := range args {
				a = append(a, fmt.Sprintf("%v", arg))
			}
			lw.messages = append(lw.messages,
				fmt.Sprintf("query: '%s' args: [%s]", query, strings.Join(a, ", ")))
			return
		}
		lw.messages = append(lw.messages,
			fmt.Sprintf("query: %s", query))
	}
}
