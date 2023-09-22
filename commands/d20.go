package commands

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"tg-group-scheduler/tokenizer"
)

type ValidDie int

const (
	FOUR    ValidDie = 4
	SIX     ValidDie = 6
	EIGHT   ValidDie = 8
	TEN     ValidDie = 10
	TWELVE  ValidDie = 12
	TWENTY  ValidDie = 20
	PERCENT ValidDie = 100
)

type diceNotation struct {
	NumberOfDice  int
	NumberOfSides int
	Modifier      int
}

func (d *diceNotation) dieIsValid() bool {
	sides := d.NumberOfSides
	return sides == int(FOUR) || sides == int(SIX) || sides == int(EIGHT) || sides == int(TEN) || sides == int(TWELVE) || sides == int(TWENTY) || sides == int(PERCENT)
}

func (d *diceNotation) roll() (string, int) {
	var explanation string
	var modifierOperation string
	var total int

	if d.Modifier >= 0 {
		modifierOperation = "+"
	} else {
		modifierOperation = "-"
	}

	for i := 0; i < d.NumberOfDice; i++ {
		v := rand.Intn(d.NumberOfSides-1) + 1
		total += v
	}

	total += d.Modifier

	explanation = fmt.Sprintf("%dd%d%s%d = %d", d.NumberOfDice, d.NumberOfSides, modifierOperation, int(math.Abs(float64(d.Modifier))), total)

	return explanation, total
}

func newDiceNotation(dice int, sides int, modifier int) (*diceNotation, error) {
	if dice == 0 {
		dice = 1
	}

	notation := &diceNotation{
		NumberOfDice:  dice,
		NumberOfSides: sides,
		Modifier:      modifier,
	}

	if !notation.dieIsValid() {
		return nil, errors.New(fmt.Sprintf("Invalid dice notation: %dd%d + %d", dice, sides, modifier))
	}

	return notation, nil
}

func executeOne(stringNotation string) (string, int, error) {
	var explanation string
	var numberOfDice int
	var numberOfSides int
	var modifier int
	var err error
	modifierMultiplier := 1

	splitOnD := strings.Split(stringNotation, "d")

	if len(splitOnD) == 1 {
		numberOfDice = 1
		numberOfSides, err = strconv.Atoi(splitOnD[0])
	} else {
		numberOfDice, err = strconv.Atoi(splitOnD[0])
	}

	if err != nil {
		return explanation, 0, err
	}

	var modifierSep string

	if strings.Index(splitOnD[1], "+") != -1 {
		modifierSep = "+"
	} else if strings.Index(splitOnD[1], "-") != -1 {
		modifierMultiplier = -1
		modifierSep = "-"
	}

	if modifierSep != "" {
		splitOnMod := strings.Split(splitOnD[1], modifierSep)

		numberOfSides, err = strconv.Atoi(splitOnMod[0])
		modifier, err = strconv.Atoi(splitOnMod[1])
	} else {
		numberOfSides, err = strconv.Atoi(splitOnD[1])
	}

	if err != nil {
		return explanation, 0, err
	}

	diceNotation, err := newDiceNotation(numberOfDice, numberOfSides, modifier*modifierMultiplier)

	if err != nil {
		return explanation, 0, nil
	}

	exp, rollResult := diceNotation.roll()

	explanation += exp + "\n"

	return explanation, rollResult, nil
}

func ExecuteRollD20(token tokenizer.Token) (string, error) {
	var result string

	if token.Type != tokenizer.COMMAND {
		return result, errors.New("Invalid token passed to d20 command")
	}

	parts := strings.Split(token.Value, " ")

	if strings.ToLower(parts[0]) != "roll" {
		return result, errors.New("Non-d20 command passed to d20 executor")
	}

	for _, part := range parts[1:] {
		explanation, _, err := executeOne(part)

		if err != nil {
			return result, err
		}

		result += explanation + "\n"
	}

	return result, nil
}
