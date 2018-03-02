package main

type Person struct {
	ID        int   `json:"id"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
}
type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}