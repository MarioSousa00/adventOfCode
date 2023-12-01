package advent2023

import (
	"adventOfCode/utils"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

func Day01_2(inputFilePath string) {
	var numbersMapper map[string]int = initNumberMap();

	var inputs []string = utils.ReadFile(inputFilePath);

	var calibrationSum int = 0;

	for _, line := range inputs {
		firstDigit, err := findFirstNumber(numbersMapper, line);
		utils.CheckError(err);
		lastDigit, err := findLastNumber(numbersMapper, line);
		utils.CheckError(err);

		calibration, err := strconv.Atoi(string(firstDigit) + string(lastDigit));
		utils.CheckError(err);

		calibrationSum += calibration;
	}

	fmt.Println(calibrationSum);
}

func initNumberMap() map[string]int {
	var result map[string]int = make(map[string]int);

	result["one"] = 49;
	result["two"] = 50;
	result["three"] = 51;
	result["four"] = 52;
	result["five"] = 53;
	result["six"] = 54;
	result["seven"] = 55;
	result["eight"] = 56;
	result["nine"] = 57;

	return result;
}

func findFirstNumber(numbersMap map[string]int, searchString string) (int, error) {
	for i := 0; i < len(searchString); i++ {
		var char byte = searchString[i];
		if char >= 48 && char <= 57 {
			return int(char), nil;
		}

		var subString string = searchString[i:];

		for key, value := range numbersMap {
			if strings.HasPrefix(subString, key) {
				return value, nil;
			}
		}
	}

	return -1, errors.New("no digits found");
}

func findLastNumber(numbersMap map[string]int, searchString string) (int, error) {
	for i := len(searchString); i > 0; i-- {
		var char byte = searchString[i - 1]
		if char >= 48 && char <= 57 {
			return int(char), nil;
		}

		var subString string = searchString[:i];

		for key, value := range numbersMap {
			if strings.HasSuffix(subString, key) {
				return value, nil;
			}
		}
	}

	return -1, errors.New("no digits found");
}
