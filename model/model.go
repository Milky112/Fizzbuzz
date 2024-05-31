package model

import "time"

type FizzBuzzRequest struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type FizzbuzzResult struct {
	ProcessTime time.Duration `json:"process_time"`
	Result      []string      `json:"result"`
	Reason      string        `json:"reason"`
}
