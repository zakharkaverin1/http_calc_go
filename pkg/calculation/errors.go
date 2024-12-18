package calculation

import "errors"

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrInvalidExpression   = errors.New("invalid expression")
	ErrSomethingWentWrong  = errors.New("something went wrong")
	ErrExpressionNotString = errors.New("expression not a string")
)
