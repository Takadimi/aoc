package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Takadimi/aoc/2022/file"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	strategyGuideEntries, err := file.Lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part One:", partOne(strategyGuideEntries))
	fmt.Println("Part Two:", partTwo(strategyGuideEntries))
}

type Choice int

const (
	Choice_Unknown Choice = iota
	Choice_Rock
	Choice_Paper
	Choice_Scissors
)

type Outcome int

const (
	Outcome_Unknown = -1
	Outcome_Lose    = 0
	Outcome_Draw    = 3
	Outcome_Win     = 6
)

func determineOutcome(opponentsMove, myMove Choice) Outcome {
	if opponentsMove == myMove {
		return Outcome_Draw
	}

	switch opponentsMove {
	case Choice_Rock:
		if myMove == Choice_Paper {
			return Outcome_Win
		}
		if myMove == Choice_Scissors {
			return Outcome_Lose
		}
	case Choice_Paper:
		if myMove == Choice_Scissors {
			return Outcome_Win
		}
		if myMove == Choice_Rock {
			return Outcome_Lose
		}
	case Choice_Scissors:
		if myMove == Choice_Rock {
			return Outcome_Win
		}
		if myMove == Choice_Paper {
			return Outcome_Lose
		}
	}

	return Outcome_Unknown
}

func partOne(entries []string) int {
	strategyGuide := parseStrategyGuidePartOne(entries)

	cumulativeScore := 0
	for _, round := range strategyGuide {
		opponentsMove := round[0]
		myMove := round[1]

		cumulativeScore += int(myMove)

		outcome := determineOutcome(opponentsMove, myMove)
		if outcome == Outcome_Unknown {
			panic("unknown outcome for round")
		}
		cumulativeScore += int(outcome)
	}

	return cumulativeScore
}

func parseStrategyGuidePartOne(entries []string) [][2]Choice {
	var choiceByGuide = map[string]Choice{
		"A": Choice_Rock,
		"B": Choice_Paper,
		"C": Choice_Scissors,

		"X": Choice_Rock,
		"Y": Choice_Paper,
		"Z": Choice_Scissors,
	}

	strategyGuide := [][2]Choice{}
	for _, line := range entries {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			panic("expected two moves on each strategy guide line")
		}
		firstMove, isValidChoice := choiceByGuide[fields[0]]
		if !isValidChoice {
			panic("expected valid choice")
		}
		secondMove, isValidChoice := choiceByGuide[fields[1]]
		if !isValidChoice {
			panic("expected valid choice")
		}

		strategyGuide = append(strategyGuide, [2]Choice{firstMove, secondMove})
	}

	return strategyGuide
}

type Round struct {
	OpponentsMove   Choice
	IntendedOutcome Outcome
}

func partTwo(entries []string) int {
	strategyGuide := parseStrategyGuidePartTwo(entries)

	winningMove := map[Choice]Choice{
		Choice_Rock:     Choice_Paper,
		Choice_Paper:    Choice_Scissors,
		Choice_Scissors: Choice_Rock,
	}

	losingMove := map[Choice]Choice{
		Choice_Rock:     Choice_Scissors,
		Choice_Paper:    Choice_Rock,
		Choice_Scissors: Choice_Paper,
	}

	cumulativeScore := 0
	for _, round := range strategyGuide {
		opponentsMove := round.OpponentsMove

		var myMove Choice
		switch round.IntendedOutcome {
		case Outcome_Lose:
			myMove = losingMove[opponentsMove]
		case Outcome_Draw:
			myMove = opponentsMove
		case Outcome_Win:
			myMove = winningMove[opponentsMove]
		}

		if myMove == Choice_Unknown {
			panic("unknown choice for round")
		}

		cumulativeScore += int(myMove)
		cumulativeScore += int(round.IntendedOutcome)
	}

	return cumulativeScore
}

func parseStrategyGuidePartTwo(entries []string) []Round {
	var choiceByGuide = map[string]Choice{
		"A": Choice_Rock,
		"B": Choice_Paper,
		"C": Choice_Scissors,
	}

	var outcomeByGuide = map[string]Outcome{
		"X": Outcome_Lose,
		"Y": Outcome_Draw,
		"Z": Outcome_Win,
	}

	strategyGuide := []Round{}
	for _, line := range entries {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			panic("expected two moves on each strategy guide line")
		}
		firstMove, isValidChoice := choiceByGuide[fields[0]]
		if !isValidChoice {
			panic("expected valid choice")
		}
		intendedOutcome, isValidOutcome := outcomeByGuide[fields[1]]
		if !isValidOutcome {
			panic("expected valid outcome")
		}

		strategyGuide = append(strategyGuide, Round{firstMove, intendedOutcome})
	}

	return strategyGuide
}
