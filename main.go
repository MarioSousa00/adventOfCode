package main

import (
	"os"
	"adventOfCode/advent2023"
)

func main() {
	var year string = os.Args[1];
	var day string = os.Args[2];
	var challenge string = os.Args[3];
	var inputFilePath string = os.Args[4];

	executeChallenge(year, day, challenge, inputFilePath);
}

func executeChallenge(year string, day string, challenge string, inputFilePath string) {
	var challengesMap map[string]func(string) = initChallengesMap();

	var challengeKey = year + "_" + day + "_" + challenge;
	challengesMap[challengeKey](inputFilePath);
}

func initChallengesMap() map[string]func(string) {
	var result map[string]func(string) = make(map[string]func(string));

	result["2023_01_01"] = advent2023.Day01_1;
	result["2023_01_02"] = advent2023.Day01_2;
	result["2023_02_01"] = advent2023.Day02_1;
	result["2023_02_02"] = advent2023.Day02_2;

	return result;
}
