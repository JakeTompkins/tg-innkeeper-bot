package services

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type D20Service struct {
	service Service
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func NewD20Service() *D20Service {
	return &D20Service{
		service: Service{
			Name:          "D20 Roller",
			ValidCommands: []string{"/roll"},
		},
	}
}

func (d *D20Service) CommandIsValid(command *ServiceCommand) bool {
	return d.service.CommandIsValid(command)
}

func (d *D20Service) Execute(command *ServiceCommand) (*ServiceResult, error) {
	var result string

	switch command.Command {
	case "/roll":
		r, err := rollDice(command.Args)
		if err != nil {
			return nil, err
		}
		result = r
	}

	return &ServiceResult{
		Code: SUCCESS,
		Data: result,
	}, nil
}

func (d *D20Service) HelpMessage() string {
	return `
   Valid Commands:

	 /roll <dice_notation>
	`
}

type diceRoll struct {
	NumberOfDice   int
	Sides          int
	FlatAdjustment int
}

func parseFlatAdjustmentString(s string) int {
	var multiplier int
	var adjustment int

	if strings.Contains("-", s) {
		multiplier = -1
	} else {
		multiplier = 1
	}

	r := regexp.MustCompile(`\d+$`)
	digits := r.FindString(s)

	adjustment, _ = strconv.Atoi(digits)

	return adjustment * multiplier
}

func (d *diceRoll) Roll() string {
	n := max(d.NumberOfDice, 1)
	s := d.Sides
	f := d.FlatAdjustment

	var result string
	var total int
	for i := 0; i < n; i++ {
		rollValue := rand.Intn(s) + 1
		adjustment := f
		total += rollValue + adjustment

		result += fmt.Sprintf("1d%d + %d", rollValue, adjustment)

		if n-i == 1 {
			result += fmt.Sprintf(" = %d", total)
		}
	}

	return result
}

func rollDice(diceStrings []string) (string, error) {
	var result string
	diceStringRegex := regexp.MustCompile(`^(?P<NumberOfDice>\d+)?d(?P<Sides>\d+)\s?(?P<FlatAdjustment>[+-]\s?\d+)?$`)

	for _, diceString := range diceStrings {
		newRoll := &diceRoll{}
		matches := diceStringRegex.FindStringSubmatch(diceString)

		for i, name := range diceStringRegex.SubexpNames() {
			value := matches[i]

			switch name {
			case "NumberOfDice":
				newRoll.NumberOfDice, _ = strconv.Atoi(value)
			case "Sides":
				newRoll.Sides, _ = strconv.Atoi(value)
			case "FlatAdjustment":
				newRoll.FlatAdjustment = parseFlatAdjustmentString(value)
			}
		}

		if newRoll.Sides <= 0 {
			return "0", errors.New("Invalid dice notation")
		}

		result += newRoll.Roll() + "\n"
	}

	return result, nil
}
