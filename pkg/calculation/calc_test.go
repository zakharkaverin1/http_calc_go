package calculation

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedResult float64
		err            error
	}{
		{
			name:           "обычный",
			expression:     "1+1",
			expectedResult: 2,
			err:            nil,
		},
		{
			name:           "со скобками",
			expression:     "2*(2+2)",
			expectedResult: 8,
			err:            nil,
		},
		{
			name:           "с двумя скобками",
			expression:     "(9+1)*(2+2)",
			expectedResult: 40,
			err:            nil,
		},
		{
			name:           "умножение",
			expression:     "2+2*2",
			expectedResult: 6,
			err:            nil,
		},
		{
			name:           "деление",
			expression:     "1/2",
			expectedResult: 0.5,
			err:            nil,
		},
		{
			name:           "деление на ноль",
			expression:     "5/0",
			expectedResult: 0,
			err:            ErrDivisionByZero,
		},

		{
			name:           "бэаэб",
			expression:     "1+32",
			expectedResult: 0,
			err:            ErrSomethingWentWrong,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}
}
