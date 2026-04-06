package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get(*cfg.Next)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	if response.Next != nil {
		cfg.Next = response.Next
	}
	if response.Previous != nil {
		cfg.Previous = response.Previous
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	res, err := http.Get(*cfg.Previous)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	if response.Next != nil {
		cfg.Next = response.Next
	}
	if response.Previous != nil {
		cfg.Previous = response.Previous
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}
