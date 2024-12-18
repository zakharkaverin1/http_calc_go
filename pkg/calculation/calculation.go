package calculation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if regexp.MustCompile(`[a-zA-Zа-яА-Я]`).MatchString(expression) {
		return 0, ErrInvalidExpression
	}
	var posledovatelnost []string
	var err error
	counter3 := 0
	counter4 := 0
	// проверяем, чтобы количество знаков операций скобок не превышало количества чисел. Иначе выводим ошибку
	for i := 0; i < len(expression); i++ {
		if expression[i] == '0' || expression[i] == '1' || expression[i] == '2' || expression[i] == '3' || expression[i] == '4' || expression[i] == '5' || expression[i] == '6' || expression[i] == '7' || expression[i] == '8' || expression[i] == '9' {
			counter3 += 1
		} else if expression[i] == '*' || expression[i] == '/' || expression[i] == '-' || expression[i] == '+' {
			counter4 += 1
		}
	}
	if counter3 <= counter4 {
		return 0, ErrInvalidExpression
	}
	// проверяем, есть ли скобки в выражении. если да, то активируем специальные функции
	if strings.Contains(expression, "(") {
		counter1 := 0
		counter2 := 0
		// проверям, совпадает ли количество закрывающихся и открывающихся скобок. Иначе выводим ошибку
		for i := 0; i < len(expression); i++ {
			if expression[i] == '(' {
				counter1 += 1
			} else if expression[i] == ')' {
				counter2 += 1
			}
		}

		if counter1 != counter2 {
			return 0, ErrInvalidExpression
		}
		posledovatelnost = append(posledovatelnost, find_brackets(expression)...)
		new := calc_brackets(posledovatelnost)
		posledovatelnost = posledovatelnost[:0]
		for i := 0; i < len(new); i++ {
			posledovatelnost = append(posledovatelnost, new[i])
		}
		// находим все выражения в скобках, добавляем их в специальный список
		for i := 0; i < len(expression); i++ {
			if expression[i] == '(' {
				counter := i
				skobki := expression[i+1:]
				for j := 0; j < len(skobki); j++ {
					if skobki[j] != ')' {
						counter += 1
					} else if skobki[j] == ')' {
						counter += 2
						break
					}
				}
				// считаем выражение в скобках и заменяем их в основном выражении
				expression = strings.Replace(expression, expression[i:counter], posledovatelnost[0], 1)
				posledovatelnost = posledovatelnost[1:]
				counter = 0
			}
		}
	}
	// досчитываем
	expression, err = mult_div(expression)
	if err != nil {
		return 0, err
	}
	expression = add_sub(expression)
	ans, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0, ErrSomethingWentWrong
	}
	return ans, nil
}

// с помощью неё находим находим все выражения в скобках
func find_brackets(expression string) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var sk []string
	var stroka string
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			skobki := expression[i+1:]
			for j := 0; j < len(skobki); j++ {
				if skobki[j] != ')' {
					stroka = stroka + string(skobki[j])
				} else if skobki[j] == ')' {
					break
				}
			}
			sk = append(sk, stroka)
			stroka = stroka[:0]
		}
	}
	return sk
}

// рассчитываем все скобочные выраженея
func calc_brackets(sk []string) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var itog []string
	for i := 0; i < len(sk); i++ {
		if strings.Contains(sk[i], "+") {
			parts := strings.Split(sk[i], "+")
			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				itog = append(itog, strconv.Itoa(first+second))
			}
		} else if strings.Contains(sk[i], "-") {
			parts := strings.Split(sk[i], "-")
			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				itog = append(itog, strconv.Itoa(first-second))
			}
		} else if strings.Contains(sk[i], "*") {
			parts := strings.Split(sk[i], "*")
			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				itog = append(itog, strconv.Itoa(first*second))
			}
		} else if strings.Contains(sk[i], "/") {
			parts := strings.Split(sk[i], "/")
			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil && second != 0 {
				itog = append(itog, strconv.Itoa(first/second))
			}
		}
	}
	return itog
}

// считаем операции умножения и деления в выражени без скобок
func mult_div(expression string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for i := 0; i < len(expression); i++ {
		if expression[i] == '*' {
			first, err1 := strconv.Atoi(string(expression[i-1]))
			second, err2 := strconv.Atoi(string(expression[i+1]))
			if err1 == nil && err2 == nil {
				expression = strings.Replace(expression, expression[i-1:i+2], strconv.Itoa(first*second), 1)
			}
		} else if expression[i] == '/' {
			first, err1 := strconv.ParseFloat(string(expression[i-1]), 64)
			second, err2 := strconv.ParseFloat(string(expression[i+1]), 64)
			if second == 0.0 {
				return "", ErrDivisionByZero
			}
			if err1 == nil && err2 == nil && second != 0 {
				expression = strings.Replace(expression, expression[i-1:i+2], strconv.FormatFloat((first/second), 'f', 2, 64), 1)
			}
		}
	}
	return expression, nil
}

// считаем операции сложения и вычитания в выражении без скобок
func add_sub(expression string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' {
			first, err1 := strconv.Atoi(string(expression[i-1]))
			second, err2 := strconv.Atoi(string(expression[i+1]))
			if err1 == nil && err2 == nil {
				expression = strings.Replace(expression, expression[i-1:i+2], strconv.Itoa(first+second), 1)
			}
		} else if expression[i] == '-' {
			first, err1 := strconv.Atoi(string(expression[i-1]))
			second, err2 := strconv.Atoi(string(expression[i+1]))
			if err1 == nil && err2 == nil {
				expression = strings.Replace(expression, expression[i-1:i+2], strconv.Itoa(first-second), 1)
			}
		}
	}
	return expression
}
