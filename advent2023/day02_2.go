package advent2023

import (
	"adventOfCode/utils"
	"regexp"
	"fmt"
)

func Day02_2(inputFilePath string) {
	var inputs []string = utils.ReadFile(inputFilePath);
	
	var accumulator int = 0;
	for _, game := range inputs {
		var minimumRedCubes int = findMinimumRedCubes(game);
		var minimumGreenCubes int = findMinimumGreenCubes(game);
		var minimumBlueCubes int = findMinimumBlueCubes(game);

		accumulator += minimumRedCubes * minimumGreenCubes * minimumBlueCubes;
	}

	fmt.Println(accumulator);
}

func findMinimumRedCubes(game string) int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) red`);

	return findLowerCubeOccurrence(re, game);
}

func findMinimumGreenCubes(game string) int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) green`);

	return findLowerCubeOccurrence(re, game);
}

func findMinimumBlueCubes(game string) int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) blue`);

	return findLowerCubeOccurrence(re, game);
}

func findLowerCubeOccurrence(re *regexp.Regexp, game string) int {
	var cubeOccurrences []int = utils.ExtractAllIntFromString(re, game);

	return arrayHigherValue(cubeOccurrences);
}

func arrayHigherValue(array []int) int {
	var higherValue int = 0;

	for _, value := range array {
		if value > higherValue {
			higherValue = value;
		}
	}
	
	return higherValue;
}
