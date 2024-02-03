package handlers

import (
	"WebhookRunner/types"
	"errors"
	"os/exec"
)

type CommandHandler struct {
	command     *exec.Cmd
	commandRaw  string
	handlerType types.CommandHandler
	lastOutput  string
}

func (handler *CommandHandler) SetType(handlerType types.CommandHandler) {
	handler.handlerType = handlerType
}

func (handler *CommandHandler) SetCommand(command string) {
	handler.commandRaw = command
	handler.command = exec.Command(string(handler.handlerType), handler.commandRaw)
}

func (handler *CommandHandler) Execute() (string, error) {
	if handler.command != nil {
		tmpOut, err := handler.command.Output()
		if err != nil {
			return "", err
		}

		handler.lastOutput = string(tmpOut)
		handler.command = nil
		handler.commandRaw = ""
	} else {
		return "", errors.New("can't execute: command was not set")
	}

	return handler.lastOutput, nil
}

func (handler *CommandHandler) GetLastOutput() string {
	return handler.lastOutput
}

func (handler *CommandHandler) GetRawCommand() string {
	return handler.commandRaw
}

func (handler *CommandHandler) GetHandlerType() types.CommandHandler {
	return handler.handlerType
}
