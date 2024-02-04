package handlers

import (
	"WebhookRunner/enums"
	"WebhookRunner/utils"
	"errors"
	"flag"
	"fmt"
	"net/http"
)

type WebHandler struct {
	commandHandler CommandHandler
	port           int
	isLogCommands  bool
	isLogOutput    bool
}

func (handler *WebHandler) SetPort(port int) {
	handler.port = port
}

func (handler *WebHandler) GetPort() int {
	return handler.port
}

func (handler *WebHandler) ListenAndServe() error {
	handler.init()

	if handler.port == 0 {
		return errors.New("port value is not set for WebHandler")
	}

	http.HandleFunc("/", handler.executeCode)
	fmt.Println("Server is listening on port " + fmt.Sprint(handler.port))
	return http.ListenAndServe(":"+fmt.Sprint(handler.port), nil)
}

func (handler *WebHandler) init() {
	handler.commandHandler = CommandHandler{}

	isMinimal := flag.Bool("minimal", false, "use /bin/sh instead of /bin/bash to execute scripts on Linux systems") // TODO move to a standalone handler
	isLogCommands := flag.Bool("log-commands", false, "log received commands")
	isLogOutput := flag.Bool("log-output", false, "log executions output")
	flag.Parse()

	handler.isLogCommands = *isLogCommands
	handler.isLogOutput = *isLogOutput

	if utils.IsWindows() {
		handler.commandHandler.SetType(enums.Windows)
	} else {
		if *isMinimal {
			handler.commandHandler.SetType(enums.LinuxMinimal)
		} else {
			handler.commandHandler.SetType(enums.Linux)
		}
	}
}

func (handler *WebHandler) executeCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		cmd := r.Form.Get("command")
		if cmd != "" {
			handler.commandHandler.SetCommand(cmd)
			out, err := handler.commandHandler.Execute()
			if err != nil {
				fmt.Fprint(w, fmt.Sprint(err))
				fmt.Println(err)
			} else {
				fmt.Fprint(w, out)

				if handler.isLogCommands {
					fmt.Println(cmd)
				}

				if handler.isLogOutput {
					fmt.Println(out)
				}
			}
		}
	}
}
