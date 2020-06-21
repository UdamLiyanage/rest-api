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
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	AccountType string `json:"accountType"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
