# Advent of Code

## Description

This project saves the resolution of Advent of Code challenges. To see the challenges, please visit the following website: https://adventofcode.com/.

The software used to complete the challenges is developed in Go https://go.dev/.

The main goal of the project is not to give the solution of the challenges, but to share one of the possible resolution,s so that software developers who interact with the project have a space to discuss resolution ideas.

## Usage

### Dependencies

 * Go (version 1.21.4 recommended)

### Execution

To execute the resolution, follow the folling steps:

```bash
git clone https://github.com/MariosSousa00/adventOfCode
```
```bash
cd adventOfCode
```

On the following command, replace:
 * `<year>` with the year of the challenge you pretend to execute
 * `<day>` with the day of the challenge you pretend to execute
 * `<challenge>` with the number of the challenge you pretend to execute
 * `<inputFilePath>` with the path to the input file relative to the project root

```bash
go run main.go <year> <day> <challenge> <inputFilePath>
```

A valid example, to run the second challeng of the December 1st of 2023, is:
```bash
go run main.go 2023 01 02 inputs/2023/day01Input.txt
```

### Input Files

The repository already contains the input files given on the challenge on the `input` directory.
