package telegram

import (
	"regexp"
	"strings"

	s "tg-group-scheduler/services"
)

// TODO: This logic should be moved to a struct that has access to alls ervices in order to construct the VALID_COMMANDS var
var VALID_COMMANDS = []string{"/roll"}

func parseUpdateToCommand(update update) s.ServiceCommand {
	message := update.Message
	text := message.Text
	re := regexp.MustCompile(`(?P<command>^/\w+)(?P<args>[\w\d\s]+)?$`)

	var serviceCommand s.ServiceCommand

	match := re.FindStringSubmatch(text)

	for i, name := range re.SubexpNames() {
		if name == "command" {
			serviceCommand.Command = match[i]
		} else if name == "args" {
			serviceCommand.Args = strings.Split(strings.Trim(match[i], " "), " ")
		}
	}

	return serviceCommand
}
