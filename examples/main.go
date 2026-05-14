package main

import (
	"fmt"
	"net/http"

	"github.com/skvdmt/erw"
)

// err1 Возвращает детализированную ошибку.
func err1() error {
	return erw.New(
		erw.Status(http.StatusBadRequest),
		erw.Response("param id not found"),
		erw.Logging(
			erw.Location("pkg", "method", "action"),
			erw.Error(fmt.Errorf("internal error for logging")),
			erw.SQL("select now();", "my param", 123),
		),
	)
}

// err2 Возвращает ошибку с минимальной детализацией.
func err2() error {
	return erw.New(
		erw.Status(http.StatusInternalServerError),
		erw.Logging(
			erw.Error(fmt.Errorf("internal error for logging")),
		),
	)
}

func main() {
	fmt.Println("Примеры использования:")
	fmt.Println("Обертка ошибки содержащая ее статус, краткое сообщение клиенту и детальное описание ошибки для записи в журнал ошибок.")
	fmt.Println()

	// Получаем детализированную ошибку.
	e := err1()
	err := e.(*erw.ErrorWrapper)
	fmt.Printf("Детали ошибки:\n")
	if err != nil {
		fmt.Printf("to response code: %d\n", err.StatusCode())
		fmt.Printf("to response text: %s\n", err.ResponseText())
		fmt.Printf("to response json: %s\n", err.ResponseJSON(erw.ErrorParam("msg")))
		fmt.Printf("to logs: %s\n", err.LoggingErrorText())
	}
	fmt.Println()

	// Получаем ошибку с минимальной детализацией.
	e = err2()
	err = e.(*erw.ErrorWrapper)
	fmt.Printf("Детали ошибки:\n")
	if err != nil {
		fmt.Printf("to response code: %d\n", err.StatusCode())
		fmt.Printf("to response text: %s\n", err.ResponseText())
		fmt.Printf("to response json: %s\n", err.ResponseJSON())
		fmt.Printf("to logs: %s\n", err.LoggingErrorText())
	}
}
