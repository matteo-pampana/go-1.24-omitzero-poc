package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	type Response struct {
		ID        string    `json:"id"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	resp := Response{
		ID: "123",
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("result 1: %s\n", data)
}

func example2() {
	type Response struct {
		ID        string     `json:"id"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	resp := Response{
		ID: "123",
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("result 2: %s\n", data)
}

func example3() {
	type Response struct {
		ID        string    `json:"id"`
		UpdatedAt time.Time `json:"updated_at,omitzero"`
	}

	resp := Response{
		ID: "123",
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("result 3: %s\n", data)
}

type CanBeZero struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c CanBeZero) IsZero() bool {
	return c.Key == "zero" && c.Value == "zero"
}

func example4() {
	type Response struct {
		ID        string    `json:"id"`
		Feature   CanBeZero `json:"feature,omitzero"`
		UpdatedAt time.Time `json:"updated_at,omitzero"`
	}

	resp := Response{
		ID: "123",
		Feature: CanBeZero{
			Key:   "zero",
			Value: "zero",
		},
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("result 4: %s\n", data)
}
