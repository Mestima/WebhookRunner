package enums

type CommandHandlerType string

const (
	LinuxMinimal CommandHandlerType = "/bin/sh"
	Linux        CommandHandlerType = "/bin/bash"
	Windows      CommandHandlerType = "powershell"
)
