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
		Name         string    `json:"name" bson:"name"`
		Urn          string    `json:"urn" bson:"urn"`
		Industry     string    `json:"industry" bson:"industry"`
		Emails       []string  `json:"emails" bson:"emails"`
		PhoneNumbers []string  `json:"phone_numbers" bson:"phone_numbers"`
		Address      Address   `json:"address" bson:"address"`
		CreatedAt    time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	}

	Address struct {
		Line1    string `json:"line_1" bson:"line_1"`
		Line2    string `json:"line_2" bson:"line_2"`
		City     string `json:"city" bson:"city"`
		Postcode string `json:"postcode" bson:"postcode"`
		Country  string `json:"country" bson:"country"`
	}
)
