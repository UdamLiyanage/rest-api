package main

import "encoding/json"

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) MarshalUser() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	AccountType string      `json:"accountType"`
	Email       string      `json:"email"`
	Address     UserAddress `json:"address"`
	PhoneNumber string      `json:"phoneNumber"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type UserAddress struct {
	Primary   Billing `json:"primary"`
	Secondary Billing `json:"secondary"`
	Billing   Billing `json:"billing"`
}

type Billing struct {
	Line1    string `json:"line_1"`
	Line2    string `json:"line_2"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}
