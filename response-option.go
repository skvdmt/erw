package erw

// ResponseOption функциональная опция ответов.
type ResponseOption func(s *string)

// ErrorParam Принимает строку name  устанавливает ее
// в качестве ключа текста ошибки в JSON структуре.
func ErrorParam(name string) ResponseOption {
	return func(s *string) {
		*s = name
	}
}
