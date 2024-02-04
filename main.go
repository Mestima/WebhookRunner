package main

import (
	"WebhookRunner/handlers"
	"fmt"
)

func main() {
	handler := handlers.WebHandler{}
	handler.SetPort(8081)
	err := handler.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
