package main

import (
	"WebhookRunner/handlers"
	"flag"
	"fmt"
)

func main() {
	isMinimal := flag.Bool("minimal", false, "use /bin/sh instead of /bin/bash to execute scripts on Linux systems")
	logCommands := flag.Bool("log-commands", false, "log received commands")
	logOutput := flag.Bool("log-output", false, "log executions output")
	port := flag.String("port", "8080", "server port")
	flag.Parse()

	handler := handlers.WebHandler{}

	handler.SetMinimal(*isMinimal)
	handler.SetLogCommands(*logCommands)
	handler.SetLogOutput(*logOutput)
	handler.SetPort(*port)

	err := handler.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
