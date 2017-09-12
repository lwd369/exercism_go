package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	sum := 0
	for ix := 1; ix <= n; ix++ {
		sum += ix
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for ix := 1; ix <= n; ix++ {
		sum += ix * ix
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
