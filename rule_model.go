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
	Device    string    `json:"device"`
	Actions   []string  `json:"actions"`
	Rule      RuleClass `json:"rule"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RuleClass struct {
	Parameter string      `json:"parameter"`
	Condition string      `json:"condition"`
	Value     interface{} `json:"value"`
}
