package advent2023

import (
	"adventOfCode/utils"
	"regexp"
	"fmt"
)

func Day02_1(inputFilePath string) {
	var inputs []string = utils.ReadFile(inputFilePath);
	
	var accumulator int = 0;
	for _, game := range inputs {
		if verifyValidGame(game) {
			var gameId int = findGameId(game);

			accumulator += gameId;
		}
	}

	fmt.Println(accumulator);
}

func verifyValidGame(game string) bool {
	var validCubeMap map[string]int = initValidCubeMap();

	return !arrayContainsValueHigherThan(findRedCubes(game), validCubeMap["red"]) &&
		!arrayContainsValueHigherThan(findGreenCubes(game), validCubeMap["green"]) &&
		!arrayContainsValueHigherThan(findBlueCubes(game), validCubeMap["blue"]);
}

func initValidCubeMap() map[string]int {
	var result map[string]int = make(map[string]int);

	result["red"] = 12;
	result["green"] = 13;
	result["blue"] = 14;

	return result;
}

func findRedCubes(searchString string) []int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) red`);

	return utils.ExtractAllIntFromString(re, searchString);
}

func findGreenCubes(searchString string) []int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) green`);

	return utils.ExtractAllIntFromString(re, searchString);
}

func findBlueCubes(searchString string) []int {
	var re *regexp.Regexp = regexp.MustCompile(`([0-9]+) blue`);

	return utils.ExtractAllIntFromString(re, searchString);
}

func arrayContainsValueHigherThan(array []int, searchNumber int) bool {
	for _, value := range array {
		if value > searchNumber {
			return true;
		}
	}
	
	return false;
}

func findGameId(searchString string) int {
	var re *regexp.Regexp = regexp.MustCompile(`Game ([0-9]+):*`);

	return utils.ExtractIntFromString(re, searchString);
}
