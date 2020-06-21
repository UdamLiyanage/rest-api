package main

import (
	"encoding/json"
	"time"
)

func UnmarshalEnterprise(data []byte) (Enterprise, error) {
	var r Enterprise
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Enterprise) MarshalEnterprise() ([]byte, error) {
	return json.Marshal(r)
}

type (
	Enterprise struct {
		Name         string    `json:"name"`
		Urn          string    `json:"urn"`
		Industry     string    `json:"industry"`
		Emails       []string  `json:"emails"`
		PhoneNumbers []string  `json:"phone_numbers"`
		Address      Address   `json:"address"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	Address struct {
		Line1    string `json:"line_1"`
		Line2    string `json:"line_2"`
		City     string `json:"city"`
		Postcode string `json:"postcode"`
		Country  string `json:"country"`
	}
)
