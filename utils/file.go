package utils

import (
	"os"
	"bufio"
)

func ReadFile(fileName string) []string {
	directory, err := os.Getwd();
	CheckError(err);

	file, err := os.Open(directory + "/" + fileName);
	CheckError(err);
	defer file.Close();

	var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

	return lines;
}
