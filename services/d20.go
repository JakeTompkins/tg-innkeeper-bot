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
		result = fmt.Sprintf("%d", r)
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

func (d *diceRoll) Roll() int {
	n := max(d.NumberOfDice, 1)
	s := d.Sides
	f := d.FlatAdjustment

	var result int
	for i := 0; i < n; i++ {
		result += rand.Intn(s) + 1 + f
	}

	return result
}

func rollDice(diceStrings []string) (int, error) {
	var result int
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
			return 0, errors.New("Invalid dice notation")
		}

		result += newRoll.Roll()
	}

	return result, nil
}
