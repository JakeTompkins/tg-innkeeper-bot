package telegram

import (
	"regexp"
	"strings"

	s "tg-group-scheduler/services"
)

var VALID_COMMANDS = []string{"/roll"}

type updateParser struct{}

func newUpdateParser() *updateParser {
	var parser updateParser

	return &parser
}

func (p *updateParser) Parse(update update) s.ServiceCommand {
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
