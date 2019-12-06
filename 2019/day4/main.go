package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(findNumPossiblePasswords(372304, 847060))
	fmt.Println(" Possible Passwords that meet criteria ")
}

func findNumPossiblePasswords(startNumber, endNumber int) int {
	numOfPossiblePasswords := 0

	for currPossiblePasswordNum := startNumber; currPossiblePasswordNum <= endNumber; currPossiblePasswordNum++ {
		if determineIfNumberIsPassword(currPossiblePasswordNum) == true {
			numOfPossiblePasswords++
		}
	}

	return numOfPossiblePasswords
}

func determineIfNumberIsPassword(potentialPassword int) bool {
	potentialPasswordString := strconv.Itoa(potentialPassword)

	hasAdjacentDigits := atLeastStrictlyTwoAdjacentDigitsIdentical(potentialPasswordString)
	nextDigitsNeverDecreased := digitsNeverDecreased(potentialPasswordString)

	return nextDigitsNeverDecreased && hasAdjacentDigits
}

func digitsNeverDecreased(numberString string) bool {
	currMaxDigitValue, _ := strconv.Atoi(numberString[0:1])
	for ind := 1; ind < len(numberString); ind++ {
		currDigit, _ := strconv.Atoi(numberString[ind-1 : ind])
		nextDigit, _ := strconv.Atoi(numberString[ind : ind+1])
		if currMaxDigitValue > currDigit || currDigit > nextDigit {
			return false
		}
		currMaxDigitValue = currDigit
	}
	return true
}

func atLeastTwoAdjacentDigitsIdentical(numberString string) bool {
	foundTwoIdenticalAdjacentDigits := false

	for ind := 1; ind < len(numberString); ind++ {
		currDigit := numberString[ind-1 : ind]
		nextDigit := numberString[ind : ind+1]
		if foundTwoIdenticalAdjacentDigits = currDigit == nextDigit; foundTwoIdenticalAdjacentDigits == true {
			return true
		}
	}
	return foundTwoIdenticalAdjacentDigits
}

// For Part 2 of adventofcode 12/04/2019
func atLeastStrictlyTwoAdjacentDigitsIdentical(numberString string) bool {
	foundTwoIdenticalAdjacentDigits := false

	numberOfDigits := len(numberString)
	for ind := 1; ind < numberOfDigits; ind++ {
		currDigit := numberString[ind-1 : ind]
		nextDigit := numberString[ind : ind+1]

		digitAfterNextIsNotIdentical := true
		hasDigitAfterNext := ind+2 <= numberOfDigits
		if hasDigitAfterNext {
			digitAfterNextDigit := numberString[ind+1 : ind+2]
			digitAfterNextIsNotIdentical = nextDigit != digitAfterNextDigit
		}

		if foundTwoIdenticalAdjacentDigits = ((currDigit == nextDigit) && digitAfterNextIsNotIdentical); foundTwoIdenticalAdjacentDigits == true {
			return true
		}

		if (currDigit == nextDigit) && !digitAfterNextIsNotIdentical {
			ind = ind + 2
			for ind < numberOfDigits && (numberString[ind-1:ind] == numberString[ind:ind+1]) {
				ind++
			}
			continue
		}
	}
	return foundTwoIdenticalAdjacentDigits
}
