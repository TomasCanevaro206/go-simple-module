
package main

import (
	"fmt"
	"errors"
)

func operation(op string) (func(...int) int, error) {

	var calcFunc func(...int) int

	switch op {
		case "minimum":
			calcFunc = func (num ...int) int{
				min := num[0]
				for _,n := range num {
					if n < min {
						min = n
					}
				}
				return min
			}
		case "average":
			calcFunc = func (num ...int) int{
				sum := 0
				for _,n := range num {
					sum += n
				}
				return sum/len(num)
			}
		case "maximum":
			calcFunc = func (num ...int) int{
				max := num[0]
				for _,n := range num {
					if n > max {
						max = n
					}
				}
				return max
			}
		default:
			return nil, errors.New("Operacion invalida")
	}
	return calcFunc, nil
}

func main() {

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	 )
	 
	minFunc, err := operation(minimum)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(minValue)
	fmt.Println(err)

	averageFunc, err := operation(average)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(averageValue)
	fmt.Println(err)

	maxFunc, err := operation(maximum)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(maxValue)
	fmt.Println(err)
}