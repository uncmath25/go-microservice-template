package networking

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func EncodeResponse(resWriter http.ResponseWriter, statusCode int, data interface{}) error {
	resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := responseData{
		Status: getHttpStatusCode(statusCode),
		Data:   data,
	}

	return json.NewEncoder(resWriter).Encode(body)
}

func EncodeErrorResponse(resWriter http.ResponseWriter, statusCode int, err error) {
	errString := "error is nil"
	if err != nil {
		errString = err.Error()
	}

	resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := responseData{
		Status:  getHttpStatusCode(statusCode),
		Message: errString,
	}

	json.NewEncoder(resWriter).Encode(body)
}

func getHttpStatusCode(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "success"
	case code >= 300 && code < 400:
		return "redirect"
	case code >= 400 && code < 500:
		return "client error"
	case code >= 500 && code < 600:
		return "server error"
	default:
		return "unknown"
	}
}
