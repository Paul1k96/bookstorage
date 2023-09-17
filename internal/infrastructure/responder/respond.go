package responder

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Paul1k96/bookstorage/internal/infrastructure/response"
)

// Responder интерфейс, описывающий парсер для выдачи результата в JSON формате
type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})

	ErrorBadRequest(w http.ResponseWriter, err error)
}

type Respond struct {
}

func NewResponder() Responder {
	return &Respond{}
}

// OutputJSON метод выдаёт успешный результат выполнения в формате JSON
func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("responder json encode error", err)
	}
}

// ErrorBadRequest метод выдаёт 400 ошибку и причину ошибки в формате JSON
func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	log.Println("http response bad request status code", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		log.Println("response writer error on write", err)
	}
}
