package main

import "encoding/json"

func UnmarshalRule(data []byte) (Rule, error) {
	var r Rule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Rule) MarshalRule() ([]byte, error) {
	return json.Marshal(r)
}

type Rule struct {
	Device     string    `json:"device"`
	ActionID   string    `json:"action_id"`
	ActionType string    `json:"action_type"`
	Rule       RuleClass `json:"rule"`
}

type RuleClass struct {
	Parameter string      `json:"parameter"`
	Condition string      `json:"condition"`
	Value     interface{} `json:"value"`
}
