package handlers

import (
	"WebhookRunner/enums"
	"WebhookRunner/utils"
	"errors"
	"fmt"
	"net/http"
)

type WebHandler struct {
	commandHandler CommandHandler
	port           string
	logCommands    bool
	logOutput      bool
	isMinimal      bool
}

func (handler *WebHandler) SetPort(port string) {
	handler.port = port
}

func (handler *WebHandler) GetPort() string {
	return handler.port
}

func (handler *WebHandler) SetLogCommands(value bool) {
	handler.logCommands = value
}

func (handler *WebHandler) GetLogCommands() bool {
	return handler.logCommands
}

func (handler *WebHandler) SetLogOutput(value bool) {
	handler.logOutput = value
}

func (handler *WebHandler) GetLogOutput() bool {
	return handler.logOutput
}

func (handler *WebHandler) SetMinimal(value bool) {
	handler.isMinimal = value
}

func (handler *WebHandler) GetMinimal() bool {
	return handler.isMinimal
}

func (handler *WebHandler) ListenAndServe() error {
	handler.init()

	if handler.port == "" {
		return errors.New("port value is not set for WebHandler")
	}

	http.HandleFunc("/", handler.executeCode)
	fmt.Println("Server is listening on port " + handler.port)
	return http.ListenAndServe(":"+handler.port, nil)
}

func (handler *WebHandler) init() {
	handler.commandHandler = CommandHandler{}

	if utils.IsWindows() {
		handler.commandHandler.SetType(enums.Windows)
	} else {
		if handler.isMinimal {
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
		filename := r.Form.Get("filename")

		if cmd != "" || filename != "" {
			if cmd != "" {
				handler.commandHandler.SetCommand(cmd, false)
			} else {
				handler.commandHandler.SetCommand(filename, true)
			}

			out, err := handler.commandHandler.Execute()
			if err != nil {
				fmt.Fprint(w, fmt.Sprint(err))
				fmt.Println(err)
			} else {
				fmt.Fprint(w, out)

				if handler.logCommands {
					fmt.Println(cmd)
				}

				if handler.logOutput {
					fmt.Println(out)
				}
			}
		}
	}
}
