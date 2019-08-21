package jsonapi

import (
	"net/http"
	"time"
)

type Link struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type Data struct {
	ID        string      `json:"id"`
	Type      string      `json:"type,omitempty"`
	Attribute interface{} `json:"attributes,omitempty"`
}

type ResponseMultiData struct {
	Total    int64  `json:"total,omitempty"`
	Link     Link   `json:"links,omitempty"`
	Data     []Data `json:"data,omitempty"`
	Included []Data `json:"included,omitempty"`
}

type ResponseData struct {
	Data   Data `json:"data,omitempty"`
	Status int  `json:"status,omitempty"`
}

type Error struct {
	Status      string  `json:"status"`
	Message     string  `json:"error"`
	ProcessTime float64 `json:"server_process_time"`
}

type ResponseDataV4 struct {
	Status      string      `json:"status"`
	Config      *string     `json:"config"`
	ProcessTime float64     `json:"server_process_time"`
	Data        interface{} `json:"data,omitempty"`
}

// ResponseDataV5 make data success similar with error struct
type ResponseDataGeneral struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	ProcessTime float64     `json:"server_process_time"`
	Data        interface{} `json:"data,omitempty"`
}

type Module struct {
	start    time.Time
	response http.ResponseWriter
	callback string
}
