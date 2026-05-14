package erw

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
)

// ErrorWrapper Обретка ошибок.
type ErrorWrapper struct {
	statusCode      int
	responseMessage string
	loggingError    error
}

// Error вернуть текст ошибки.
func (e *ErrorWrapper) Error() string {
	return e.loggingError.Error()
}

// StatusCode Возвращает код состояния ошибки.
func (e *ErrorWrapper) StatusCode() int {
	return e.statusCode
}

// ResponseText Возвращает текст сообщения преднозначеного клиенту сделавшему запрос.
func (e *ErrorWrapper) ResponseText() string {
	if utf8.RuneCountInString(e.responseMessage) == 0 {
		if 0 <= e.statusCode && e.statusCode <= 16 {
			return gRPCStatusText(codes.Code(e.statusCode))
		} else if 100 <= e.statusCode && e.statusCode <= 511 {
			return http.StatusText(int(e.statusCode))
		}
	}
	return e.responseMessage
}

// ResponseText Возвращает сообщение преднозначеного
// клиенту сделавшему запрос в формате JSON.
func (e ErrorWrapper) ResponseJSON(options ...ResponseOption) []byte {
	pn := "error"
	for _, o := range options {
		o(&pn)
	}
	r, _ := json.Marshal(map[string]string{
		pn: e.ResponseText(),
	})
	return r
}

// LoggingError Возвращает ошибку преднозначеную для записи в журнал.
func (e *ErrorWrapper) LoggingError() error {
	return e.loggingError
}

// LoggingError Возвращает текст ошибки преднозначеный для записи в журнал.
func (e *ErrorWrapper) LoggingErrorText() string {
	if e.loggingError != nil {
		return e.loggingError.Error()
	}
	return ""
}

// New Конструктор обертки ошибок.
func New(options ...Option) *ErrorWrapper {
	e := &ErrorWrapper{}
	// options processing
	for _, o := range options {
		o(e)
	}
	return e
}
