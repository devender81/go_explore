package main

import (
	"fmt"
)

func main() {
	fmt.Println("Prime numbers up to 100:")
	primes := generatePrimes(100)
	fmt.Println(primes)

	fmt.Println("\nFibonacci sequence up to 1000:")
	fibonacci := generateFibonacci(1000)
	fmt.Println(fibonacci)
}

func generatePrimes(limit int) []int {
	primes := []int{}

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes

}

// Check if a number is prime
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Generate Fibonacci sequence up to a given limit
func generateFibonacci(limit int) []int {
	fibonacci := []int{0, 1}
	for i := 2; ; i++ {
		fibNum := fibonacci[i-1] + fibonacci[i-2]
		if fibNum > limit {
			break
		}
		fibonacci = append(fibonacci, fibNum)
	}
	return fibonacci
}
