package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"source/cmd/packages/console"
	"source/cmd/source/solutions"
	"strconv"
	"time"
)

var ARGUMENTS = [...]string{
	"--sample", "--input",
	"--parta", "--partb",
	"--solution01",
	"--solution02", "--solution03",
	"--solution04", "--solution05",
	"--solution06", "--solution07",
	"--solution08", "--solution09",
	"--solution10", "--solution11",
	"--solution12", "--solution13",
	"--solution14", "--solution15",
	"--solution16", "--solution17",
	"--solution18", "--solution19",
	"--solution20", "--solution21",
	"--solution22", "--solution23",
	"--solution24", "--solution25",
}

var INPUT_PATHS = [...]string{
	"Input01.input",
	"Input02.input", "Input03.input",
	"Input04.input", "Input05.input",
	"Input06.input", "Input07.input",
	"Input08.input", "Input09.input",
	"Input10.input", "Input11.input",
	"Input12.input", "Input13.input",
	"Input14.input", "Input15.input",
	"Input16.input", "Input17.input",
	"Input18.input", "Input19.input",
	"Input20.input", "Input21.input",
	"Input22.input", "Input23.input",
	"Input24.input", "Input25.input",
}

var SAMPLE_PATHS = [...]string{
	"Sample01.input",
	"Sample02.input", "Sample03.input",
	"Sample04.input", "Sample05.input",
	"Sample06.input", "Sample07.input",
	"Sample08.input", "Sample09.input",
	"Sample10.input", "Sample11.input",
	"Sample12.input", "Sample13.input",
	"Sample14.input", "Sample15.input",
	"Sample16.input", "Sample17.input",
	"Sample18.input", "Sample19.input",
	"Sample20.input", "Sample21.input",
	"Sample22.input", "Sample23.input",
	"Sample24.input", "Sample25.input",
}

var SOLUTIONS_A = [...]func(string){
	solutions.PrintSolution01A,
	solutions.PrintSolution02A,
	solutions.PrintSolution03A,
	solutions.PrintSolution04A,
	solutions.PrintSolution05A,
	solutions.PrintSolution06A,
	solutions.PrintSolution07A,
	solutions.PrintSolution08A,
	solutions.PrintSolution09A,
	solutions.PrintSolution10A,
	solutions.PrintSolution11A,
	solutions.PrintSolution12A,
	solutions.PrintSolution13A,
	solutions.PrintSolution14A,
	solutions.PrintSolution15A,
	solutions.PrintSolution16A,
	solutions.PrintSolution17A,
	solutions.PrintSolution18A,
	solutions.PrintSolution19A,
	solutions.PrintSolution20A,
	solutions.PrintSolution21A,
	solutions.PrintSolution22A,
	solutions.PrintSolution23A,
	solutions.PrintSolution24A,
	solutions.PrintSolution25A,
}

var SOLUTIONS_B = [...]func(string){
	solutions.PrintSolution01B,
	solutions.PrintSolution02B,
	solutions.PrintSolution03B,
	solutions.PrintSolution04B,
	solutions.PrintSolution05B,
	solutions.PrintSolution06B,
	solutions.PrintSolution07B,
	solutions.PrintSolution08B,
	solutions.PrintSolution09B,
	solutions.PrintSolution10B,
	solutions.PrintSolution11B,
	solutions.PrintSolution12B,
	solutions.PrintSolution13B,
	solutions.PrintSolution14B,
	solutions.PrintSolution15B,
	solutions.PrintSolution16B,
	solutions.PrintSolution17B,
	solutions.PrintSolution18B,
	solutions.PrintSolution19B,
	solutions.PrintSolution20B,
	solutions.PrintSolution21B,
	solutions.PrintSolution22B,
	solutions.PrintSolution23B,
	solutions.PrintSolution24B,
	solutions.PrintSolution25B,
}

type Arguments struct {
	sample          bool
	input           bool
	parta           bool
	partb           bool
	solutionsInt    []int
	solutionsString []string
}

var TITLE string = "Advent of Code"
var STAR string = console.YELLOW_TEXT + " * " + console.RESET_COLOR

// Main Function.
func AdventOfCode() {
	PrintTitle()                // Print Title.
	arguments := ReadArgs()     // Get Arguments.
	PrintArguments(arguments)   // Print Arguments.
	ExecuteSolutions(arguments) // Execute Solutions.
}

