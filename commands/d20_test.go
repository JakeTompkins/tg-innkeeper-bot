package commands

import (
	"fmt"
	"testing"
	"tg-group-scheduler/tokenizer"
)

func TestRollD20(t *testing.T) {
	input := tokenizer.Token{Type: tokenizer.COMMAND, Value: "roll 3d4+2 4d20-1 1d100"}
	output, err := ExecuteRollD20(input)

	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(output)
	}
}

func TestExecuteOne(t *testing.T) {
	input := "2d6+5"

	for i := 0; i < 100; i++ {
		_, total, err := executeOne(input)
		if err != nil {
			t.Fatal(err)
		} else {
			if total < 7 {
				t.Fatalf("Got %d, expected no less than 7", total)
			}
			if total > 17 {
				t.Fatalf("Got %d, expected no more than 27", total)
			}
		}
	}
}
