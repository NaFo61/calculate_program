package calculation

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/NaFo61/calculate_program/custom_errors"
)

func precedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func isValidExpression(expression string) bool {
	for _, r := range expression {
		if !(unicode.IsDigit(r) || r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')') {
			return false
		}
	}
	return true
}

func Calc(expression string) (float64, error) {
	if !isValidExpression(expression) {
		return 0, custom_errors.ErrInvalidExpression
	}

	var output []string
	stack := []byte{}

	var token string
	for i := 0; i < len(expression); i++ {
		r := expression[i]
		if unicode.IsDigit(rune(r)) {
			token += string(r)
		} else {
			if token != "" {
				output = append(output, token)
				token = ""
			}
			if r == '(' {
				stack = append(stack, r)
			} else if r == ')' {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					output = append(output, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					return 0, custom_errors.ErrInvalidExpression
				}
				stack = stack[:len(stack)-1]
			} else if precedence(r) > 0 { // Обработка операторов
				for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(r) {
					output = append(output, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, r)
			}
		}
		// важное изменение: добавление токена в конце цикла
		if i == len(expression)-1 && token != "" {
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	result, err := OPZ_to_result(strings.Join(output, " "))
	if err != nil {
		return 0, err // Передаем ошибку дальше
	}
	return result, nil
}

func OPZ_to_result(expression string) (float64, error) {
	stack := []float64{}
	tokens := strings.Fields(expression)
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, custom_errors.ErrInvalidExpression
			}
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, custom_errors.ErrDivisionByZero
				}
				stack = append(stack, a/b)
			}
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, custom_errors.ErrInvalidExpression
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, custom_errors.ErrInvalidExpression
	}
	return stack[0], nil
}
