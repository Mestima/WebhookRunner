package enums

import "WebhookRunner/types"

const (
	LinuxMinimal types.CommandHandlerType = "/bin/sh"
	Linux        types.CommandHandlerType = "/bin/bash"
	Windows      types.CommandHandlerType = "powershell"
)
