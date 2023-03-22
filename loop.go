package loop

func sumOfFirst(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func isPrime(n int) bool {
	count := 0
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count == 1
}
