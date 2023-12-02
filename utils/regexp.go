package utils

import (
	"regexp"
	"strconv"
)

func ExtractAllIntFromString(re *regexp.Regexp, searchString string) []int {
	var regexpResult [][][]byte = re.FindAllSubmatch([]byte(searchString), -1);

	var results []int;
	for _, match := range regexpResult {
		numberOfCubes, err := strconv.Atoi(string(match[1]));
		CheckError(err);

		results = append(results, numberOfCubes);
	}

	return results;
}

func ExtractIntFromString(re *regexp.Regexp, searchString string) int {
	var result [][]byte = re.FindSubmatch([]byte(searchString));

	integer, err := strconv.Atoi(string(result[1]))
	CheckError(err);

	return integer;
}