func PrintTitle() {
	print(STAR)
	color := 0
	for _, character := range TITLE {
		if color == 0 {
			fmt.Printf("%s%c%s", console.GREEN_TEXT, character, console.RESET_COLOR)
			color = 1
		} else {
			fmt.Printf("%s%c%s", console.RED_TEXT, character, console.RESET_COLOR)
			color = 0
		}
	}
	print(STAR + console.YELLOW_TEXT + "2022\n\n" + console.RESET_COLOR)
}

func ReadArgs() Arguments {
	var result Arguments
	arguments := os.Args
	if len(arguments) < 2 {
		PrintArgumentError()
		os.Exit(1)
	}
	if len(arguments) >= 2 {
		arguments = arguments[1:]
		for _, argument := range arguments {
			switch argument {
			case "--input":
				result.input = true
			case "--sample":
				result.sample = true
			case "--parta":
				result.parta = true
			case "--partb":
				result.partb = true
			default:
				solution, err := strconv.Atoi(argument[len(argument)-2:])
				if err != nil {
					log.Fatal(err)
				}
				result.solutionsInt = append(result.solutionsInt, solution)
				result.solutionsString = append(result.solutionsString, argument)
			}
		}
		if !result.parta && !result.partb {
			result.parta = true
			result.partb = true
		}

		if !result.sample && !result.input {
			result.sample = true
			result.input = true
		}
		slices.Sort(result.solutionsInt)
		slices.Sort(result.solutionsString)
	}
	return result
}

func PrintArguments(arguments Arguments) {
	print(STAR + "Command line arguments: \n\n")
	print(console.GREEN_TEXT)
	if arguments.input {
		println("\t--input")
	}
	if arguments.sample {
		println("\t--sample")
	}
	if arguments.parta {
		println("\t--parta")
	}
	if arguments.partb {
		println("\t--partb")
	}
	for _, solution := range arguments.solutionsString {
		println("\t" + solution)
	}
	println(console.RESET_COLOR)
}

func ExecuteSolutions(arguments Arguments) {
	for _, day := range arguments.solutionsInt {
		for index := 0; index <= 1; index++ {
			if index == 0 && !arguments.parta {
				continue
			}
			if index == 1 && !arguments.partb {
				continue
			}
			print(STAR + "Running solution " + console.CYAN_TEXT)
			print(day)
			print(console.RESET_COLOR + ", part " + console.CYAN_TEXT)
			var solver func(string)
			if index == 0 {
				solver = SOLUTIONS_A[day-1]
				print("A")
			}
			if index == 1 {
				print("B")
				solver = SOLUTIONS_B[day-1]
			}
			print(console.RESET_COLOR)

			var inputpaths []string
			directory, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			if arguments.sample {
				inputpaths = append(inputpaths, directory+"\\source\\samples\\"+SAMPLE_PATHS[day-1])
			}
			if arguments.input {
				inputpaths = append(inputpaths, directory+"\\source\\inputs\\"+INPUT_PATHS[day-1])
			}

			for _, path := range inputpaths {
				print("\n\n\tUsing input file: " + console.YELLOW_TEXT + path + console.RESET_COLOR + "\n\n")
				start := time.Now().UnixMicro()
				solver(path)
				end := time.Now().UnixMicro()
				duration := end - start
				print("\n\tSolution executed in: " + console.YELLOW_TEXT)
				print(duration)
				print(" us" + console.RESET_COLOR + "\n\n")
			}
		}
	}
}

func PrintArgumentError() {
	print(console.RED_TEXT)
	print("No solutions specified for execution. Pass desired solutions as command-line arguments. Sample usage:\n")
	print(console.RESET_COLOR)

	print("\tAdventOfCode.exe ")
	print(console.YELLOW_TEXT)
	print("--solution04\n")
	print(console.RESET_COLOR)

	print("\tAdventOfCode.exe ")
	print(console.YELLOW_TEXT)
	print("--solution01 --solution02\n")
	print(console.RESET_COLOR)

	print("\tAdventOfCode.exe ")
	print(console.YELLOW_TEXT)
	print("--partb --input --solution02\n")
	print(console.RESET_COLOR)

	print("\tAdventOfCode.exe ")
	print(console.YELLOW_TEXT)
	print("--parta --sample --puzzle03\n")
	print(console.RESET_COLOR)
}
