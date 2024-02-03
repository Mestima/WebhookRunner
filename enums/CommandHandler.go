package enums

import "WebhookRunner/types"

const (
	LinuxMinimal types.CommandHandler = "/bin/sh"
	Linux        types.CommandHandler = "/bin/bash"
	Windows      types.CommandHandler = "powershell"
)
