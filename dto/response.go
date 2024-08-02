package dto

type Response struct {
	Data      any    `json:"data,omitempty"`
	Message   string `json:"message,omitempty"`
	PageCount int    `json:"page_count,omitempty"`
	ItemCount int    `json:"item_count,omitempty"`
}
