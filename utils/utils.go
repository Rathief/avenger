package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

type APIError struct {
	StatusCode          int
	ResponseCode        string
	ResponseDescription string
}

// api errors
var (
	ErrDuplicateData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "001",
		ResponseDescription: "Error: Data already exist",
	}
	ErrInsertData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "002",
		ResponseDescription: "Error: Failed inserting data",
	}
	ErrReadData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "003",
		ResponseDescription: "Error: Failed fetching data",
	}
	ErrUpdateData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "004",
		ResponseDescription: "Error: Failed updating data",
	}
	ErrDeleteData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "005",
		ResponseDescription: "Error: Failed deleting data. Dependency?",
	}
)

func JsonResponseWriter(w http.ResponseWriter, code int, input any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(input)
}

func Logger(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		currentTime := time.Now()
		log := fmt.Sprintf("[%v] - HTTP Request sent to %v %v", currentTime.Format("2006/01/02 15:04:05"), r.Method, r.URL.Path)
		fmt.Println(log)
		next(w, r, p)
	}
}
