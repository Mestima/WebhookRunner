package handlers

import (
	"WebhookRunner/types"
	"errors"
	"os/exec"
)

type CommandHandler struct {
	command     *exec.Cmd
	commandRaw  string
	handlerType types.CommandHandlerType
	lastOutput  string
}

func (handler *CommandHandler) SetType(handlerType types.CommandHandlerType) {
	handler.handlerType = handlerType
}

func (handler *CommandHandler) SetCommand(command string, isFile bool) {
	handler.commandRaw = command
	if isFile {
		handler.command = exec.Command(string(handler.handlerType), handler.commandRaw)
	} else {
		handler.command = exec.Command(string(handler.handlerType), "-c", handler.commandRaw)
	}
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

func (handler *CommandHandler) GetHandlerType() types.CommandHandlerType {
	return handler.handlerType
}
