package main

import (
	"encoding/json"
	"time"
)

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) MarshalUser() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	FirstName   string      `json:"first_name" bson:"first_name"`
	LastName    string      `json:"last_name" bson:"last_name"`
	AccountType string      `json:"account_type" bson:"account_type"`
	Email       string      `json:"email" bson:"email"`
	Address     UserAddress `json:"address" bson:"address"`
	PhoneNumber string      `json:"phone_number" bson:"phone_number"`
	CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}

type UserAddress struct {
	Primary   Billing `json:"primary" bson:"primary"`
	Secondary Billing `json:"secondary" bson:"secondary"`
	Billing   Billing `json:"billing" bson:"billing"`
}

type Billing struct {
	Line1    string `json:"line_1" bson:"line_1"`
	Line2    string `json:"line_2" bson:"line_2"`
	City     string `json:"city" bson:"city"`
	Postcode string `json:"postcode" bson:"postcode"`
	Country  string `json:"country" bson:"country"`
}
