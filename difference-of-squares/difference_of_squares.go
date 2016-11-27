package diffsquares

func SquareOfSums(num int) int {
	s := sum(num)
	return s * s
}

func sum(num int) int {
	return (num * (num + 1)) / 2
}

func SumOfSquares(num int) int {
	return (num * (num + 1) * (2 * num + 1)) / 6
}

func Difference(num int) int {
	return SquareOfSums(num)-SumOfSquares(num)
}