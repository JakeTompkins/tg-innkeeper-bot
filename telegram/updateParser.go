package telegram

import (
	"regexp"
	"strings"
)

var VALID_COMMANDS = []string{"/roll"}

type updateParser struct{}

func newUpdateParser() *updateParser {
	var parser updateParser

	return &parser
}

func (p *updateParser) Parse(update update) {
	message := update.Message
	text := message.Text
	re := regexp.MustCompile(`(?P<command>^\\\w+)(?P<args>[\w\d\s]+)?$`)

	var command string
	var args []string

	match := re.FindStringSubmatch(text)

	for i, name := range re.SubexpNames() {
			if name == "command" {
				command = match[i]
			} else if name == "args" {
				args = strings.Split(strings.Trim(match[i], " "), " ")
			}
	}


  case "\\roll":
	switch command {
	
  }
}
