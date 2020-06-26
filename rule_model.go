package main

import (
	"encoding/json"
	"time"
)

func UnmarshalRule(data []byte) (Rule, error) {
	var r Rule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Rule) MarshalRule() ([]byte, error) {
	return json.Marshal(r)
}

type Rule struct {
	Urn       string    `json:"urn" bson:"urn"`
	Device    string    `json:"device" bson:"device"`
	Actions   []string  `json:"actions" bson:"actions"`
	Rule      RuleClass `json:"rule" bson:"rule"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type RuleClass struct {
	Parameter string      `json:"parameter" bson:"parameter"`
	Condition string      `json:"condition" bson:"condition"`
	Value     interface{} `json:"value" bson:"value"`
}
