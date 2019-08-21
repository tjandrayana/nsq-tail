package jsonapi

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	"net/http"
)

func New(res http.ResponseWriter, call string) *Module {
	return &Module{
		start:    time.Now(),
		response: res,
		callback: call,
	}
}

func (m *Module) SuccessWriter(dataResponse interface{}) {
	dataBytes, err := json.Marshal(dataResponse)
	if err != nil {
		log.Fatalln(err)
	}

	if m.callback != "" {
		m.response.Header().Set("Content-Type", "application/javascript")
		m.response.WriteHeader(http.StatusOK)
		fmt.Fprintf(m.response, "%s(%s)", m.callback, dataBytes)
	} else {
		m.response.Header().Set("Content-Type", "application/json")
		m.response.WriteHeader(http.StatusOK)
		m.response.Write(dataBytes)
	}
}

func (m *Module) ErrorWriter(status int, message string) {
	m.response.Header().Set("Content-Type", "application/json")
	m.response.Header().Set("Status-Code", fmt.Sprintf("%d", status))
	m.response.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization")
	m.response.WriteHeader(status)

	errorsResponse := Error{
		Status:      http.StatusText(status),
		Message:     message,
		ProcessTime: time.Since(m.start).Seconds(),
	}
	errorsBytes, err := json.Marshal(errorsResponse)
	if err != nil {
		log.Fatalln(err)
	}

	m.response.Write(errorsBytes)
}

func (m *Module) Writer(dataResponse interface{}, status int) {
	dataBytes, err := json.Marshal(dataResponse)
	if err != nil {
		log.Fatalln(err)
	}

	if m.callback != "" {
		m.response.Header().Set("Content-Type", "application/javascript")
		m.response.WriteHeader(status)
		fmt.Fprintf(m.response, "%s(%s)", m.callback, dataBytes)
	} else {
		m.response.Header().Set("Content-Type", "application/json")
		m.response.WriteHeader(status)
		m.response.Write(dataBytes)
	}
}
