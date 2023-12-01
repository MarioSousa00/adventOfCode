package advent2023

import (
	"adventOfCode/utils"
	"strconv"
	"errors"
	"fmt"
)

func Day01_1(inputFilePath string) {
	var inputs []string = utils.ReadFile(inputFilePath);

	var calibrationSum int = 0;

	for _, line := range inputs {
		firstDigit, err := findFirstDigit(line);
		utils.CheckError(err);
		lastDigit, err := findLastDigit(line);
		utils.CheckError(err);

		calibration, err := strconv.Atoi(string(firstDigit) + string(lastDigit));
		utils.CheckError(err);

		calibrationSum += calibration;
	}

	fmt.Println(calibrationSum);
}

func findFirstDigit(charSequence string) (rune, error) {
	runes := []rune(charSequence);

	for _, char := range runes {
		if char >= 48 && char <= 57 {
			return char, nil;
		}
	}

	return -1, errors.New("no digits found");
}

func findLastDigit(charSequence string) (rune, error) {
	runes := []rune(charSequence);

	for i := len(runes); i > 0; i-- {
		var char rune = runes[i - 1];
		if char >= 48 && char <= 57 {
			return char, nil;
		}
	}

	return -1, errors.New("no digits found");
}
