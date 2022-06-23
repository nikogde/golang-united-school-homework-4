package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	nums_re := regexp.MustCompile("[0-9]+") 
	nums := nums_re.FindAllString(input, -1)

	str_re := regexp.MustCompile("[a-zA-Z]+") 
	str := str_re.FindAllString(input, -1)

	if len(str) != 0 {
		_, err := strconv.Atoi(str[0])
		if err != nil {
			return "", fmt.Errorf(
				"Error (you must use only int values): %w", err,
			)
		}
	}
	if input == "" {
		return "", fmt.Errorf(
			"Error (must be at least two operands): %w", errorEmptyInput,
		)
	}
	if len(nums) != 2 {
		return "", fmt.Errorf(
			"Error (must be two operands): %w", errorNotTwoOperands,
		)
	}
	
	values := strings.Split(input, "+")

	if len(values) == 1 {
		last_index := strings.LastIndex(input, "-")
		return compute_result([]string{input[:last_index], input[last_index:]})
	}

	return compute_result([]string{values[0], values[1]})
}

func compute_result(operands []string) (string, error) {

	var result int

	for _, j := range operands {
		num, err := strconv.Atoi(j)
		if err != nil {
			return "", fmt.Errorf("Error: %w", err)
		}
		result += num
	}

	return strconv.Itoa(result), nil
}
