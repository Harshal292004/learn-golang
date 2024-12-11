package errorss

import "fmt"

type divideError struct {
	Dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("error because you stupid dumb as fuck tried to divide the %v by 0", d.Dividend)
}

func Divide(Dividend, Divisor float64) (float64, error) {
	if Divisor == 0 {
		return 0, divideError{Dividend: Dividend}
	}
	return Dividend / Divisor, nil
}
