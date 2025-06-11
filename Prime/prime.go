package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

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

// برای هر عدد، اگه اول باشه، می‌فرستتش داخل channel
func findPrime(n int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	if isPrime(n) {
		ch <- n
	}
}

func main() {
	//دو عدد از کاربر می‌گیریم (شروع و پایان بازه)
	var start, end int
	fmt.Print("عدد شروع را وارد کن: ")
	fmt.Scan(&start)
	fmt.Print("عدد پایان را وارد کن: ")
	fmt.Scan(&end)
	//برای هر عدد از اون بازه، یک Goroutine اجرا می‌کنیم
	//با WaitGroup صبر می‌کنیم تا همه Goroutine‌ها تموم شن
	var wg sync.WaitGroup
	resultChan := make(chan int, end-start+1)

	for i := start; i <= end; i++ {
		wg.Add(1)
		go findPrime(i, &wg, resultChan)
	}
	//بعد از اتمام همه، channel رو می‌بندیم و اعداد اول رو چاپ می‌کنیم
	wg.Wait()
	close(resultChan)

	// ذخیره نتایج در slice
	var primes []int
	for p := range resultChan {
		primes = append(primes, p)
	}

	// مرتب‌سازی
	sort.Ints(primes)

	// نمایش نتایج
	fmt.Println("اعداد اول مرتب‌ شده بین بازه انتخاب‌ شده:")
	for _, p := range primes {
		fmt.Println(p)
	}
}
