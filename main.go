package main

import (
	"WebhookRunner/enums"
	"WebhookRunner/handlers"
	"fmt"
)

func main() {
	var cmd string
	fmt.Scan(&cmd)

	cmdHandler := handlers.CommandHandler{}
	cmdHandler.SetType(enums.Windows)
	cmdHandler.SetCommand(cmd)
	out, err := cmdHandler.Execute()
	fmt.Println(out, err)
}
