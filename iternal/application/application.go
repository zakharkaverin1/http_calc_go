package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/zakharkaverin1/http_calc_go/pkg/calculation"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json:"expression"`
}

type Result struct {
	Result string `json:"result"`
}

type Error struct {
	Error string `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Критическая ошибка: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: calculation.ErrExpressionNotString.Error()})
		return
	}

	reqExp := strings.ReplaceAll(request.Expression, " ", "")
	result, err := calculation.Calc(reqExp)
	if err != nil {
		log.Println(err)
		var errorResponse Error
		if err == calculation.ErrInvalidExpression {
			log.Printf("Критическая ошибка: Посторонние символы")
			errorResponse = Error{Error: err.Error()}
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else if err == calculation.ErrDivisionByZero {
			log.Printf("Критическая ошибка: Деление на ноль")
			errorResponse = Error{Error: calculation.ErrDivisionByZero.Error()}
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else if err == calculation.ErrSomethingWentWrong {
			log.Printf("Критическая ошибка: Что-то пошло не так")
			errorResponse = Error{Error: calculation.ErrSomethingWentWrong.Error()}
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	log.Printf("Успешное выполнение запроса")
	json.NewEncoder(w).Encode(Result{Result: fmt.Sprintf("%f", result)})
}

func (a *Application) Run() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Printf("Сервер запущен")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
		return nil
	} else {
		return err
	}
}
