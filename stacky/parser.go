package stacky

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInstruction(toParse string) (instruction, error) {
	var val instruction

	// Ignore literal values and characters
	if num, err := strconv.ParseUint(toParse, 10, 16); err == nil {
		return instruction(num), nil
	} else if char := strings.Trim(toParse, "'"); len(char) == 1 {
		return instruction(char[0]), nil
	}

	switch toParse {
	case "Print":
		val = instPrint
	case "Add":
		val = instAdd
	case "Sub":
		val = instSub
	case "Mult":
		val = instMult
	case "Div":
		val = instDiv
	case "Literal":
		val = instLiteral
	case "DBGSTK":
		val = instDBGSTK
	default:
		return val, errors.New("Unknown instruction")
	}

	return val, nil
}

func isEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}

func isComment(line string) bool {
	isComment := regexp.MustCompile("^;")
	return isComment.MatchString(line)
}

func parse(toParse []string) (instructions, error) {
	var err error
	parsed := make(instructions, 0)

	for n, line := range toParse {
		if isEmptyLine(line) {
			continue
		} else if isComment(line) {
			continue
		}

		inst, err := getInstruction(line)
		if err != nil {
			return parsed, errors.New(fmt.Sprintln(err, line, "at line", n+1))
		}
		parsed = append(parsed, instruction(inst))
	}

	return parsed, err
}

// Read parses a stacky file and returns instructions that the VM can
// understand.
func Read(filename string) (instructions, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if scanner.Err() != nil {
		var instructions instructions
		return instructions, scanner.Err()
	}

	return parse(data)
}
