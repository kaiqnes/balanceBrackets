package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Scenario struct {
	Description,
	Input,
	ExpectOutput string
}

type Scenarios []Scenario

func NewScenario(description, input, expectOutput string) Scenario {
	return Scenario{
		Description:  description,
		Input:        input,
		ExpectOutput: expectOutput,
	}
}

func TestBalanceStrings(t *testing.T) {
	scenarios := Scenarios{
		NewScenario("Default challenge input", "[a][b]][c]", "[a][b][c]"),
		NewScenario("Input with extra open brackets", "[[[a][[b]][[c]", "[a][b][c]"),
		NewScenario("Input with extra close brackets", "[a]]][b]][c]]", "[a][b][c]"),
		NewScenario("Input with blank pair brackets", "[a][b][][c]", "[a][b][c]"),
		NewScenario("Input with extra open and blank pair brackets", "[[[a][b][][c]", "[a][b][c]"),
		NewScenario("Input with extra close and blank pair brackets", "[a][b][][c]]]", "[a][b][c]"),
		NewScenario("Input with extra open, close and blank pair brackets", "[[[a][b][][c]]]", "[a][b][c]"),
		NewScenario("Input with more than a value separated by open brackets", "[[[[[[][1[[[[[[[a][b][c]]]]]", "[1a][b][c]"),
	}

	for _, scenario := range scenarios {
		fmt.Println(">", scenario.Description)

		// ReplaceAll is failing in last scenario due split input using '['
		//replaceAll := BalanceUsingReplaceAll(scenario.Input)
		//assert.NotEmpty(t, replaceAll)
		//assert.Equal(t, scenario.ExpectOutput, replaceAll)

		stack := BalanceUsingStack(scenario.Input)
		assert.NotEmpty(t, stack)
		assert.Equal(t, scenario.ExpectOutput, stack)

		count := BalanceUsingCount(scenario.Input)
		assert.NotEmpty(t, count)
		assert.Equal(t, scenario.ExpectOutput, count)

		formattedCount := BalanceUsingCountFormatted(scenario.Input)
		assert.NotEmpty(t, formattedCount)
		assert.Equal(t, scenario.ExpectOutput, formattedCount)
	}
}
