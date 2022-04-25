package main

import "fmt"

// convertible int types.
type convertible interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// convert integer types from type A to B
func convert[A, B convertible](a A) B {
	return B(a)
}

type Person interface {
	Greet()
}

type Employee[T Person] struct{}

func (e *Employee[T]) Get(data T) T {
	return data
}

func main() {
	// basic generics with contract (meta type)
	var a int8 = 127
	b := convert[int8, uint8](a)
	fmt.Printf("Converted %T to %T with value of %d\n", a, b, b)

}
