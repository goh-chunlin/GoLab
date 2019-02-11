package models

// APIStatus is a record of API action status
type APIStatus struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
