package main

import (
	"fmt"
	"math"
	"sync"
)

// برای بررسی کردن اعداد که اول هستن یا نه
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// اینجا بررسی میکنه که عددهای وارد شده اول هستند یا نه
func checkNumber(n int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	if isPrime(n) {
		results <- fmt.Sprintf("%d عدد اول است.", n)
	} else {
		results <- fmt.Sprintf("%d عدد اول نیست.", n)
	}
}

func main() {
	var nums []int
	var count int

	fmt.Print("چند عدد می‌خوای وارد کنی؟ ")
	fmt.Scan(&count)

	fmt.Println("اعداد رو وارد کن:")
	for i := 0; i < count; i++ {
		var x int
		fmt.Scan(&x)
		nums = append(nums, x)
	}

	var wg sync.WaitGroup
	results := make(chan string, len(nums))

	for _, n := range nums {
		wg.Add(1)
		go checkNumber(n, &wg, results)
	}

	wg.Wait()
	close(results)

	fmt.Println("\nنتایج:")
	for res := range results {
		fmt.Println(res)
	}
}
