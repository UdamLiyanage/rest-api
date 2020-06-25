package main

import (
	"encoding/json"
	"time"
)

func UnmarshalAction(data []byte) (Action, error) {
	var r Action
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Action) MarshalAction() ([]byte, error) {
	return json.Marshal(r)
}

type Action struct {
	Schema        string      `json:"schema" bson:"schema"`
	Type          string      `json:"type" bson:"type"`
	Configuration interface{} `json:"configuration" bson:"configuration"`
	CreatedAt     time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at" bson:"updated_at"`
}
