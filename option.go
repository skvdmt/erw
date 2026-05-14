package erw

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
)

// Option функциональная опция обертки ошибок.
type Option func(e *ErrorWrapper)

type Code interface {
	int | codes.Code
}

// Status Устанавливает код статуса и сообщение ему соответствующее.
func Status[T Code](code T) Option {
	return func(e *ErrorWrapper) {
		e.statusCode = int(code)
		if 0 <= code && code <= 16 {
			e.responseMessage = gRPCStatusText(codes.Code(code))
		} else if 100 <= code && code <= 511 {
			e.responseMessage = http.StatusText(int(code))
		} else {
			panic(fmt.Sprintf("unknown status code %d", code))
		}
	}
}

// Response установка описания ошибки, которое отдается клиенту в ответ на его запрос.
func Response(message string) Option {
	return func(e *ErrorWrapper) {
		e.responseMessage = message
	}
}

// Logging установка расширенного описания ошибки для логирования в журнале ошибок.
// Принимает функциональные опции LoggingOption
func Logging(options ...LoggingOption) Option {
	return func(e *ErrorWrapper) {
		lw := &LoggingWrapper{}
		for _, option := range options {
			option(lw)
		}
		e.loggingError = fmt.Errorf("%s", strings.Join(lw.messages, "; "))
	}
}
