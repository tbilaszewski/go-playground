package option

import (
	"fmt"
	"math"
)

type Option[T any] struct {
	value T
	err   error
}

func (o Option[T]) String() string {
	return fmt.Sprintf("Option {value:%v, error:%v}", o.value, o.err)
}

func ToOption[T any](v T, e error) Option[T] {
	return Option[T]{v, e}
}

func UnwrapOr[T any](option Option[T], alternative T) T {
	if option.err == nil {
		return option.value
	}
	return alternative
}

func UnwrapElse[T any](option Option[T], f func(error) T) T {
	if option.err == nil {
		return option.value
	}
	return f(option.err)
}

func Unwrap[T any](option Option[T]) T {
	return UnwrapElse(option, func(e error) T {
		panic(e)
	})
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Negative number %v", float64(e))
}

func Sqrt(v float64) (float64, error) {
	if v < 0 {
		return 0, ErrNegativeSqrt(v)
	}
	return math.Sqrt(v), nil
}

func main() {
	v1 := ToOption(Sqrt(4))
	v2 := ToOption(Sqrt(-4))
	v3 := ToOption(func(v float64) (float64, error) {
		if math.IsNaN(v) {
			return 0, fmt.Errorf("Nan number")
		}
		return v, nil
	}(math.Sqrt(-4)))
	fmt.Println("v3", v3)
	fmt.Println("v1", v1)

	fmt.Println("Option", Unwrap(v1))
	fmt.Println("Option", UnwrapOr(v2, 100))
	UnwrapElse(ToOption(Sqrt(-4)), func(e error) float64 {
		fmt.Printf("Error happended %v\n", e)
		return 0
	})
}
